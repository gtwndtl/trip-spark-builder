
import React from 'react';
import { 
  Table, 
  TableHeader, 
  TableBody, 
  TableFooter,
  TableHead, 
  TableRow, 
  TableCell,
  TableCaption 
} from "@/components/ui/table";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card";
import { Calendar, Clock, MapPin, Bed, Utensils } from "lucide-react";
import { FullItinerary } from '@/types';

interface ItinerarySummaryProps {
  itinerary: FullItinerary;
}

const ItinerarySummary: React.FC<ItinerarySummaryProps> = ({ itinerary }) => {
  if (!itinerary || !itinerary.days || itinerary.days.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>ยังไม่มีแผนการเดินทาง</CardTitle>
          <CardDescription>กรุณาสร้างแผนการเดินทางก่อน</CardDescription>
        </CardHeader>
      </Card>
    );
  }

  return (
    <div className="w-full space-y-6">
      <Card className="glass-card">
        <CardHeader className="bg-tripPurple text-white rounded-t-lg">
          <CardTitle className="text-2xl">สรุปแผนการเดินทาง {itinerary.preferences.destination}</CardTitle>
          <CardDescription className="text-white/80">
            {itinerary.preferences.duration} วัน • {itinerary.preferences.budget} • {itinerary.preferences.style}
          </CardDescription>
        </CardHeader>
        
        <CardContent className="p-6">
          {itinerary.days.map((day, index) => (
            <div key={index} className="mb-8">
              <h3 className="text-xl font-medium flex items-center gap-2 mb-4">
                <Calendar className="h-5 w-5 text-tripPurple" />
                วันที่ {day.day} - {day.date}
              </h3>
              
              <Table>
                <TableHeader>
                  <TableRow className="bg-muted/40">
                    <TableHead className="w-[150px]">เวลา</TableHead>
                    <TableHead className="w-[200px]">กิจกรรม</TableHead>
                    <TableHead>รายละเอียด</TableHead>
                  </TableRow>
                </TableHeader>
                
                <TableBody>
                  {day.activities.map((activity, actIdx) => (
                    <React.Fragment key={actIdx}>
                      <TableRow className="hover:bg-muted/30">
                        <TableCell className="font-medium">
                          <div className="flex items-center gap-2">
                            <Clock className="h-4 w-4 text-tripOrange" /> 
                            {activity.time}
                          </div>
                        </TableCell>
                        <TableCell>
                          {activity.title.includes("เช็คอินที่โรงแรม") || activity.title.toLowerCase().includes("hotel") ? (
                            <div className="flex items-center gap-2">
                              <Bed className="h-4 w-4 text-tripPurple" /> 
                              {activity.title}
                            </div>
                          ) : activity.title.includes("อาหาร") || activity.title.includes("ทาน") || activity.title.toLowerCase().includes("food") ? (
                            <div className="flex items-center gap-2">
                              <Utensils className="h-4 w-4 text-tripOrange" /> 
                              {activity.title}
                            </div>
                          ) : (
                            <div className="flex items-center gap-2">
                              <MapPin className="h-4 w-4 text-tripPurple" /> 
                              {activity.title}
                            </div>
                          )}
                        </TableCell>
                        <TableCell>{activity.description}</TableCell>
                      </TableRow>
                      
                      {actIdx < day.routes.length && (
                        <TableRow className="bg-muted/20">
                          <TableCell colSpan={3} className="text-sm text-muted-foreground py-2">
                            <div className="flex items-center">
                              <div className="h-4 w-4 text-muted-foreground mr-2">→</div>
                              <span>
                                เดินทางจาก {day.routes[actIdx].from} ไป {day.routes[actIdx].to} ({day.routes[actIdx].transport}, {day.routes[actIdx].duration})
                              </span>
                            </div>
                          </TableCell>
                        </TableRow>
                      )}
                    </React.Fragment>
                  ))}
                </TableBody>
              </Table>
            </div>
          ))}
        </CardContent>
      </Card>
    </div>
  );
};

export default ItinerarySummary;
