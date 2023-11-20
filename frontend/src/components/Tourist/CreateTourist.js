// CreateTourist.js

import React, { useState } from 'react';

const CreateTourist = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
    tourist_first_name: '',
    tourist_last_name: '',
    tourist_email: '',
  });

  const handleCreateTourist = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/create-tourist', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Tourist created successfully
        console.log('Tourist created successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to create tourist');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating tourist:', error.message);
    }
  };

  return (
    <div>
      <h2>Create Tourist</h2>
      <form>
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value, 10) || 0 })}
        />
        <label>First Name:</label>
        <input
          type="text"
          value={formData.tourist_first_name}
          onChange={(e) => setFormData({ ...formData, tourist_first_name: e.target.value })}
        />

        <label>Last Name:</label>
        <input
          type="text"
          value={formData.tourist_last_name}
          onChange={(e) => setFormData({ ...formData, tourist_last_name: e.target.value })}
        />

        <label>Email:</label>
        <input
          type="email"
          value={formData.tourist_email}
          onChange={(e) => setFormData({ ...formData, tourist_email: e.target.value })}
        />

        <button type="button" onClick={handleCreateTourist}>
          Create Tourist
        </button>
      </form>
    </div>
  );
};

export default CreateTourist;
