import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';

function PostContent() {
    const { id } = useParams();
    const [post, setPost] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`http://localhost:8080/api/v1/posts/${id}`)
            .then(response => setPost(response.data))
            .catch(error => console.error("Error fetching post:", error));
    }, [id]);

    if (!post) {
        return <div>Loading...</div>
    }

    return (
        <div>
            <h2>{post.Title}</h2>
            <p>{post.Content}</p>
            <p><strong>Created At:</strong> {post.CreatedAt}</p>
            <button onClick={() => navigate(-1)}>Go Back</button>
        </div>
    );
}

export default PostContent;