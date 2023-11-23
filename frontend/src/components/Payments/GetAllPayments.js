import React, { useState } from 'react';

const GetAllPayments = () => {
  const [paymentData, setPaymentData] = useState([]);
  const [error, setError] = useState(null);

  const handleGetAllPayments = async () => {
    try {
      const response = await fetch('http://3.134.76.216:8080/get-all-payment-methods', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.ok) {
        const data = await response.json();
        setPaymentData(data); // Set the payment data directly
        setError(null);
      } else {
        const errorData = await response.json();
        setError(`Failed to acquire payments. Status: ${response.status}, Error: ${errorData.message}`);
        setPaymentData([]); // Set an empty array in case of error
      }
    } catch (error) {
      setError(`Error acquiring payments: ${error.message}`);
      setPaymentData([]); // Set an empty array in case of error
    }
  };

  return (
    <div>
      <h2>Get All Payments</h2>
      <form>
        <button type="button" onClick={handleGetAllPayments}>
          Get All Payments
        </button>
      </form>

      {error ? (
        <div>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      ) : (
        paymentData.length > 0 && (
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
                {paymentData.map((payment) => (
                  <tr key={payment.payment_id}>
                    <td>{payment.payment_id}</td>
                    <td>{payment.payment_method}</td>
                    {/* Add more details as needed */}
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

export default GetAllPayments;
