import React, { useState } from 'react';

const GetPaymentByID = () => {
  const [paymentData, setPaymentData] = useState(null);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({
    payment_id: '',
  });

  const handleGetPaymentByID = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-payment-method-by-id/' + formData.payment_id, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setPaymentData(data); // Set the payment data
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire payment. Status: ${response.status}, Error: ${errorData.message}`);
        setPaymentData(null); // Set null in case of error
      }
    } catch (error) {
      setError(`Error acquiring payment: ${error.message}`);
      setPaymentData(null); // Set null in case of error
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
          onChange={(e) => setFormData({ ...formData, payment_id: parseInt(e.target.value, 10) || '' })}
        />

        <button type="button" onClick={handleGetPaymentByID}>
          Get Payment By ID
        </button>
      </form>

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        paymentData && (
          <div>
            <h3>Payment Data:</h3>
            <table>
              <thead>
                <tr>
                  <th>Payment ID</th>
                  <th>Payment Method</th>
                  {/* Add more details as needed */}
                </tr>
              </thead>
              <tbody>
                <tr key={paymentData.payment_id}>
                  <td>{paymentData.payment_id}</td>
                  <td>{paymentData.payment_method}</td>
                  {/* Add more details as needed */}
                </tr>
              </tbody>
            </table>
          </div>
        )
      )}
    </div>
  );
};

export default GetPaymentByID;
