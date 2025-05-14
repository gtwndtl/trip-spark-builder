
import React from 'react';

const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white py-12">
      <div className="container mx-auto px-4">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <h3 className="text-xl font-bold mb-4">
              <span className="text-tripPurple">TRIP</span> PLANNER
            </h3>
            <p className="text-gray-400">
              Your personal AI-powered travel assistant for perfect trip planning.
            </p>
          </div>
          
          <div>
            <h4 className="text-lg font-semibold mb-4">เกี่ยวกับเรา</h4>
            <ul className="space-y-2">
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">เกี่ยวกับบริษัท</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">ทีมงาน</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">ร่วมงานกับเรา</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">ข่าวสาร</a>
              </li>
            </ul>
          </div>
          
          <div>
            <h4 className="text-lg font-semibold mb-4">บริการ</h4>
            <ul className="space-y-2">
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">วางแผนทริป</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">ที่พัก</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">กิจกรรม</a>
              </li>
              <li>
                <a href="#" className="text-gray-400 hover:text-white transition-colors">ร้านอาหาร</a>
              </li>
            </ul>
          </div>
          
          <div>
            <h4 className="text-lg font-semibold mb-4">ติดต่อเรา</h4>
            <ul className="space-y-2">
              <li className="text-gray-400">
                <span className="block">อีเมล:</span>
                <a href="mailto:info@tripplanner.com" className="hover:text-white transition-colors">info@tripplanner.com</a>
              </li>
              <li className="text-gray-400">
                <span className="block">โทรศัพท์:</span>
                <a href="tel:+6622223333" className="hover:text-white transition-colors">+66 2 222 3333</a>
              </li>
              <li className="text-gray-400">
                <span className="block">ที่อยู่:</span>
                <span>123 ถนนสุขุมวิท กรุงเทพฯ 10110</span>
              </li>
            </ul>
          </div>
        </div>
        
        <div className="border-t border-gray-700 mt-8 pt-8 text-center text-gray-400">
          <p>&copy; {new Date().getFullYear()} Trip Planner. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
