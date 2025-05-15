
import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useToast } from '@/hooks/use-toast';
import { useIsMobile } from '@/hooks/use-mobile';
import { FullItinerary, DayItinerary } from '@/types';

type ItineraryDisplayProps = {
  itinerary: DayItinerary[];
  preferences: {
    destination: string | null;
    duration: number | null;
    budget: string | null;
    style: string | null;
  };
  fullItinerary: FullItinerary | null;
};

const ItineraryDisplay = ({ itinerary, preferences, fullItinerary }: ItineraryDisplayProps) => {
  const { toast } = useToast();
  const isMobile = useIsMobile();
  const navigate = useNavigate();

  const handleViewFullSummary = () => {
    if (fullItinerary) {
      navigate('/trip-summary', { state: { itinerary: fullItinerary } });
    } else {
      toast({
        title: "ไม่พบข้อมูลแผนการเดินทาง",
        description: "กรุณาสร้างแผนการเดินทางให้เสร็จสมบูรณ์ก่อน",
      });
    }
  };

  return (
    <>
      <div className="p-4 bg-tripPurple text-white">
        <h3 className="text-xl font-light">แผนการเดินทางของคุณที่ {preferences.destination}</h3>
        <p className="text-sm">
          {preferences.duration} วัน • {preferences.budget} • {preferences.style}
        </p>
      </div>
      
      <div className="p-6 overflow-auto" style={{ 
        maxHeight: isMobile ? 'calc(50vh)' : 'calc(100vh - 16rem)' 
      }}>
        {itinerary.map((day, index) => (
          <div key={index} className="mb-8">
            <h4 className="text-xl font-medium mb-4">วันที่ {day.day} - {day.date}</h4>
            
            <div className="ml-4">
              {day.activities.map((activity: any, actIdx: number) => (
                <div key={actIdx} className="itinerary-day">
                  <div className="itinerary-time">{activity.time}</div>
                  <div className="itinerary-activity">
                    <div className="itinerary-activity-title">{activity.title}</div>
                    <div className="itinerary-activity-description">{activity.description}</div>
                  </div>
                  
                  {actIdx < day.routes.length && (
                    <div className="itinerary-route">
                      <span className="itinerary-route-icon">→</span>
                      <span>
                        {day.routes[actIdx].from} ถึง {day.routes[actIdx].to} ({day.routes[actIdx].transport}, {day.routes[actIdx].duration})
                      </span>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        ))}
        
        <div className="flex gap-3 mt-4">
          <button 
            className="apple-button flex-1 py-3 text-white rounded-md hover:bg-opacity-90"
            onClick={() => {
              toast({
                title: "แผนการเดินทางถูกบันทึกแล้ว",
                description: "คุณสามารถเข้าถึงแผนการเดินทางได้ในหน้าโปรไฟล์ของคุณ",
              });
            }}
          >
            บันทึกแผนการเดินทาง
          </button>
          
          <button 
            className="bg-tripOrange/80 flex-1 py-3 text-white rounded-md hover:bg-tripOrange transition-colors"
            onClick={handleViewFullSummary}
          >
            ดูสรุปแผนการเดินทาง
          </button>
        </div>
      </div>
    </>
  );
};

export default ItineraryDisplay;
