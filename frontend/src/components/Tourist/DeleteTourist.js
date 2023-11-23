// DeleteTourist.js

import React, { useState } from 'react';

const DeleteTourist = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
  });
  const [message, setMessage] = useState(null);

  const handleDeleteTourist = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/delete-tourist/' + formData.tourist_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Tourist deleted successfully
        console.log('Tourist deleted successfully');
        setMessage('Tourist deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete tourist');
        setMessage('Failed to delete tourist');

      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting tourist:', error.message);
      setMessage('Error deleting tourist:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Tourist</h2>
      <form>
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeleteTourist}>
          Delete Tourist
        </button>
      </form>
      <hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default DeleteTourist;
