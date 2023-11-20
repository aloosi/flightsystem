// DeletePayment.js

import React, { useState } from 'react';

const DeletePayment = () => {
  const [formData, setFormData] = useState({
    payment_id: '',
  });

  const handleDeletePayment = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/delete-payment-method/' + formData.payment_id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Payment deleted successfully
        console.log('Payment deleted successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to delete payment');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error deleting payment:', error.message);
    }
  };

  return (
    <div>
      <h2>Delete Payment</h2>
      <form>
        <label>Payment ID:</label>
        <input
          type="text"
          value={formData.payment_id}
          onChange={(e) => setFormData({ ...formData, payment_id: parseInt(e.target.value, 10) || 0 })}
        />

        <button type="button" onClick={handleDeletePayment}>
          Delete Payment
        </button>
      </form>
    </div>
  );
};

export default DeletePayment;
