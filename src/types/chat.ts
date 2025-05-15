
import { TripPreferences } from './index';

export type Message = {
  id: string;
  text: string;
  sender: 'user' | 'bot';
  timestamp: Date;
};

export type Activity = {
  time: string;
  title: string;
  description: string;
};

export type Route = {
  from: string;
  to: string;
  transport: string;
  duration: string;
};

export type DayItinerary = {
  day: number;
  date: string;
  activities: Activity[];
  routes: Route[];
};
