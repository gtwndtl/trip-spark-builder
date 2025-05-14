
import React from 'react';

const Hero = () => {
  return (
    <div className="relative h-[60vh] bg-hero-pattern bg-cover bg-center">
      <div className="absolute inset-0 hero-overlay"></div>
      <div className="relative z-10 flex flex-col items-center justify-center h-full text-white text-center px-4">
        <h1 className="text-5xl md:text-6xl font-bold mb-4">
          TRIP PLANNER
        </h1>
        <p className="text-xl md:text-2xl max-w-2xl">
          Let our AI assistant help you plan your perfect trip with personalized recommendations.
        </p>
        <button 
          className="mt-8 px-6 py-3 bg-tripPurple text-white rounded-full hover:bg-opacity-90 transition-colors shadow-lg"
          onClick={() => document.getElementById('chatSection')?.scrollIntoView({ behavior: 'smooth' })}
        >
          Start Planning
        </button>
      </div>
    </div>
  );
};

export default Hero;
