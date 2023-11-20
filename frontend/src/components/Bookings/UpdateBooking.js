import React, { useState } from 'react';

const UpdateBooking = () => {
  const [formData, setFormData] = useState({
    booking_id: '',
    tourist_id: '',
    flight_id: '',
    payment_id: '',
    booking_date: '',
  });

  const handleUpdateBooking = async () => {
    try {
      const response = await fetch('http://localhost:8080/update-booking', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Booking updated successfully
        console.log('Booking updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update booking');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating booking:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Booking</h2>
      <form>
        <label>Booking ID:</label>
        <input
          type="text"
          value={formData.booking_id}
          onChange={(e) => setFormData({ ...formData, booking_id: parseInt(e.target.value) })}
        />
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value) })}
        />

        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value) })}
        />

        <label>Payment ID:</label>
        <input
          type="text"
          value={formData.payment_id}
          onChange={(e) => setFormData({ ...formData, payment_id: parseInt(e.target.value) })}
        />

        <label>Booking Date:</label>
        <input
          type="date"
          value={formData.booking_date}
          onChange={(e) => setFormData({ ...formData, booking_date: e.target.value })}
        />

        <button type="button" onClick={handleUpdateBooking}>
          Create Booking
        </button>
      </form>
    </div>
  );
};

export default UpdateBooking;
