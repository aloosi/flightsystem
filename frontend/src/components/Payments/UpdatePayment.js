import React, { useState } from 'react';

const UpdatePayment = () => {
  const [formData, setFormData] = useState({
    payment_id: '',
    payment_type: '',
    processing_fees: '',
    payment_status: '',
  });

  const handleUpdatePayment = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/update-payment-method', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Payment updated successfully
        console.log('Payment updated successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Handle error response
        console.error('Failed to update payment');
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error updating payment:', error.message);
    }
  };

  return (
    <div>
      <h2>Update Payment</h2>
      <form>
      <label>Payment ID:</label>
        <input
          type="text"
          value={formData.payment_id}
          onChange={(e) => setFormData({ ...formData, payment_id: parseInt(e.target.value) })}
        />
        <label>Payment Type:</label>
        <input
          type="text"
          value={formData.payment_type}
          onChange={(e) => setFormData({ ...formData, payment_type: e.target.value })}
        />
        <label>Processing Fees:</label>
        <input
          type="text"
          value={formData.processing_fees}
          onChange={(e) => setFormData({ ...formData, processing_fees: parseInt(e.target.value) })}
        />
        <label>Payment Status:</label>
        <input
          type="text"
          value={formData.payment_status}
          onChange={(e) => setFormData({ ...formData, payment_status: e.target.value })}
        />

        <button type="button" onClick={handleUpdatePayment}>
          Update Payment
        </button>
      </form>
    </div>
  );
};

export default UpdatePayment;
