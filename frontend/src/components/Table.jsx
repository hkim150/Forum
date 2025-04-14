import axios from "axios";
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

function Table() {
    const [data, setData] = useState([]);
    
    useEffect(() => {
        axios.get("http://localhost:8080/api/v1/posts")
            .then(response => setData(response.data))
            .catch(error => console.error("Error fetching data:", error))
    }, []);

    return (
        <table>
        <thead>
            <tr>
                <th>Id</th>
                <th>Title</th>
                <th>Date</th>
            </tr>
        </thead>
        <tbody>
            {data.map(post => (
            <tr key={post.Id}>
                <td>{post.Id}</td>
                <td>
                    <Link to={`/posts/${post.Id}`}>{post.Title}</Link>
                </td>
                <td>{post.CreatedAt}</td>
            </tr>
            ))}
        </tbody>
        </table>
    );
}

export default Table;