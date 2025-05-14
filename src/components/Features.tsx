
import React from 'react';
import { CheckCircle, Globe, Calendar } from 'lucide-react';

const Features = () => {
  const features = [
    {
      icon: <CheckCircle className="w-10 h-10 text-tripPurple" />,
      title: "Personalized Recommendations",
      description: "Get tailored suggestions based on your preferences, budget, and travel style."
    },
    {
      icon: <Calendar className="w-10 h-10 text-tripPurple" />,
      title: "Detailed Itineraries",
      description: "Receive day-by-day plans with activities, routes, and timing recommendations."
    },
    {
      icon: <Globe className="w-10 h-10 text-tripPurple" />,
      title: "Destination Insights",
      description: "Discover hidden gems and local favorites at your chosen destination."
    }
  ];

  return (
    <section className="py-16 bg-gray-50">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl font-bold text-center mb-12">How It Works</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {features.map((feature, index) => (
            <div key={index} className="bg-white p-6 rounded-lg shadow-md text-center">
              <div className="flex justify-center mb-4">
                {feature.icon}
              </div>
              <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
              <p className="text-gray-600">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Features;
