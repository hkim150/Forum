import React, { useEffect, useState } from "react";
import PostTable from "../components/PostTable";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function Main() {
    const [posts, setPosts] = useState([]);
    const navigate = useNavigate();
    
    useEffect(() => {
        axios.get(`${import.meta.env.VITE_BACKEND_API_URL}/posts`)
            .then(response => setPosts(response.data))
            .catch(error => console.error("Error fetching data:", error))
    }, []);

    const handleDelete = (id) => {
        axios.delete(`${import.meta.env.VITE_BACKEND_API_URL}/posts${id}`)
            .then(() => {
                setPosts(posts.filter(post => post.Id !== id));
            })
            .catch(error => console.error("Error deleting post:", error));
    }

    return (
        <>
            <button onClick={() => navigate("/create-post")}>Create Post</button>
            <PostTable posts={posts} onDelete={handleDelete}/>
        </>
    );
}

export default Main;