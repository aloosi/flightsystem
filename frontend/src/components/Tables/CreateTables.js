import React, { useState } from 'react';

const CreateTables = () => {
  const [formData, setFormData] = useState({
    table_id: '',
    table_name: '',
  });

  const [message, setMessage] = useState(null);

  const handleCreateTables = async () => {
    try {
        const response = await fetch('http://3.134.76.216:8080/create-tables',{
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
            },
      });
  
      if (response.ok) {
        // Table created successfully
        console.log('Tables created successfully');
        setMessage('Tables created successfully');

        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to create tables. Status:', response.status);
        const errorData = await response.json();
        console.error('Error message:', errorData);
        setMessage('Error message:', errorData);

      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating tables:', error.message);
      setMessage('Error creating tables:', error.message);

    }

  };
  

  return (
    <div>
      <h2>Create Tables</h2>
      <form>

        <button type="button" onClick={handleCreateTables}>
          Create Tables
        </button>
      </form>
      <hr />
      {message && <p>{message}</p>}
    </div>
    
  );
};

export default CreateTables;
