import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes, Switch } from 'react-router-dom';

import CreateFlight from '../components/Flights/CreateFlight';
import DeleteFlight from '../components/Flights/DeleteFlight';
import GetAllFlights from '../components/Flights/GetAllFlights';
import GetFlightByID from '../components/Flights/GetFlightByID';
import UpdateFlight from '../components/Flights/UpdateFlight';

import CreateAirline from '../components/Airlines/CreateAirline';
import DeleteAirline from '../components/Airlines/DeleteAirline';
import GetAllAirlines from '../components/Airlines/GetAllAirlines';
import GetAirlineByID from '../components/Airlines/GetAirlineByID';
import UpdateAirline from '../components/Airlines/UpdateAirline';

import CreateBooking from '../components/Bookings/CreateBooking';
import DeleteBooking from '../components/Bookings/DeleteBooking';
import GetAllBookings from '../components/Bookings/GetAllBookings';
import GetBookingByID from '../components/Bookings/GetBookingByID';
import UpdateBooking from '../components/Bookings/UpdateBooking';

import CreateReview from '../components/Reviews/CreateReview';
import DeleteReview from '../components/Reviews/DeleteReview';
import GetAllReviews from '../components/Reviews/GetAllReviews';
import GetReviewByID from '../components/Reviews/GetReviewByID';
import UpdateReview from '../components/Reviews/UpdateReview';

import CreatePayment from '../components/Payments/CreatePayment';
import DeletePayment from '../components/Payments/DeletePayment';
import GetAllPayments from '../components/Payments/GetAllPayments';
import GetPaymentByID from '../components/Payments/GetPaymentByID';
import UpdatePayment from '../components/Payments/UpdatePayment';
import "./pages.css" ;



function Home() {
    return (
    <>
    <div>
        <h2>CPS510 A9: Flightsystem </h2>
    </div>
        
        </>
    );
}

export default Home