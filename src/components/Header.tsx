
import React from 'react';

const Header = () => {
  return (
    <header className="flex items-center justify-between py-4 px-6 bg-white shadow-sm">
      <div className="flex items-center">
        <h1 className="text-2xl font-bold text-tripDark">
          <span className="text-tripPurple">TRIP</span> PLANNER
        </h1>
      </div>
      <nav>
        <ul className="flex space-x-6">
          <li>
            <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors">
              Home
            </a>
          </li>
          <li>
            <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors">
              About
            </a>
          </li>
          <li>
            <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors">
              Contact
            </a>
          </li>
          <li>
            <a 
              href="#" 
              className="px-4 py-2 bg-tripPurple text-white rounded-md hover:bg-opacity-90 transition-colors"
            >
              Sign Up
            </a>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
