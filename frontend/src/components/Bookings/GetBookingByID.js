import React, { useState } from 'react';

const GetBookingByID = () => {
  const [bookingData, setBookingData] = useState(null);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({
    booking_id: '',
  });

  const handleGetBookingByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-booking-by-id/' + formData.booking_id, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setBookingData(data); // Set the booking data directly
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire booking. Status: ${response.status}, Error: ${errorData.message}`);
        setBookingData(null); // Set null in case of error
      }
    } catch (error) {
      setError(`Error acquiring booking: ${error.message}`);
      setBookingData(null); // Set null in case of error
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

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        bookingData && (
          <div>
            <h3>Booking Data:</h3>
            <p>Booking ID: {bookingData.booking_id}</p>
            <p>Flight ID: {bookingData.flight_id}</p>
            <p>Tourist ID: {bookingData.tourist_id}</p>
            {/* Add more details as needed */}
          </div>
        )
      )}
    </div>
  );
};

export default GetBookingByID;
