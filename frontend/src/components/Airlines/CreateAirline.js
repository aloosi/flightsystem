import React, { useState } from 'react';

const CreateAirline = () => {
  const [formData, setFormData] = useState({
    airline_id: '',
    airline_name: '',
  });
  const [message, setMessage] = useState(null);


  const handleCreateAirline = async () => {
    try {
      console.log('Form Data:', formData); // Log the formData
  
      const response = await fetch('http://3.134.76.216:8080/create-airline', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
  
      if (response.ok) {
        // Airline created successfully
        console.log('Airline created successfully');
        setMessage('Airline created successfully');

        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to create airline. Status:', response.status);
        const errorData = await response.json();
        console.error('Error message:', errorData);
        setMessage('Error message:', errorData);
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating airline:', error.message);
      setMessage('Error creating airline:', error.message);
    }
  };
  

  return (
    <div>
      <h2>Create Airline</h2>
      <form>
        <label>Airline ID:</label>
        <input
          type="text"
          value={formData.airline_id}
          onChange={(e) => setFormData({ ...formData, airline_id: parseInt(e.target.value) })}
        />
        <label>Airline Name:</label>
        <input
          type="text"
          value={formData.airline_name}
          onChange={(e) => setFormData({ ...formData, airline_name: e.target.value })}
        />

        <button type="button" onClick={handleCreateAirline}>
          Create Airline
        </button>
      </form>
    </div>
  );
};

export default CreateAirline;
