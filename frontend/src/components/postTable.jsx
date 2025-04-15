import React from "react";
import { Link } from "react-router-dom";

function PostTable(props) {
    return (
        <table>
        <thead>
            <tr>
                <th>Id</th>
                <th>Title</th>
                <th>Date</th>
                <th>Delete</th>
            </tr>
        </thead>
        <tbody>
            {props.posts.map(post => (
            <tr key={post.Id}>
                <td>{post.Id}</td>
                <td>
                    <Link to={`/posts/${post.Id}`}>{post.Title}</Link>
                </td>
                <td>{post.CreatedAt}</td>
                <td>
                    <button onClick={() => props.onDelete(post.Id)}>Delete</button>
                </td>
            </tr>
            ))}
        </tbody>
        </table>
    );
}

export default PostTable;