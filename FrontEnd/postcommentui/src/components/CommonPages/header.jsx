
import React from "react";
import './css/header.css'
// import logo from './components/asset/logo.svg';

function Header() {
    return (
        <header className="header">
            {/* right side */}
            {/* Logo */}
            <img alt="cpylogo" src=""></img>
            <div className="headerLeftcontainer">
                {/* Company name */}
                <h1>HTSG COMPANY</h1>
                {/* title */}
                <span className="titleyear">since:2014</span>
            </div>
            {/* left side */}
            <div className="headerRightcontainer">
                <ul className="menu">
                    <li><a href="/"><span className="fa"></span>Profile</a></li>
                    <li><a href="/"><span className="fa"></span>Settings</a></li>
                    <li><a href="/"><span className="fa"></span>Notifications</a></li>
                    <li><a href="/"><span className="fa"></span>Logout</a></li>
                </ul>
            </div>
        </header>
    )
}
export default Header;