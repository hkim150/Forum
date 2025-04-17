import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';

function EditPost() {
    const { id } = useParams();
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`${import.meta.env.VITE_BACKEND_API_URL}/posts/${id}`)
            .then(response => {
                setTitle(response.data.Title);
                setContent(response.data.Content);
            })
            .catch(error => console.error("Error fetching post:", error));
    }, [id]);

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.put(`${import.meta.env.VITE_BACKEND_API_URL}/posts/${id}`, { title, content })
            .then(() => navigate(`/posts/${id}`))
            .catch(error => console.error("Error updating post:", error));
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Title:</label>
                <input
                    type="text"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Content:</label>
                <textarea
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Save Changes</button>
            <button type="button" onClick={() => navigate(-1)}>Cancel</button>
        </form>
    );
}

export default EditPost;