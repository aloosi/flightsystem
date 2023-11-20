import React, { useState } from 'react';

const UpdateTourist = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
    // Include other fields you want to update
    // For example:
    tourist_first_name: '',
    tourist_last_name: '',
    tourist_email: '',
    // Add more fields as needed
  });

  const handleUpdateTourist = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/update-tourist', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Tourist updated successfully
        console.log('Tourist updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update tourist');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating tourist:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Tourist</h2>
      <form>
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value, 10) || 0 })}
        />

        {/* Include other input fields for the data you want to update */}
        <label>First Name:</label>
        <input
          type="text"
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, tourist_first_name: e.target.value })}
        />

        <label>Last Name:</label>
        <input
          type="text"
          value={formData.country}
          onChange={(e) => setFormData({ ...formData, tourist_last_name: e.target.value })}
        />

        <label>Email:</label>
        <input
          type="text"
          value={formData.country}
          onChange={(e) => setFormData({ ...formData, tourist_email: e.target.value })}
        />

        {/* Add more input fields as needed for other tourist information */}

        <button type="button" onClick={handleUpdateTourist}>
          Update Tourist
        </button>
      </form>
    </div>
  );
};

export default UpdateTourist;
