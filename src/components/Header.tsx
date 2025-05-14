
import React, { useState } from 'react';
import { Menu, X } from 'lucide-react';

const Header = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };

  return (
    <header className="sticky top-0 z-50 backdrop-blur-lg bg-white/70 shadow-sm">
      <div className="container mx-auto px-4 py-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center">
            <h1 className="text-2xl font-semibold text-tripDark">
              <span className="text-tripPurple">TRIP</span> PLANNER
            </h1>
          </div>

          {/* Desktop Navigation */}
          <nav className="hidden md:block">
            <ul className="flex space-x-8">
              <li>
                <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors font-light text-base">
                  หน้าแรก
                </a>
              </li>
              <li>
                <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors font-light text-base">
                  เกี่ยวกับเรา
                </a>
              </li>
              <li>
                <a href="#" className="text-gray-600 hover:text-tripPurple transition-colors font-light text-base">
                  ติดต่อ
                </a>
              </li>
              <li>
                <a 
                  href="#" 
                  className="apple-button px-4 py-2 text-white rounded-full hover:bg-opacity-90 transition-all"
                >
                  เริ่มต้นใช้งาน
                </a>
              </li>
            </ul>
          </nav>

          {/* Mobile Menu Button */}
          <button 
            className="md:hidden text-gray-600 focus:outline-none"
            onClick={toggleMobileMenu}
          >
            {isMobileMenuOpen ? (
              <X size={24} />
            ) : (
              <Menu size={24} />
            )}
          </button>
        </div>
      </div>

      {/* Mobile Navigation */}
      <div className={`mobile-menu ${isMobileMenuOpen ? 'active' : ''} md:hidden`}>
        <div className="absolute top-4 right-4">
          <button 
            onClick={toggleMobileMenu}
            className="text-gray-600 focus:outline-none"
          >
            <X size={24} />
          </button>
        </div>
        <ul className="p-4">
          <li className="py-2">
            <a 
              href="#" 
              className="text-lg text-gray-800 hover:text-tripPurple transition-colors"
              onClick={toggleMobileMenu}
            >
              หน้าแรก
            </a>
          </li>
          <li className="py-2">
            <a 
              href="#" 
              className="text-lg text-gray-800 hover:text-tripPurple transition-colors"
              onClick={toggleMobileMenu}
            >
              เกี่ยวกับเรา
            </a>
          </li>
          <li className="py-2">
            <a 
              href="#" 
              className="text-lg text-gray-800 hover:text-tripPurple transition-colors"
              onClick={toggleMobileMenu}
            >
              ติดต่อ
            </a>
          </li>
          <li className="py-4">
            <a 
              href="#" 
              className="apple-button px-6 py-3 text-white rounded-full block text-center"
              onClick={toggleMobileMenu}
            >
              เริ่มต้นใช้งาน
            </a>
          </li>
        </ul>
      </div>
    </header>
  );
};

export default Header;
