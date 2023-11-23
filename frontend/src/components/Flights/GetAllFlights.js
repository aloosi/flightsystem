import React, { useState } from 'react';

const GetAllFlights = () => {
  const [flightData, setFlightData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllFlights = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-flights', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setFlightData(data); // Set the flight data directly, assuming it's an array of objects
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire flights. Status: ${response.status}, Error: ${errorData.message}`);
        setFlightData([]);
      }
    } catch (error) {
      setError(`Error acquiring flights: ${error.message}`);
      setFlightData([]);
    }
  };

  return (
    <div>
      <h2>Get All Flights</h2>
      <form>
        <button type="button" onClick={handleGetAllFlights}>
          Get All Flights
        </button>
      </form>
<hr />

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        flightData.length > 0 && (
          <div>
            <h3>Flights Data:</h3>
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
                {flightData.map((flight) => (
                  <tr key={flight.flight_id}>
                    <td>{flight.flight_id}</td>
                    <td>{flight.departure_city}</td>
                    <td>{flight.destination_city}</td>
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

export default GetAllFlights;
