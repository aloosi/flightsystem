import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function PaymentQueries() {
  return (
    <div className='paymentqueries'>
      <h1>Payment Queries</h1>
      <div>
            <nav>
            <ul className="app-ul">
                <li className="app-li">
                <Link to="./create-payment">Create Payment</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-payment">Delete Payment</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-payments">Get All Payments</Link>
                </li>
                <li className="app-li">
                <Link to="./get-payment-by-id">Get Payment By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-payment">Update Payment</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default PaymentQueries;