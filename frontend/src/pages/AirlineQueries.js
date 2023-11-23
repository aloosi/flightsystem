import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function AirlineQueries() {
  return (
    <div className='airlinequeries'>
      <h1>Airline Queries</h1>
      <div>
            <nav>
            <ul className='app-ul'>
                <li className='app-li'>
                <Link to="./create-airline">Create Airline</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-airline">Delete Airline</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-airlines">Get All Airlines</Link>
                </li>
                <li className="app-li">
                <Link to="./get-airline-by-id">Get Airline By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-airline">Update Airline</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default AirlineQueries;