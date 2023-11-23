import React, { useState } from 'react';

const GetAllBookings = () => {
  const [bookingsData, setBookingsData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllBookings = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-bookings', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setBookingsData(data); // Set the bookings data directly
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire bookings. Status: ${response.status}, Error: ${errorData.message}`);
        setBookingsData([]); // Set empty array in case of error
      }
    } catch (error) {
      setError(`Error acquiring bookings: ${error.message}`);
      setBookingsData([]); // Set empty array in case of error
    }
  };

  return (
    <div>
      <h2>Get All Bookings</h2>
      <form>
        <button type="button" onClick={handleGetAllBookings}>
          Get All Bookings
        </button>
      </form>

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        bookingsData.length > 0 && (
          <div>
            <h3>Bookings Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Booking ID</th>
                  <th>Flight ID</th>
                  <th>Tourist ID</th>
                  {/* Add more columns as needed */}
                </tr>
              </thead>
              <tbody>
                {bookingsData.map((booking) => (
                  <tr key={booking.booking_id}>
                    <td>{booking.booking_id}</td>
                    <td>{booking.flight_id}</td>
                    <td>{booking.tourist_id}</td>
                    {/* Add more cells for additional data */}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )
      )}
    </div>
  );
};

export default GetAllBookings;
