
import React from 'react';
import { Mail, Phone, MapPin } from 'lucide-react';

const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white">
      <div className="container mx-auto px-4">
        {/* Main Footer Content */}
        <div className="py-12">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div>
              <h3 className="text-xl font-light mb-6">
                <span className="text-tripPurple font-normal">TRIP</span> PLANNER
              </h3>
              <p className="text-gray-400 font-light leading-relaxed">
                ผู้ช่วยวางแผนการท่องเที่ยวส่วนตัวที่ขับเคลื่อนด้วย AI สำหรับการวางแผนทริปที่สมบูรณ์แบบ
              </p>
            </div>
            
            <div className="lg:ml-auto">
              <h4 className="text-lg font-light mb-6">เกี่ยวกับเรา</h4>
              <ul className="space-y-3">
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">เกี่ยวกับบริษัท</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">ทีมงาน</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">ร่วมงานกับเรา</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">ข่าวสาร</a>
                </li>
              </ul>
            </div>
            
            <div>
              <h4 className="text-lg font-light mb-6">บริการ</h4>
              <ul className="space-y-3">
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">วางแผนทริป</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">ที่พัก</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">กิจกรรม</a>
                </li>
                <li>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors font-light text-sm">ร้านอาหาร</a>
                </li>
              </ul>
            </div>
            
            <div>
              <h4 className="text-lg font-light mb-6">ติดต่อเรา</h4>
              <ul className="space-y-4">
                <li className="flex items-start">
                  <Mail className="w-5 h-5 text-tripPurple mr-3 mt-0.5" />
                  <div>
                    <span className="block text-sm font-light text-gray-400">อีเมล:</span>
                    <a href="mailto:info@tripplanner.com" className="text-gray-300 hover:text-white transition-colors">info@tripplanner.com</a>
                  </div>
                </li>
                <li className="flex items-start">
                  <Phone className="w-5 h-5 text-tripPurple mr-3 mt-0.5" />
                  <div>
                    <span className="block text-sm font-light text-gray-400">โทรศัพท์:</span>
                    <a href="tel:+6622223333" className="text-gray-300 hover:text-white transition-colors">+66 2 222 3333</a>
                  </div>
                </li>
                <li className="flex items-start">
                  <MapPin className="w-5 h-5 text-tripPurple mr-3 mt-0.5" />
                  <div>
                    <span className="block text-sm font-light text-gray-400">ที่อยู่:</span>
                    <span className="text-gray-300">123 ถนนสุขุมวิท กรุงเทพฯ 10110</span>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
        
        {/* Footer Bottom */}
        <div className="border-t border-gray-700 py-6 text-center text-gray-400">
          <p className="text-sm font-light">&copy; {new Date().getFullYear()} Trip Planner. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
