import React, { useState } from 'react';

const UpdateReview = () => {
  const [formData, setFormData] = useState({
    review_id: '',
    review_rating: '',
    tourist_id: '',
    flight_id: '',
  });
  const [message, setMessage] = useState(null);


  const handleUpdateReview = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/update-review', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Review updated successfully
        console.log('Review updated successfully');
        setMessage('Review updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update review');
        setMessage('Failed to update review');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating review:', error.message);
      setMessage('Error updating review:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Review</h2>
      <form>
      <label>Review ID:</label>
        <input
          type="text"
          value={formData.review_id}
          onChange={(e) => setFormData({ ...formData, review_id: parseInt(e.target.value) })}
        />
        <label>Review Rating:</label>
        <input
          type="text"
          value={formData.review_rating}
          onChange={(e) => setFormData({ ...formData, review_rating: parseInt(e.target.value) })}
        />

        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value) })}
        />

        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value) })}
        />

        <button type="button" onClick={handleUpdateReview}>
          Update Review
        </button>
      </form>
      <hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default UpdateReview;
