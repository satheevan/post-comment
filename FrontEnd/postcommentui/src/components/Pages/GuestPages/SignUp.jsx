import React from "react";

function SignUp() {
    return (<>
        <h1>SignUp Page</h1>
        <div className="RegisterContainer">
            <form className="RegisterContainer">
                <label>EmailId :</label>
                <input type="text" placeholder="EmailId" required></input>
                <label>Receiver :</label>
                <input type="password" placeholder="Receiver Code" required></input>
                <button type="button">Submit</button>
            </form>
        </div>
    </>)
}
export default SignUp