package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/navaz-alani/go-talk/env"

	"github.com/dgrijalva/jwt-go"
)

/*
The following variables specify the options used
when generating JWTs. When a variable is derived
form the environment variables, the key under which
it should be specified is also highlighted.

Please note that if any of the variables initialized
from the application's env vars is not specified, the
JWTCheck function will fail and the app will exit
with error code 1.
*/
var (
	/*
		tokenLife specifies the time till expiry
		of a token after issuance.
		This value cannot be specified from the
		environment variables: it must be set
		before building the app.
	*/
	tokenLife int64 = 3600 * 6
	/*
		signingSecret is the key used to sign
		JWT tokens. It is derived from the
		application's env variables under the
		key "JWT_SS".
	*/
	signingSecret []byte
	/*
		issuer is a string specifying the issuing
		authority of the JWT. It is derived from the
		application's env variables under the key
		"JWT_ISS".
	*/
	issuer string
)

/*
JWTCheck ensures that all of the information needed to issue
a JSON web token has been initialized correctly.

A successful check will allow program execution to continue,
otherwise the program will exit, printing the issues that
were encountered.
*/
func JWTCheck() {
	// read parameters from env variables
	signingSecret = []byte(env.Get("JWT_SS"))
	issuer = env.Get("JWT_ISS")
}

/*
EncodePayload returns a signed JWT token containing
the given payload string.
*/
func EncodePayload(payload string) (string, error) {
	claims := &jwt.MapClaims{
		"exp": time.Now().Unix() + tokenLife,
		"iss": issuer,
		// custom claim
		"uid": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingSecret)
}

/*
JWTVerify returns the associated JWT in the request and
a HTTP Status code for the authorization.
*/
func JWTVerify(r *http.Request) (*jwt.Token, int) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")

	if len(splitToken) != 2 {
		return nil, http.StatusUnauthorized
	}

	token, err := VerifyJWT(strings.TrimSpace(splitToken[1]))
	if err == nil && token.Valid {
		return token, http.StatusOK
	} else {
		return token, http.StatusUnauthorized
	}
}

/*
VerifyJWT verifies the given JWT string.
*/
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				token.Header["alg"])
		}

		return signingSecret, nil
	})
}

/*
JWTClaims returns the claim set associated with
the given JWT.
*/
func JWTClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
