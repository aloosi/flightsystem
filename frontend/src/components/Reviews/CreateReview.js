import React, { useState } from 'react';

const CreateReview = () => {
  const [formData, setFormData] = useState({
    review_id: '',
    review_rating: '',
    tourist_id: '',
    flight_id: '',
  });
  const [message, setMessage] = useState(null);


  const handleCreateReview = async () => {
    try {
      console.log('Form Data:', formData); // Log the formData
  
      const response = await fetch('http://3.134.76.216:8080/create-review', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
  
      if (response.ok) {
        // Review created successfully
        console.log('Review created successfully');
        setMessage('Review created successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to create review. Status:', response.status);
        const errorData = await response.json();
        console.error('Error message:', errorData);
        setMessage('Error message:', errorData);
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating review:', error.message);
      setMessage('Error creating review:', error.message);
    }
  };
  

  return (
    <div>
      <h2>Create Review</h2>
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


        <button type="button" onClick={handleCreateReview}>
          Create Review
        </button>
      </form>
<hr />
      {message && <p>{message}</p>}
    </div>
  );
};

export default CreateReview;
