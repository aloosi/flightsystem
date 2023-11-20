// GetAllAirlines.js

import React, { useState } from 'react';

const GetAllAirlines = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
  });

  const handleGetAllAirlines = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-airlines',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Airline deleted successfully
        console.log('airlines acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire airlines');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring airlines:', error.message);
    }
  };

  return (
    <div>
      <h2>Get All airlines</h2>
      <form>
        <button type="button" onClick={handleGetAllAirlines}>
          Get All airlines
        </button>
      </form>
    </div>
  );
};

export default GetAllAirlines;
