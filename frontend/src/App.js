import React from 'react';
import { BrowserRouter as Router, Route, Link, Routes} from 'react-router-dom';

import Navbar from './components/Navbar/Navbar';
import Home from './pages/Home';
import TableFunctions from './pages/TableFunctions';
import TouristQueries from './pages/TouristQueries';
import AirlineQueries from './pages/AirlineQueries';
import FlightQueries from './pages/FlightQueries';
import BookingQueries from './pages/BookingQueries';
import PaymentQueries from './pages/PaymentQueries';
import ReviewQueries from './pages/ReviewQueries';

import DeleteTables from './components/Tables/DeleteTables';
import CreateTables from './components/Tables/CreateTables';
import PopulateTables from './components/Tables/PopulateTables';

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
    <>
    <Router>
      <Navbar> </Navbar>
      <Routes>
        <Route path='*' element={<Home />} />
        <Route path='tablefunctions/*' element={ <TableFunctions />} />
        <Route path='touristqueries/*' element={ <TouristQueries />} />
        <Route path='airlinequeries/*' element={ <AirlineQueries />} />
        <Route path='flightqueries/*' element={ <FlightQueries />} />
        <Route path='bookingqueries/*' element={ <BookingQueries />} />
        <Route path='paymentqueries/*' element={ <PaymentQueries />} />
        <Route path='reviewqueries/*' element={ <ReviewQueries />} />

        <Route path="/tablefunctions/delete-tables" element={<DeleteTables />} />
        <Route path="/tablefunctions/create-tables" element={<CreateTables />} />
        <Route path="/tablefunctions/populate-tables" element={<PopulateTables />} />

        <Route path="/touristqueries/create-tourist" element={<CreateTourist />} />
        <Route path="/touristqueries/delete-tourist" element={<DeleteTourist />} />
        <Route path="/touristqueries/get-all-tourists" element={<GetAllTourists />} />
        <Route path="/touristqueries/get-tourist-by-id" element={<GetTouristByID />} />
        <Route path="/touristqueries/update-tourist" element={<UpdateTourist />} />

        <Route path="/airlinequeries/create-airline" element={<CreateAirline />} />
        <Route path="/airlinequeries/delete-airline" element={<DeleteAirline />} />
        <Route path="/airlinequeries/get-all-airlines" element={<GetAllAirlines />} />
        <Route path="/airlinequeries/get-airline-by-id" element={<GetAirlineByID />} />
        <Route path="/airlinequeries/update-airline" element={<UpdateAirline />} />

        <Route path="/flightqueries/create-flight" element={<CreateFlight />} />
        <Route path="/flightqueries/delete-flight" element={<DeleteFlight />} />
        <Route path="/flightqueries/get-all-flights" element={<GetAllFlights />} />
        <Route path="/flightqueries/get-flight-by-id" element={<GetFlightByID />} />
        <Route path="/flightqueries/update-flight" element={<UpdateFlight />} />

        <Route path="/bookingqueries/create-booking" element={<CreateBooking />} />
        <Route path="/bookingqueries/delete-booking" element={<DeleteBooking />} />
        <Route path="/bookingqueries/get-all-bookings" element={<GetAllBookings />} />
        <Route path="/bookingqueries/get-booking-by-id" element={<GetBookingByID />} />
        <Route path="/bookingqueries/update-booking" element={<UpdateBooking />} />

        <Route path="/reviewqueries/create-review" element={<CreateReview />} />
        <Route path="/reviewqueries/delete-review" element={<DeleteReview />} />
        <Route path="/reviewqueries/get-all-reviews" element={<GetAllReviews />} />
        <Route path="/reviewqueries/get-review-by-id" element={<GetReviewByID />} />
        <Route path="/reviewqueries/update-review" element={<UpdateReview />} />

        <Route path="/paymentqueries/create-payment" element={<CreatePayment />} />
        <Route path="/paymentqueries/delete-payment" element={<DeletePayment />} />
        <Route path="/paymentqueries/get-all-payments" element={<GetAllPayments />} />
        <Route path="/paymentqueries/get-payment-by-id" element={<GetPaymentByID />} />
        <Route path="/paymentqueries/update-payment" element={<UpdatePayment />} />
      </Routes>
    </Router>
    </>
  );
}

export default App;
