import React, { useState } from 'react';

const GetAllAirlines = () => {
  const [airlineData, setAirlineData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllAirlines = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-airlines', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setAirlineData(data); // Set the airline data directly, assuming it's an array of objects
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire airlines. Status: ${response.status}, Error: ${errorData.message}`);
        setAirlineData([]);
      }
    } catch (error) {
      setError(`Error acquiring airlines: ${error.message}`);
      setAirlineData([]);
    }
  };

  return (
    <div>
      <h2>Get All Airlines</h2>
      <form>
        <button type="button" onClick={handleGetAllAirlines}>
          Get All Airlines
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
            <h3>Airlines Data:</h3>
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

export default GetAllAirlines;
