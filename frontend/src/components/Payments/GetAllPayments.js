// GetAllPayments.js

import React, { useState } from 'react';

const GetAllPayments = () => {
  const [formData, setFormData] = useState({
    payment_id: '',
  });

  const handleGetAllPayments = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-payment-methods',{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Payment deleted successfully
        console.log('payments acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire payments');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring payments:', error.message);
    }
  };

  return (
    <div>
      <h2>Get All payments</h2>
      <form>
        <button type="button" onClick={handleGetAllPayments}>
          Get All payments
        </button>
      </form>
    </div>
  );
};

export default GetAllPayments;
