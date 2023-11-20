// GetAirlineByID.js

import React, { useState } from 'react';

const GetAirlineByID = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
  });

  const handleGetAirlineByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-airline-by-id/' + formData.airline_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Airline deleted successfully
        console.log('Airline acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire airline');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring airline:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Airline By ID</h2>
      <form>
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetAirlineByID}>
          Get Airline By ID
        </button>
      </form>
    </div>
  );
};

export default GetAirlineByID;
