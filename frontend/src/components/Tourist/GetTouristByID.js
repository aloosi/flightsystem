// GetTouristByID.js

import React, { useState } from 'react';

const GetTouristByID = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
  });

  const handleGetTouristByID = async () => {
    try {
      const response = await fetch('http://localhost:8080/get-tourist-by-id/' + formData.tourist_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Tourist deleted successfully
        console.log('Tourist acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire tourist');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring tourist:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Tourist By ID</h2>
      <form>
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetTouristByID}>
          Get Tourist By ID
        </button>
      </form>
    </div>
  );
};

export default GetTouristByID;
