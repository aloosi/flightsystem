import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes } from 'react-router-dom';
import CreateTourist from './components/Tourist/CreateTourist';
import DeleteTourist from './components/Tourist/DeleteTourist';
import GetAllTourists from './components/Tourist/GetAllTourists';
import GetTouristByID from './components/Tourist/GetTouristByID';
import UpdateTourist from './components/Tourist/UpdateTourist';

import CreateFlight from './components/Flights/CreateFlight';
import DeleteFlight from './components/Flights/DeleteFlight';
import GetAllFlights from './components/Flights/GetAllFlights';
import GetFlightByID from './components/Flights/GetFlightByID';
import UpdateFlight from './components/Flights/UpdateFlight';

import CreateAirline from './components/Airlines/CreateAirline';
import DeleteAirline from './components/Airlines/DeleteAirline';
import GetAllAirlines from './components/Airlines/GetAllAirlines';
import GetAirlineByID from './components/Airlines/GetAirlineByID';
import UpdateAirline from './components/Airlines/UpdateAirline';

import CreateBooking from './components/Bookings/CreateBooking';
import DeleteBooking from './components/Bookings/DeleteBooking';
import GetAllBookings from './components/Bookings/GetAllBookings';
import GetBookingByID from './components/Bookings/GetBookingByID';
import UpdateBooking from './components/Bookings/UpdateBooking';

import CreateReview from './components/Reviews/CreateReview';
import DeleteReview from './components/Reviews/DeleteReview';
import GetAllReviews from './components/Reviews/GetAllReviews';
import GetReviewByID from './components/Reviews/GetReviewByID';
import UpdateReview from './components/Reviews/UpdateReview';

import CreatePayment from './components/Payments/CreatePayment';
import DeletePayment from './components/Payments/DeletePayment';
import GetAllPayments from './components/Payments/GetAllPayments';
import GetPaymentByID from './components/Payments/GetPaymentByID';
import UpdatePayment from './components/Payments/UpdatePayment';

const App = () => {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/tourist/create-tourist">Create Tourist</Link>
              <Link to="/tourist/delete-tourist">Delete Tourist</Link>
              <Link to="/tourist/get-all-tourists">Get All Tourists</Link>
              <Link to="/tourist/get-tourist-by-id">Get Tourist By ID</Link>
              <Link to="/tourist/update-tourist">Update Tourist</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/tourist/create-tourist" element={<CreateTourist />} />
        <Route path="/tourist/delete-tourist" element={<DeleteTourist />} />
        <Route path="/tourist/get-all-tourists" element={<GetAllTourists />} />
        <Route path="/tourist/get-tourist-by-id" element={<GetTouristByID />} />
        <Route path="/tourist/update-tourist" element={<UpdateTourist />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
      </div>

      <div>
        <nav>
          <ul>
            <li>
              <Link to="/flights/create-flight">Create Flight</Link>
              <Link to="/flights/delete-flight">Delete Flight</Link>
              <Link to="/flights/get-all-flights">Get All Flights</Link>
              <Link to="/flights/get-flight-by-id">Get Flight By ID</Link>
              <Link to="/flights/update-flight">Update Flight</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/flights/create-flight" element={<CreateFlight />} />
        <Route path="/flights/delete-flight" element={<DeleteFlight />} />
        <Route path="/flights/get-all-flights" element={<GetAllFlights />} />
        <Route path="/flights/get-flight-by-id" element={<GetFlightByID />} />
        <Route path="/flights/update-flight" element={<UpdateFlight />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
        </div>
        <div>
        <nav>
          <ul>
            <li>
              <Link to="/airlines/create-airline">Create Airline</Link>
              <Link to="/airlines/delete-airline">Delete Airline</Link>
              <Link to="/airlines/get-all-airlines">Get All Airlines</Link>
              <Link to="/airlines/get-airline-by-id">Get Airline By ID</Link>
              <Link to="/airlines/update-airline">Update Airline</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/airlines/create-airline" element={<CreateAirline />} />
        <Route path="/airlines/delete-airline" element={<DeleteAirline />} />
        <Route path="/airlines/get-all-airlines" element={<GetAllAirlines />} />
        <Route path="/airlines/get-airline-by-id" element={<GetAirlineByID />} />
        <Route path="/airlines/update-airline" element={<UpdateAirline />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
      </div>

      <div>
        <nav>
          <ul>
            <li>
              <Link to="/bookings/create-booking">Create Booking</Link>
              <Link to="/bookings/delete-booking">Delete Booking</Link>
              <Link to="/bookings/get-all-bookings">Get All Bookings</Link>
              <Link to="/bookings/get-booking-by-id">Get Booking By ID</Link>
              <Link to="/bookings/update-booking">Update Booking</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/bookings/create-booking" element={<CreateBooking />} />
        <Route path="/bookings/delete-booking" element={<DeleteBooking />} />
        <Route path="/bookings/get-all-bookings" element={<GetAllBookings />} />
        <Route path="/bookings/get-booking-by-id" element={<GetBookingByID />} />
        <Route path="/bookings/update-booking" element={<UpdateBooking />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
      </div>


      <div>
        <nav>
          <ul>
            <li>
              <Link to="/reviews/create-review">Create Review</Link>
              <Link to="/reviews/delete-review">Delete Review</Link>
              <Link to="/reviews/get-all-reviews">Get All Reviews</Link>
              <Link to="/reviews/get-review-by-id">Get Review By ID</Link>
              <Link to="/reviews/update-review">Update Review</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/reviews/create-review" element={<CreateReview />} />
        <Route path="/reviews/delete-review" element={<DeleteReview />} />
        <Route path="/reviews/get-all-reviews" element={<GetAllReviews />} />
        <Route path="/reviews/get-review-by-id" element={<GetReviewByID />} />
        <Route path="/reviews/update-review" element={<UpdateReview />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
      </div>

      <div>
        <nav>
          <ul>
            <li>
              <Link to="/payments/create-payment">Create Payment</Link>
              <Link to="/payments/delete-payment">Delete Payment</Link>
              <Link to="/payments/get-all-payments">Get All Payments</Link>
              <Link to="/payments/get-payment-by-id">Get Payment By ID</Link>
              <Link to="/payments/update-payment">Update Payment</Link>


            </li>
            {/* Add links for other CRUD operations */}
          </ul>
        </nav>

        <hr />
        <Routes>
        <Route path="/payments/create-payment" element={<CreatePayment />} />
        <Route path="/payments/delete-payment" element={<DeletePayment />} />
        <Route path="/payments/get-all-payments" element={<GetAllPayments />} />
        <Route path="/payments/get-payment-by-id" element={<GetPaymentByID />} />
        <Route path="/payments/update-payment" element={<UpdatePayment />} />



        {/* Add routes for other CRUD operations */}
        </Routes>
      </div>
    </Router>
  );
};

export default App;
