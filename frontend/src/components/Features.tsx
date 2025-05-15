
import React from 'react';
import { CheckCircle, Globe, Calendar } from 'lucide-react';

const Features = () => {
  const features = [
    {
      icon: <CheckCircle className="w-10 h-10 text-tripPurple" />,
      title: "Personalized Recommendations",
      description: "รับคำแนะนำที่ปรับแต่งตามความต้องการ งบประมาณ และสไตล์การท่องเที่ยวของคุณ"
    },
    {
      icon: <Calendar className="w-10 h-10 text-tripPurple" />,
      title: "Detailed Itineraries",
      description: "รับแผนการเดินทางแบบวันต่อวันพร้อมกิจกรรม เส้นทาง และคำแนะนำด้านเวลา"
    },
    {
      icon: <Globe className="w-10 h-10 text-tripPurple" />,
      title: "Destination Insights",
      description: "ค้นพบสถานที่ลับและสถานที่ยอดนิยมในท้องถิ่นที่จุดหมายปลายทางที่คุณเลือก"
    }
  ];

  return (
    <section className="py-20 bg-gray-50">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl font-light text-center mb-3">
          <span className="font-normal">How</span> It Works
        </h2>
        <p className="text-center text-gray-600 mb-12">วิธีการใช้งานที่ง่ายและรวดเร็ว</p>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {features.map((feature, index) => (
            <div key={index} className="glass-card p-8 rounded-xl text-center transition-all hover:transform hover:scale-[1.02] hover:shadow-lg">
              <div className="flex justify-center mb-6 relative">
                {feature.icon}
                <div className="absolute -z-10 w-16 h-16 bg-purple-100 rounded-full opacity-70 blur-lg"></div>
              </div>
              <h3 className="text-xl font-medium mb-3">{feature.title}</h3>
              <p className="text-gray-600">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Features;
