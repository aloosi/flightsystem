// GetPaymentByID.js

import React, { useState } from 'react';

const GetPaymentByID = () => {
  const [formData, setFormData] = useState({
    payment_id: '',
  });

  const handleGetPaymentByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-payment-method-by-id/' + formData.payment_id,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        // Payment deleted successfully
        console.log('Payment acquired successfully', data);
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to acquire payment');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error acquiring payment:', error.message);
    }
  };

  return (
    <div>
      <h2>Get Payment By ID</h2>
      <form>
        <label>Payment ID:</label>
        <input
          type="text"
          value={formData.payment_id}
          onChange={(e) => setFormData({ ...formData, payment_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleGetPaymentByID}>
          Get Payment By ID
        </button>
      </form>
    </div>
  );
};

export default GetPaymentByID;
