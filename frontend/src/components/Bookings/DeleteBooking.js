// DeleteBooking.js

import React, { useState } from 'react';

const DeleteBooking = () => {
  const [formData, setFormData] = useState({
    booking_id: '',
  });
  const [message, setMessage] = useState(null);

  const handleDeleteBooking = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/delete-booking/' + formData.booking_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Booking deleted successfully
        console.log('Booking deleted successfully');
        setMessage('Booking deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete booking');
        setMessage('Failed to delete booking');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting booking:', error.message);
      setMessage('Error deleting booking:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Booking</h2>
      <form>
        <label>Booking ID:</label>
        <input
          type="text"
          value={formData.booking_id}
          onChange={(e) => setFormData({ ...formData, booking_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeleteBooking}>
          Delete Booking
        </button>
      </form>
<hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default DeleteBooking;
