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

    return (
        <>
            <button onClick={() => navigate("/create-post")}>Create Post</button>
            <PostTable posts={posts}/>
        </>
    );
}

export default Main;