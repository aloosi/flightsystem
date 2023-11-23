import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;



function TouristQueries() {
  return (
    <div className='touristqueries'>
      <h1>Tourist Queries</h1>
      <div>
            <nav className='options'>
            <ul className="app-ul">
                <li className="app-li">
                <Link to="./create-tourist">Create Tourist</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-tourist">Delete Tourist</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-tourists">Get All Tourists</Link>
                </li>
                <li className="app-li">
                <Link to="./get-tourist-by-id">Get Tourist By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-tourist">Update Tourist</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default TouristQueries;