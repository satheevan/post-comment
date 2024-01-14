import React, { useState, useEffect } from "react";
import axios from "axios";

function PostList(props) {
    // hardcode
    // const postItems = ['English', 'developer', 'science', 'maths'];

    // const listItems = []

    // for (let i = 0; postItems.length > i; i++) {
    //     listItems.push(<li key={i}>{postItems[i]}</li>)
    // }
    //Using API call
    const [postItems, setPostItems] = useState([]);

    useEffect(() => {
        // fetch data from golang(backend) end Points
        axios.get('http://localhost:3003/api/getpost')
            .then(res => {
                console.log("@@@@=>", res.data.Data.data);
                setPostItems(res.data.Data.data);

            })
            .catch(error => {
                console.error("error feteching the data from database", error);
            });
    }, [])

    console.log("$$ ===>", postItems);

    return (
        <div className="Postlist">
            <h1>Post List</h1>
            <ul className="postTitle" >
                {/* {listItems} */}
                {/* {postItems} */}
                {
                    postItems.map(item => (
                        // <li key={item.id}><a href="#" onClick={() => props.onPostSelect(item)}>{item.Title}</a></li>
                        <li key={item.id}><a href='#' onClick={() => props.onSelectPost(item)}>{item.Title}</a></li>
                    ))
                }
            </ul>
        </div>
    )
}

export default PostList