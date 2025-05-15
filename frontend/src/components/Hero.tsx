
import React from 'react';

const Hero = () => {
  return (
    <div className="relative h-[70vh] bg-hero-pattern bg-cover bg-center">
      <div className="absolute inset-0 hero-overlay"></div>
      <div className="relative z-10 flex flex-col items-center justify-center h-full text-white text-center px-4">
        <h1 className="text-4xl md:text-6xl lg:text-7xl font-light mb-4 tracking-tight">
          <span className="font-normal">TRIP</span> PLANNER
        </h1>
        <p className="text-lg md:text-2xl max-w-2xl font-light tracking-wide mb-3">
          ให้ <span className="text-tripOrange font-normal">AI</span> ของเราช่วยวางแผนทริปในฝันของคุณ
        </p>
        <p className="text-base md:text-lg max-w-2xl text-gray-200 mb-8">
          ออกแบบการเดินทางที่สมบูรณ์แบบด้วยคำแนะนำที่ปรับให้เหมาะกับคุณ
        </p>
        
        <div className="flex flex-col sm:flex-row gap-4">
          <button 
            className="apple-button px-8 py-3 text-white rounded-full hover:bg-opacity-90 transition-all text-base shadow-lg"
            onClick={() => document.getElementById('chatSection')?.scrollIntoView({ behavior: 'smooth' })}
          >
            เริ่มการวางแผน
          </button>
          
          <button 
            className="px-8 py-3 bg-transparent border border-white text-white rounded-full hover:bg-white hover:text-tripPurple transition-all text-base"
          >
            ดูตัวอย่าง
          </button>
        </div>
        
        <div className="absolute bottom-10 w-full flex justify-center animate-bounce">
          <svg 
            className="w-6 h-6 text-white" 
            fill="none" 
            stroke="currentColor" 
            viewBox="0 0 24 24" 
            xmlns="http://www.w3.org/2000/svg"
          >
            <path 
              strokeLinecap="round" 
              strokeLinejoin="round" 
              strokeWidth={2} 
              d="M19 14l-7 7m0 0l-7-7m7 7V3"
            />
          </svg>
        </div>
      </div>
    </div>
  );
};

export default Hero;
