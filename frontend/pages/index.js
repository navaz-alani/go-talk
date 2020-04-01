import React from "react";
import Link from 'next/link'
import Jumbotron from "react-bootstrap/Jumbotron"
import Button from "react-bootstrap/Button"
import Footer from "../components/footer/footer";

const App = () => {
    return (
        <div id="main-container">
        <Jumbotron id="main">
            <h1>Go-Talk</h1>
            <p>
                Talk, un-traced.
            </p>
            <p id="app-description">
                Go-Talk is an ephemeral messaging service<br/>
                which does not track any user information <br/>
                whatsoever.
            </p>
            <p>
                <Link href="/auth">
                    <Button variant="primary">Login / Sign-Up</Button>
                </Link>
            </p>
        </Jumbotron>
            <Footer/>
        </div>
    );
};

export default App;
