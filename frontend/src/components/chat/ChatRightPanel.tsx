
import React from 'react';
import RecommendedTrips from '../RecommendedTrips';
import LongdoMap from '../LongdoMap';
import ItineraryDisplay from './ItineraryDisplay';
import { DayItinerary, FullItinerary } from '@/types';

type ChatRightPanelProps = {
  showItinerary: boolean;
  itinerary: DayItinerary[];
  preferences: {
    destination: string | null;
    duration: number | null;
    budget: string | null;
    style: string | null;
  };
  fullItinerary: FullItinerary | null;
};

const ChatRightPanel = ({ 
  showItinerary, 
  itinerary, 
  preferences, 
  fullItinerary 
}: ChatRightPanelProps) => {
  return (
    <div className="glass-card rounded-xl shadow-lg overflow-hidden">
      {!showItinerary ? (
        <div className="flex flex-col space-y-4">
          <RecommendedTrips />
          <LongdoMap />
        </div>
      ) : (
        <ItineraryDisplay 
          itinerary={itinerary} 
          preferences={preferences} 
          fullItinerary={fullItinerary} 
        />
      )}
    </div>
  );
};

export default ChatRightPanel;
