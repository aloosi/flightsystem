// GetAllTourists.js

import React, { useState } from 'react';

const GetAllTourists = () => {
  const [formData, setFormData] = useState({
    tourist_id: '',
  });

  const handleGetAllTourists = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-tourists',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Tourist deleted successfully
        console.log('Tourists acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire tourists');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring tourists:', error.message);
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
    </div>
  );
};

export default GetAllTourists;
