import React, { useState } from 'react';

const UpdateAirline = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
    airline_name: '',
  });

  const handleUpdateAirline = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/update-airline', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Airline updated successfully
        console.log('Airline updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update airline');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating airline:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Airline</h2>
      <form>
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value) })}
        />
        <label>Airline Name:</label>
        <input
          type="text"
          value={formData.airline_name}
          onChange={(e) => setFormData({ ...formData, airline_name: e.target.value })}
        />


        <button type="button" onClick={handleUpdateAirline}>
          Update Airline
        </button>
      </form>
    </div>
  );
};

export default UpdateAirline;
