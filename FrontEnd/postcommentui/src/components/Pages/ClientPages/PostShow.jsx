import axios from 'axios';
import React, { useEffect, useState } from 'react'

export const PostShow = (props) => {

    const [comments, setComments] = useState(null);

    useEffect(() => { // after mount

        axios.get('http://localhost:3004/api/get-comments/' + props.Id).then((res) => {
            setComments(res.data)
        }).catch(err => {
            console.log('Error in getting the comments by post ', err)
        })
        console.log(comments);
    }, [props.Id]);

    return (
        <div>
            <h1>{props.Title}</h1>
            {comments && comments.length && (<ul>
                {comments.map((comment, idx) => (
                    <li key={idx}>
                        <div>{comment.description}</div>
                    </li>
                ))}
            </ul>)}
        </div>
    )
}