// GetBookingByID.js

import React, { useState } from 'react';

const GetBookingByID = () => {
  const [formData, setFormData] = useState({
    booking_id: '',
  });

  const handleGetBookingByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-booking-by-id/' + formData.booking_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Booking deleted successfully
        console.log('Booking acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire booking');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring booking:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Booking By ID</h2>
      <form>
        <label>Booking ID:</label>
        <input
          type="text"
          value={formData.booking_id}
          onChange={(e) => setFormData({ ...formData, booking_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetBookingByID}>
          Get Booking By ID
        </button>
      </form>
    </div>
  );
};

export default GetBookingByID;
