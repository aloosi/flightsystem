// DeleteAirline.js

import React, { useState } from 'react';

const DeleteAirline = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
  });

  const handleDeleteAirline = async () => {
    try {
      const response = await fetch('http://localhost:8080/delete-airline/' + formData.airline_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Airline deleted successfully
        console.log('Airline deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete airline');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting airline:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Airline</h2>
      <form>
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeleteAirline}>
          Delete Airline
        </button>
      </form>
    </div>
  );
};

export default DeleteAirline;
