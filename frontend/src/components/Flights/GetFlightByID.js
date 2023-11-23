import React, { useState } from 'react';

const GetFlightByID = () => {
  const [flightData, setFlightData] = useState(null);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({
    flight_id: '',
  });

  const handleGetFlightByID = async () => {
    try {
      const response = await fetch(`http://3.134.76.216:8080/get-flight-by-id/${formData.flight_id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setFlightData(data); // Set the flight data directly
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire flight. Status: ${response.status}, Error: ${errorData.message}`);
        setFlightData(null);
      }
    } catch (error) {
      setError(`Error acquiring flight: ${error.message}`);
      setFlightData(null);
    }
  };

  return (
    <div>
      <h2>Get Flight By ID</h2>
      <form>
        <label>Flight ID:</label>
        <input
          type="text"
          value={formData.flight_id}
          onChange={(e) => setFormData({ ...formData, flight_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetFlightByID}>
          Get Flight By ID
        </button>
      </form>
<hr />      

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        flightData && (
          <div>
            <h3>Flight Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Flight ID</th>
                  <th>Departure City</th>
                  <th>Destination City</th>
                  {/* Add more columns as needed */}
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>{flightData.flight_id}</td>
                  <td>{flightData.departure_city}</td>
                  <td>{flightData.destination_city}</td>
                  {/* Add more cells for additional data */}
                </tr>
              </tbody>
            </table>
          </div>
        )
      )}
    </div>
  );
};

export default GetFlightByID;
