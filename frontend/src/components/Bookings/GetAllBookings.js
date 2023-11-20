// GetAllBookings.js

import React, { useState } from 'react';

const GetAllBookings = () => {
  const [formData, setFormData] = useState({
    booking_id: '',
  });

  const handleGetAllBookings = async () => {
    try {
      const response = await fetch('http://localhost:8080/get-all-bookings',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Booking deleted successfully
        console.log('bookings acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire bookings');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring bookings:', error.message);
    }
  };

  return (
    <div>
      <h2>Get All bookings</h2>
      <form>
        <button type="button" onClick={handleGetAllBookings}>
          Get All bookings
        </button>
      </form>
    </div>
  );
};

export default GetAllBookings;
