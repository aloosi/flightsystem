// GetFlightByID.js

import React, { useState } from 'react';

const GetFlightByID = () => {
  const [formData, setFormData] = useState({
    flight_id: '',
  });

  const handleGetFlightByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-flight-by-id/' + formData.flight_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Flight deleted successfully
        console.log('Flight acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire flight');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring flight:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Flight By ID</h2>
      <form>
        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetFlightByID}>
          Get Flight By ID
        </button>
      </form>
    </div>
  );
};

export default GetFlightByID;
