import React, { useState } from 'react';

const GetAllTourists = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
  });
  const [touristData, setTouristData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllTourists = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-tourists', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setTouristData(data);
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire tourists. Status: ${response.status}, Error: ${errorData.message}`);
        setTouristData([]);
      }
    } catch (error) {
      setError(`Error acquiring tourists: ${error.message}`);
      setTouristData([]);
    }
  };

  return (
    <div>
      <h2>Get All Tourists</h2>
      <form>
        <button type="button" onClick={handleGetAllTourists}>
          Get All Tourists
        </button>
      </form>
      <hr />
      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        touristData.length > 0 && (
          <div>
            <h3>Tourist Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Tourist ID</th>
                  <th>First Name</th>
                  <th>Last Name</th>
                  <th>Email</th>
                </tr>
              </thead>
              <tbody>
                {touristData.map((tourist) => (
                  <tr key={tourist.tourist_id}>
                    <td>{tourist.tourist_id}</td>
                    <td>{tourist.tourist_first_name}</td>
                    <td>{tourist.tourist_last_name}</td>
                    <td>{tourist.tourist_email}</td>
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

export default GetAllTourists;
