import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function BookingQueries() {
  return (
    <div className='bookingqueries'>
      <h1>Booking Queries</h1>
      <div>
            <nav>
            <ul className="app-ul">
                <li className="app-li">
                <Link to="./create-booking">Create Booking</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-booking">Delete Booking</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-bookings">Get All Bookings</Link>
                </li>
                <li className="app-li">
                <Link to="./get-booking-by-id">Get Booking By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-booking">Update Booking</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default BookingQueries;