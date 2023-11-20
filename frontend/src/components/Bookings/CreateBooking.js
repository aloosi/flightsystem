import React, { useState } from 'react';

const CreateBooking = () => {
  const [formData, setFormData] = useState({
    booking_id: '',
    tourist_id: '',
    flight_id: '',
    payment_id: '',
    booking_date: '',
  });

  const handleCreateBooking = async () => {
    try {
      console.log('Form Data:', formData); // Log the formData
  
      const response = await fetch('http://3.134.76.216:8080/create-booking', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
  
      if (response.ok) {
        // Booking created successfully
        console.log('Booking created successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to create booking. Status:', response.status);
        const errorData = await response.json();
        console.error('Error message:', errorData);
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating booking:', error.message);
    }
  };
  

  return (
    <div>
      <h2>Create Booking</h2>
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

        <button type="button" onClick={handleCreateBooking}>
          Create Booking
        </button>
      </form>
    </div>
  );
};

export default CreateBooking;
