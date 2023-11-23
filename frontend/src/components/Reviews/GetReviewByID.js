import React, { useState } from 'react';

const GetReviewByID = () => {
  const [reviewData, setReviewData] = useState(null);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({
    review_id: '',
  });

  const handleGetReviewByID = async () => {
    try {
      const response = await fetch(`http://3.134.76.216:8080/get-review-by-id/${formData.review_id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setReviewData(data); // Set the review data
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire review. Status: ${response.status}, Error: ${errorData.message}`);
        setReviewData(null); // Set null in case of error
      }
    } catch (error) {
      setError(`Error acquiring review: ${error.message}`);
      setReviewData(null); // Set null in case of error
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

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        reviewData && (
          <div>
            <h3>Review Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Review ID</th>
                  {/* Add more details as needed */}
                </tr>
              </thead>
              <tbody>
                <tr key={reviewData.review_id}>
                  <td>{reviewData.review_id}</td>
                  {/* Add more details as needed */}
                </tr>
              </tbody>
            </table>
          </div>
        )
      )}
    </div>
  );
};

export default GetReviewByID;
