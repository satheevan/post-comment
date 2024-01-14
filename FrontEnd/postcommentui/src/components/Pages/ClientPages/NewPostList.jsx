import React from "react";
import './css/NewPostList.css'
import { useState } from "react";
import { IoIosSearch } from "react-icons/io";
import { IoMdAddCircle } from "react-icons/io";

// function postlist(props) {
//     return (
//         <li><a href="#" onClick={e => props.setMenutems(props.list.title)}>{props.list.title}</a></li>
//     )
// }

function NewPostList() {

    const postTitle = [{ name: "name0", title: "maths", createdAt: "yesterday" },
    { name: "name1", title: "English", createdAt: "today" },
    { name: "name2", title: "Social", createdAt: "yesterday" },
    { name: "name3", title: "science", createdAt: "today" },

    ]

    //Click items
    const [menuItems, setMenutems] = useState([])

    const [message, setpostview] = useState('');

    const handleMessageChange = (e) => {
        console.log(e)
        setpostview(e.target.value);
    }


    return (
        <>
            <div className="postListContainer">
                <nav className="sidebarContainer">
                    <div className="createPostContainer">
                        <IoIosSearch className="searchIcon" />
                        <input id="createPostInput" type="text" placeholder="New Post" size={12}></input>
                        <button type="submit" className="search-button"><IoMdAddCircle className="search-addicon" /></button>
                    </div>
                    <div className="postsContainer">
                        <strong>Post List :</strong>
                        <ul>
                            {postTitle.map(item => (
                                <li><span className="first-item">{item.name}</span><a href="#" className="second-item" onClick={e => setMenutems(item)}>{item.title}</a><span className="third-item">{item.createdAt}</span></li>
                            ))}

                        </ul>

                    </div>
                </nav>
                <section className="sectionContainer">
                    <div className="postview-container">
                        <h1>Post view</h1>
                        <h1 className="post-title-name">{menuItems.title}</h1>
                        <span className="post-createdAt">{menuItems.createdAt}</span>
                        <h1>{message}</h1>
                    </div>
                    <div className="comment-box-container">
                        <input type="text" className="comment-input-box" value={message} placeholder="Enter a comment" onChange={(e) => handleMessageChange(e)} size={40} />
                        <span><button className="comment-submit-button">submit</button></span>
                    </div>
                </section>
            </div>
            {/* cloasing */}
        </>
    )
}
export default NewPostList