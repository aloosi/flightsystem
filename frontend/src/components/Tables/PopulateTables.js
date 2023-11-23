import React, { useState } from 'react';

const PopulateTables = () => {
  const [formData, setFormData] = useState({
    table_id: '',
    table_name: '',
  });
  const [message, setMessage] = useState(null);

  const handlePopulateTables = async () => {
    try {
        const response = await fetch('http://3.134.76.216:8080/populate-tables',{
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
            },
      });
  
      if (response.ok) {
        // Table populated successfully
        console.log('Tables populated successfully');
        setMessage('Tables populated successfully');

        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to populate tables. Status:', response.status);
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
      <h2>Populate Tables</h2>
      <form>

        <button type="button" onClick={handlePopulateTables}>
          Populate Tables
        </button>
      </form>
      <hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default PopulateTables;
