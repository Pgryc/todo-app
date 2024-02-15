import React, { useEffect, useState } from 'react';
import Task from "./Task";

export default function TaskList() {
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        async function fetchData() {
            try {
                const response = await fetch('http://127.0.0.1:8080/v1/tasks');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const jsonData = await response.json();
                setData(jsonData);
                setLoading(false);
            } catch (error) {
                setError(error);
                setLoading(false);
            }
        }

        fetchData();
    }, []);

    async function updateData(id, updatedData) {
        try {
            const response = await fetch('http://127.0.0.1:8080/v1/tasks/${id}', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(updatedData),
            });
            if (!response.ok) {
                throw new Error('Network responsw was not ok')
            }
            const updatedItem = await response.json();

            setData(prevData =>
                prevData.map(item => (item.id === id ? updatedItem : item))
            );
        } catch (error) {
            setError(error);
        }
    }

    if (loading) return (
        <div>Loading...</div>
    )

    if (error) return (
        <div>Error: {error.message}</div>
    )

    return (
        <div className="TaskList">
            TASK LIST
            <Task></Task>
            <Task></Task>
        </div>
    )
}