// GetAllReviews.js

import React, { useState } from 'react';

const GetAllReviews = () => {
  const [formData, setFormData] = useState({
    review_id: '',
  });

  const handleGetAllReviews = async () => {
    try {
      const response = await fetch('http://localhost:8080/get-all-reviews',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Review deleted successfully
        console.log('reviews acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire reviews');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring reviews:', error.message);
    }
  };

  return (
    <div>
      <h2>Get All reviews</h2>
      <form>
        <button type="button" onClick={handleGetAllReviews}>
          Get All reviews
        </button>
      </form>
    </div>
  );
};

export default GetAllReviews;
