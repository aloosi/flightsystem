import React, { useState } from 'react';

const CreatePayment = () => {
  const [formData, setFormData] = useState({
    payment_id: '',
    payment_type: '',
    processing_fees: '',
    payment_status: '',
  });

  const handleCreatePayment = async () => {
    try {
      console.log('Form Data:', formData); // Log the formData
  
      const response = await fetch('http://3.134.76.216:8080/create-payment-method', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
  
      if (response.ok) {
        // Payment created successfully
        console.log('Payment created successfully');
        // You might want to redirect the user or show a success message
      } else {
        // Log the response status and error message
        console.error('Failed to create payment. Status:', response.status);
        const errorData = await response.json();
        console.error('Error message:', errorData);
      }
    } catch (error) {
      // Handle network error or other issues
      console.error('Error creating payment:', error.message);
    }
  };
  

  return (
    <div>
      <h2>Create Payment</h2>
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

        <button type="button" onClick={handleCreatePayment}>
          Create Payment
        </button>
      </form>
    </div>
  );
};

export default CreatePayment;
