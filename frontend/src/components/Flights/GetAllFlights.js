// GetAllFlights.js

import React, { useState } from 'react';

const GetAllFlights = () => {
  const [formData, setFormData] = useState({
    flight_id: '',
  });

  const handleGetAllFlights = async () => {
    try {
      const response = await fetch('http://localhost:8080/get-all-flights',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Flight deleted successfully
        console.log('flights acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire flights');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring flights:', error.message);
    }
  };

  return (
    <div>
      <h2>Get All flights</h2>
      <form>
        <button type="button" onClick={handleGetAllFlights}>
          Get All flights
        </button>
      </form>
    </div>
  );
};

export default GetAllFlights;
