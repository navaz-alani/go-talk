package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/navaz-alani/go-talk/core/auth"
)

// authReq defines an authentication request for
// a user in the context of this application.
type authReq struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Auth is an http HandlerFunc which handles HTTP
// requests to handle authentication requests to
// the core service.
func Auth(w http.ResponseWriter, r *http.Request) {
	var req authReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, "request read fail", http.StatusBadRequest)
		return
	}

	user := User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
	}

	if req.Type == "create" {
		res := Users.PStorage.FindOne(context.TODO(),
			bson.M{"$or": []bson.M{
				{"username": user.Username},
				{"email": user.Email},
			}},
		)
		if res.Err() != mongo.ErrNoDocuments {
			http.Error(w, "non-unique", http.StatusUnauthorized)
			return
		} else {
			h, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
			if err != nil {
				http.Error(w, "pwd hash fail",
					http.StatusInternalServerError)
				return
			}
			user.PwdHash = string(h)
			user.UID = auth.Key(25)

			_, err = Users.Add(user)
			if err != nil {
				http.Error(w, "failed to add",
					http.StatusInternalServerError)
				return
			}

			jwt, err := auth.EncodePayload(user.UID)
			if err != nil {
				http.Error(w, "token gen fail",
					http.StatusInternalServerError)
				return
			}

			_, _ = w.Write([]byte(jwt))
		}
	} else if req.Type == "verify" {
		if ok, err := Users.Exists(user, &user); err != mongo.ErrNoDocuments &&
			err != nil {
			log.Println(err)
			http.Error(w, "database error",
				http.StatusInternalServerError)
			return
		} else {
			if !ok {
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}

			// verify password hash
			if err := bcrypt.CompareHashAndPassword(
				[]byte(user.PwdHash),
				[]byte(req.Password),
			); err != nil {
				http.Error(w, "verification fail",
					http.StatusInternalServerError)
				return
			}

			// user is allowed, send jwt
			jwt, err := auth.EncodePayload(user.UID)
			if err != nil {
				http.Error(w, "token gen fail",
					http.StatusInternalServerError)
				return
			}
			_, _ = w.Write([]byte(jwt))
		}
	} else {
		http.Error(w, "protocol unknown", http.StatusBadRequest)
		return
	}
}
