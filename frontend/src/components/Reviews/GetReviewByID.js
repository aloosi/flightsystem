// GetReviewByID.js

import React, { useState } from 'react';

const GetReviewByID = () => {
  const [formData, setFormData] = useState({
    review_id: '',
  });

  const handleGetReviewByID = async () => {
    try {
      const response = await fetch('http://localhost:8080/get-review-by-id/' + formData.review_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Review deleted successfully
        console.log('Review acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire review');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring review:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Review By ID</h2>
      <form>
        <label>Review ID:</label>
        <input
          type="text"
          value={formData.review_id}
          onChange={(e) => setFormData({ ...formData, review_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetReviewByID}>
          Get Review By ID
        </button>
      </form>
    </div>
  );
};

export default GetReviewByID;
