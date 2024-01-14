import React from "react";


function Login() {
    return (
        <div>
            <h1>Heading the Login page</h1>
            <div className="loginContainer">
                <form className="loginForm">
                    <label>UserName :</label>
                    <input type="text" placeholder="Username" required></input>
                    <label>Password :</label>
                    <input type="password" placeholder="Password" required></input>
                    <button type="button">Submit</button>
                </form>
            </div>
        </div>
    )
}

export default Login