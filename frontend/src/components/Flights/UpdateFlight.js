import React, { useState } from 'react';

const UpdateFlight = () => {
  const [formData, setFormData] = useState({
    flight_id: '',
    airline_id: '',
    flight_number: '',
    departure_country: '',
    flight_date: '',
    arrival_country: '',
    departure_time: '',
    arrival_time: '',
    flight_class: '',
    flight_cost: '',
  });

  const handleUpdateFlight = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/update-flight', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Flight updated successfully
        console.log('Flight updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update flight');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating flight:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Flight</h2>
      <form>
        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value) })}
        />
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value) })}
        />

        <label>Flight Number:</label>
        <input
          type="text"
          value={formData.flight_number}
          onChange={(e) => setFormData({ ...formData, flight_number: e.target.value })}
        />

        <label>Departure Country:</label>
        <input
          type="text"
          value={formData.departure_country}
          onChange={(e) => setFormData({ ...formData, departure_country: e.target.value })}
        />

        <label>Flight Date:</label>
        <input
          type="date"
          value={formData.flight_date}
          onChange={(e) => setFormData({ ...formData, flight_date: e.target.value })}
        />

        <label>Arrival Country:</label>
        <input
          type="text"
          value={formData.arrival_country}
          onChange={(e) => setFormData({ ...formData, arrival_country: e.target.value })}
        />

        <label>Departure Time:</label>
        <input
          type="text"
          value={formData.departure_time}
          onChange={(e) => setFormData({ ...formData, departure_time: e.target.value })}
        />

        <label>Arrival Time:</label>
        <input
          type="text"
          value={formData.arrival_time}
          onChange={(e) => setFormData({ ...formData, arrival_time: e.target.value })}
        />

        <label>Flight Class:</label>
        <input
          type="text"
          value={formData.flight_class}
          onChange={(e) => setFormData({ ...formData, flight_class: e.target.value })}
        />

        <label>Flight Cost:</label>
        <input
          type="number"
          value={formData.flight_cost}
          onChange={(e) => setFormData({ ...formData, flight_cost: parseFloat(e.target.value) })}
        />

        <button type="button" onClick={handleUpdateFlight}>
          Update Flight
        </button>
      </form>
    </div>
  );
};

export default UpdateFlight;
