import Axios from "axios";
import Cookies from "universal-cookie"

import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Tabs from "react-bootstrap/Tabs";
import Tab from "react-bootstrap/Tab";
import field from "../../components/field/fields";
import {w3cwebsocket as W3CWebsocket} from "websocket"

import getConfig from 'next/config'
import Footer from "../../components/footer/footer";
const { publicRuntimeConfig } = getConfig();

const Auth = () => {
    let cookies = new Cookies();
    let authTok = cookies.get(publicRuntimeConfig.AUTH_COOKIE);
    if (authTok !== "") {
        let ws = new W3CWebsocket(`${publicRuntimeConfig.BE.replace("http", "ws")}/connect`,
            authTok)

    }

    let name = new field(
        "Name", "", "text",
        "Your Name", /(.|\n)*?/, ""
    );
    let email = new field(
        "Email Address", "", "email",
        "Email Address",
        /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/,
        "Email invalid.",
    );
    let password = new field(
        "Password", "", "password",
        "Password", /^.{8,}$/,
        "Password too short."
    );
    let confPassword = new field(
        "Confirm Password", "", "password",
        "Password", /^.{8,}$/,
        "Password too short."
    );
    let username = new field(
        "Username", "", "text",
        "Username", /(.|\n)*?/, ""
    );

    let loginFields = [email, password];
    let signUpFields = [name, email, username, password, confPassword];

    let [loginErr, setLoginErr] = useState("");
    let [signupErr, setSignupErr] = useState("");

    const setAuthErr = (data) => {
        if (data.type === "verify") {
            setLoginErr("Error: Login Failed");
        } else {
            setSignupErr("Error: Sign-up Failed!")
        }
    };

    const auth = (type, fieldSet) => {
        let validateFail = false;
        fieldSet.map(f => { if (!f.validate()) { validateFail = true } });
        if (validateFail) {
            return;
        }

        let data = {
            type: type,
            email: email.value,
            password: password.value,
        };

        if (type === "verify") {
        } else if (type === "create") {
            if (password.value !== confPassword.value) {
                setSignupErr("Error: passwords don't match");
                return
            }
            data.name = name.value;
            data.username = username.value;
        } else {
            return;
        }

        Axios.post(`${publicRuntimeConfig.BE}/auth`, data)
            .then(r => {
                if (r.status !== 200) {
                    setAuthErr(data);
                } else {
                    cookies.set(publicRuntimeConfig.AUTH_COOKIE,
                        r.data, { path: "/" });
                }
            })
            .catch(() => {
                setAuthErr(data);
            })
    };

    return (
        <div id="main-container">
        <div id="auth-window">
            <h1 className="auth-heading main-name">go-talk</h1>
            <h2 className="auth-heading">Authentication</h2>
            <Tabs id="auth-tabs" defaultActiveKey="login">
                <Tab eventKey="login" title="Login">
                    <Form>
                        <br />
                        {loginFields.map((f) => {
                            return f.paint();
                        })}
                    </Form>
                    <div className="field-error">
                        {loginErr}
                    </div>
                    <div className="auth-submit">
                        <Button
                            onClick={() => { auth("verify", loginFields) }}>
                            Login
                        </Button>
                    </div>
                </Tab>
                <Tab eventKey="sign-up" title="Sign-Up">
                    <Form>
                        <br />
                        {signUpFields.map((f) => { return f.paint() })}
                    </Form>
                    <div className="field-error">
                        {signupErr}
                    </div>
                    <div className="auth-submit">
                        <Button
                            onClick={() => { auth("create", loginFields) }}>
                            Sign-Up
                        </Button>
                    </div>
                </Tab>
            </Tabs>
        </div>
        <Footer/>
            </div>
    );
};

export default Auth
