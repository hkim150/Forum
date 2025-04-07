import axios from "axios";
import React, { useEffect, useState } from "react";

function Table() {
    // const [data, setData] = useState([]);
    
    // useEffect(() => {
    //     axios.get("http://localhost:8080/api/data")
    //         .then(response => setData(response.data.persons))
    //         .catch(error => console.error("Error fetching data:", error))
    // }, []);

    const data = [
        { id: 1, name: "John Doe", author: "Jane Smith", date: "2023-10-01" },
        { id: 2, name: "Alice Johnson", author: "Bob Brown", date: "2023-10-02" },
        { id: 3, name: "Charlie Davis", author: "Eve White", date: "2023-10-03" },
        { id: 4, name: "David Wilson", author: "Frank Green", date: "2023-10-04" },
        { id: 5, name: "Emma Thompson", author: "Grace Black", date: "2023-10-05" }
    ]

    return (
        <table>
        <thead>
            <tr>
            <th>Name</th>
            <th>Author</th>
            <th>Date</th>
            </tr>
        </thead>
        <tbody>
            {data.map(post => (
            <tr key={post.id}>
                <td>{post.name}</td>
                <td>{post.author}</td>
                <td>{post.date}</td>
            </tr>
            ))}
        </tbody>
        </table>
    );
}

export default Table;