// DeleteReview.js

import React, { useState } from 'react';

const DeleteReview = () => {
  const [formData, setFormData] = useState({
    review_id: '',
  });

  const handleDeleteReview = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/delete-review/' + formData.review_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Review deleted successfully
        console.log('Review deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete review');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting review:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Review</h2>
      <form>
        <label>Review ID:</label>
        <input
          type="text"
          value={formData.review_id}
          onChange={(e) => setFormData({ ...formData, review_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeleteReview}>
          Delete Review
        </button>
      </form>
    </div>
  );
};

export default DeleteReview;
