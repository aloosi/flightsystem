import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function ReviewQueries() {
  return (
    <div className='reviewqueries'>
      <h1>Review Queries</h1>
      <div>
            <nav>
            <ul className='app-ul'>
                <li className='app-li'>
                <Link to="./create-review">Create Review</Link>
                </li>
                <li className="app-li">
                <Link to="./delete-review">Delete Review</Link>
                </li>
                <li className="app-li">
                <Link to="./get-all-reviews">Get All Reviews</Link>
                </li>
                <li className="app-li">
                <Link to="./get-review-by-id">Get Review By ID</Link>
                </li>
                <li className="app-li">
                <Link to="./update-review">Update Review</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default ReviewQueries;