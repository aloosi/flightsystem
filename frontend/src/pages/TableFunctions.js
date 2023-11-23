import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';
import "./pages.css" ;

function TableFunctions() {
  return (
    <div className='tablefunctions'>
      <h1>Table Functions</h1>
      <div>
            <nav className='options'>
            <ul className='app-ul'>
                <li className="app-li">
                  <Link to="/tablefunctions/delete-tables">Drop Tables</Link>          
                </li>

                <li className="app-li">
                  <Link to="/tablefunctions/create-tables">Create Tables</Link>
                </li>
                
                <li className="app-li">
                  <Link to="/tablefunctions/populate-tables">Populate Tables</Link>
                </li>
            </ul>
            </nav>

            <hr />
        </div>
    </div>
    
  );
}

export default TableFunctions;