import React, { useState } from 'react';

const GetTouristByID = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
  });
  const [touristData, setTouristData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetTouristByID = async () => {
    try {
      const response = await fetch(`http://3.134.76.216:8080/get-tourist-by-id/${formData.tourist_id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setTouristData([data]); // Wrap the data in an array to match the structure of GetAllTourists
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire tourist. Status: ${response.status}, Error: ${errorData.message}`);
        setTouristData([]);
      }
    } catch (error) {
      setError(`Error acquiring tourist: ${error.message}`);
      setTouristData([]);
    }
  };

  return (
    <div>
      <h2>Get Tourist By ID</h2>
      <form>
        <label>Tourist ID:</label>
        <input
          type="text"
          value={formData.tourist_id}
          onChange={(e) => setFormData({ ...formData, tourist_id: parseInt(e.target.value, 10) || '' })}
        />

        <button type="button" onClick={handleGetTouristByID}>
          Get Tourist By ID
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

export default GetTouristByID;
