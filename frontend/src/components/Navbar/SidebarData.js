import React from 'react'
import * as FaIcons from "react-icons/fa";
import * as Fa6Icons from "react-icons/fa6";
import * as AiIcons from "react-icons/ai";
import * as IoIcons from "react-icons/io";
import { MdAirlines } from "react-icons/md";
import { MdOutlinePayment } from "react-icons/md";
import { MdOutlineReviews } from "react-icons/md";


export const SidebarData = [
    {
        title: 'Home',
        path: '/',
        icon: <AiIcons.AiFillHome />,
        cName: 'nav-text'
    },
    {
        title: 'Table Functions',
        path: '/tablefunctions',
        icon: <FaIcons.FaTable />,
        cName: 'nav-text'
    },
    {
        title: 'Tourist Queries',
        path: '/touristqueries',
        icon: <Fa6Icons.FaPerson />,
        cName: 'nav-text'
    },
    {
        title: 'Airline Queries',
        path: '/airlinequeries',
        icon: <MdAirlines />,
        cName: 'nav-text'
    },
    {
        title: 'Flight Queries',
        path: '/flightqueries',
        icon: <IoIcons.IoIosAirplane />,
        cName: 'nav-text'
    },
    {
        title: 'Booking Queries',
        path: '/bookingqueries',
        icon: <FaIcons.FaBook />,
        cName: 'nav-text'
    },
    {
        title: 'Payments Queries',
        path: '/paymentqueries',
        icon: <MdOutlinePayment />,
        cName: 'nav-text'
    },
    {
        title: 'Reviews Queries',
        path: '/',
        icon: <MdOutlineReviews />,
        cName: 'nav-text'
    }
];