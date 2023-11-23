import React, { useState } from 'react';

const GetAirlineByID = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
  });
  const [airlineData, setAirlineData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAirlineByID = async () => {
    try {
      const response = await fetch(`http://3.134.76.216:8080/get-airline-by-id/${formData.airline_id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setAirlineData([data]); // Wrap the data in an array to match the structure of GetAllAirlines
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire airline. Status: ${response.status}, Error: ${errorData.message}`);
        setAirlineData([]);
      }
    } catch (error) {
      setError(`Error acquiring airline: ${error.message}`);
      setAirlineData([]);
    }
  };

  return (
    <div>
      <h2>Get Airline By ID</h2>
      <form>
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value, 10) || '' })}
        />

        <button type="button" onClick={handleGetAirlineByID}>
          Get Airline By ID
        </button>
      </form>
<hr />

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        airlineData.length > 0 && (
          <div>
            <h3>Airline Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Airline ID</th>
                  <th>Airline Name</th>
                  {/* Add more columns as needed */}
                </tr>
              </thead>
              <tbody>
                {airlineData.map((airline) => (
                  <tr key={airline.airline_id}>
                    <td>{airline.airline_id}</td>
                    <td>{airline.airline_name}</td>
                    {/* Add more cells for additional data */}
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

export default GetAirlineByID;
