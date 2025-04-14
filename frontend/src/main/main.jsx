import React, { useEffect, useState } from "react";
import PostTable from "../components/postTable";
import axios from "axios";

function Main() {
    const [posts, setPosts] = useState([]);
    
    useEffect(() => {
        axios.get("http://localhost:8080/api/v1/posts")
            .then(response => setPosts(response.data))
            .catch(error => console.error("Error fetching data:", error))
    }, []);

    return (
        <>
            <PostTable posts={posts}/>
        </>
    );
}

export default Main;