import React from "react";
import Link from 'next/link'
import Jumbotron from "react-bootstrap/Jumbotron"
import Button from "react-bootstrap/Button"

const App = () => {
    return (
        <Jumbotron id="main">
            <h1>Go-Talk</h1>
            <p>
                A Private Messaging Service.
            </p>
            <p>
                <Link href="/auth">
                    <Button variant="primary">Login / Sign-Up</Button>
                </Link>
            </p>
        </Jumbotron>
    );
};

export default App;
