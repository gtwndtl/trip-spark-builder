
export interface TripPreferences {
  destination: string | null;
  duration: number | null;
  budget: string | null;
  style: string | null;
}

export interface Activity {
  time: string;
  title: string;
  description: string;
}

export interface Route {
  from: string;
  to: string;
  transport: string;
  duration: string;
}

export interface DayItinerary {
  day: number;
  date: string;
  activities: Activity[];
  routes: Route[];
}

export interface FullItinerary {
  preferences: TripPreferences;
  days: DayItinerary[];
}
