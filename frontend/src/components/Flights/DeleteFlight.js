// DeleteFlight.js

import React, { useState } from 'react';

const DeleteFlight = () => {
  const [formData, setFormData] = useState({
    flight_id: '',
  });
  const [message, setMessage] = useState(null);


  const handleDeleteFlight = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/delete-flight/' + formData.flight_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Flight deleted successfully
        console.log('Flight deleted successfully');
        setMessage('Flight deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete flight');
        setMessage('Failed to delete flight');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting flight:', error.message);
      setMessage('Error deleting flight:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Flight</h2>
      <form>
        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeleteFlight}>
          Delete Flight
        </button>
      </form>
<hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default DeleteFlight;
