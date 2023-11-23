import React, { useState } from 'react';

const GetAllReviews = () => {
  const [reviewsData, setReviewsData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllReviews = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-reviews', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setReviewsData(data); // Set the reviews data
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire reviews. Status: ${response.status}, Error: ${errorData.message}`);
        setReviewsData([]); // Set an empty array in case of error
      }
    } catch (error) {
      setError(`Error acquiring reviews: ${error.message}`);
      setReviewsData([]); // Set an empty array in case of error
    }
  };

  return (
    <div>
      <h2>Get All Reviews</h2>
      <form>
        <button type="button" onClick={handleGetAllReviews}>
          Get All Reviews
        </button>
      </form>

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        reviewsData.length > 0 && (
          <div>
            <h3>Reviews Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Review ID</th>
                  {/* Add more details as needed */}
                </tr>
              </thead>
              <tbody>
                {reviewsData.map((review) => (
                  <tr key={review.review_id}>
                    <td>{review.review_id}</td>
                    {/* Add more details as needed */}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )
      )}
    </div>
  );
};

export default GetAllReviews;
