import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function FlightQueries() {
  return (
    <div className='flightqueries'>
      <h1>Flight Queries</h1>
      <div>
            <nav>
            <ul className="app-li">
                <li className="app-li">
                <Link to="./create-flight">Create Flight</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-flight">Delete Flight</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-flights">Get All Flights</Link>
                </li>
                <li className="app-li">
                <Link to="./get-flight-by-id">Get Flight By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-flight">Update Flight</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default FlightQueries;