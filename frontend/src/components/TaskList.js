import React, { useEffect, useState } from "react";
import Task from "./Task";

export default function TaskList() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const api_url = process.env.REACT_APP_API_URL;

  useEffect(() => {
    async function fetchData() {
      try {
        const response = await fetch(`${api_url}/v1/tasks`);
        console.log(response);
        if (!response.ok) {
          throw new Error("Network response was not ok");
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
  }, [api_url]);

  async function updateData(id, updatedData) {
    try {
      const response = await fetch(`${api_url}/v1/tasks/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedData),
      });
      if (!response.ok) {
        throw new Error("Network responsw was not ok");
      }
      const updatedItem = await response.json();

      setData((prevData) =>
        prevData.map((item) => (item.id === id ? updatedItem : item)),
      );
    } catch (error) {
      setError(error);
    }
  }

  if (loading) return <div>Loading...</div>;

  if (error) return <div>Error: {error.message}</div>;

  return (
    <div className="TaskList">
      TASK LIST
      <Task></Task>
      <Task></Task>
    </div>
  );
}
