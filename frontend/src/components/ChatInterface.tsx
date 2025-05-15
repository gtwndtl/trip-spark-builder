
import React, { useState } from 'react';
import { Message } from '@/types/chat';
import { useNavigate } from 'react-router-dom';
import { useToast } from '@/hooks/use-toast';
import ChatContainer from './chat/ChatContainer';
import ChatInput from './chat/ChatInput';
import ChatRightPanel from './chat/ChatRightPanel';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { TripPreferences } from '@/types';

// Initial welcome message
const initialMessages: Message[] = [
  {
    id: '1',
    text: 'สวัสดีค่ะ! ฉันคือผู้ช่วยวางแผนการท่องเที่ยวของคุณ บอกฉันเกี่ยวกับทริปที่คุณต้องการได้เลยค่ะ',
    sender: 'bot',
    timestamp: new Date()
  }
];

const ChatInterface = () => {
  const [messages, setMessages] = useState<Message[]>(initialMessages);
  const [newMessage, setNewMessage] = useState('');
  const [isTyping, setIsTyping] = useState(false);
  const [showItinerary, setShowItinerary] = useState(false);
  const [tripPreferences, setTripPreferences] = useState<TripPreferences | null>(null);
  const [showConfirmDialog, setShowConfirmDialog] = useState(false);
  const { toast } = useToast();
  const navigate = useNavigate();
  
  const handleSendMessage = async () => {
    if (!newMessage.trim()) return;
    
    // Add user message to chat
    const userMsg: Message = {
      id: Date.now().toString(),
      text: newMessage,
      sender: 'user',
      timestamp: new Date()
    };
    
    setMessages([...messages, userMsg]);
    setNewMessage('');
    setIsTyping(true);
    
    try {
      // Send message to API
      const response = await fetch('http://127.0.0.1:8000/plan-trip', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: newMessage }),
      });
      
      if (!response.ok) {
        throw new Error('API request failed');
      }
      
      const data = await response.json();
      
      // Process the parsed data
      if (data.parsed) {
        const parsedPreferences = {
          destination: data.parsed.destination || null,
          duration: data.parsed.days || null,
          budget: data.parsed.budget ? `${data.parsed.budget} บาท` : null,
          style: data.parsed.style || null,
        };
        
        setTripPreferences(parsedPreferences);
        
        // Create confirmation bot message
        const botMsg: Message = {
          id: Date.now().toString(),
          text: `ตามที่คุณต้องการ ฉันเข้าใจว่าคุณต้องการเดินทางไปที่ ${parsedPreferences.destination} เป็นเวลา ${parsedPreferences.duration} วัน ในรูปแบบ ${parsedPreferences.style} ด้วยงบประมาณ ${parsedPreferences.budget}\n\nข้อมูลถูกต้องไหมคะ? เราจะเริ่มวางแผนการเดินทางเลยมั้ย?`,
          sender: 'bot',
          timestamp: new Date()
        };
        
        setIsTyping(false);
        setMessages(prev => [...prev, botMsg]);
        setShowConfirmDialog(true);
      }
    } catch (error) {
      console.error('Error:', error);
      
      // Add error message
      const errorMsg: Message = {
        id: Date.now().toString(),
        text: 'ขออภัยค่ะ เกิดข้อผิดพลาดในการประมวลผลข้อมูล กรุณาลองใหม่อีกครั้ง',
        sender: 'bot',
        timestamp: new Date()
      };
      
      setIsTyping(false);
      setMessages(prev => [...prev, errorMsg]);
      
      toast({
        title: "เกิดข้อผิดพลาด",
        description: "ไม่สามารถประมวลผลข้อมูลได้ กรุณาลองใหม่อีกครั้ง",
        variant: "destructive",
      });
    }
  };

  const handleConfirmTrip = async () => {
    setShowConfirmDialog(false);
    setIsTyping(true);
    
    try {
      // Send trip preferences to the backend trip planner API
      const response = await fetch('http://127.0.0.1:8000/trip-planner', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          destination: tripPreferences?.destination,
          duration: tripPreferences?.duration?.toString(),
          budget: tripPreferences?.budget,
          style: tripPreferences?.style
        }),
      });
      
      if (!response.ok) {
        throw new Error('Failed to plan trip');
      }
      
      const result = await response.json();
      
      // Show processing message
      const processingMsg: Message = {
        id: Date.now().toString(),
        text: 'กำลังวางแผนการเดินทางให้คุณ กรุณารอสักครู่...',
        sender: 'bot',
        timestamp: new Date()
      };
      
      setIsTyping(false);
      setMessages(prev => [...prev, processingMsg]);
      
      // Process the JSON data from the plan field
      let itineraryData;
      if (result.plan) {
        try {
          // Extract the JSON part from the response
          const jsonMatch = result.plan.match(/```json\n([\s\S]*?)\n```/);
          if (jsonMatch && jsonMatch[1]) {
            itineraryData = JSON.parse(jsonMatch[1]);
          }
        } catch (error) {
          console.error('Error parsing JSON from result:', error);
        }
      }
      
      // Navigate to trip summary page with the result data
      setTimeout(() => {
        navigate('/trip-summary', { 
          state: { 
            itinerary: {
              preferences: {
                destination: tripPreferences?.destination,
                duration: tripPreferences?.duration,
                budget: tripPreferences?.budget,
                style: tripPreferences?.style
              },
              days: itineraryData?.days || []
            }
          } 
        });
      }, 1500);
      
    } catch (error) {
      console.error('Error planning trip:', error);
      
      const errorMsg: Message = {
        id: Date.now().toString(),
        text: 'ขออภัยค่ะ เกิดข้อผิดพลาดในการวางแผนการเดินทาง กรุณาลองใหม่อีกครั้ง',
        sender: 'bot',
        timestamp: new Date()
      };
      
      setIsTyping(false);
      setMessages(prev => [...prev, errorMsg]);
      
      toast({
        title: "เกิดข้อผิดพลาด",
        description: "ไม่สามารถวางแผนการเดินทางได้ กรุณาลองใหม่อีกครั้ง",
        variant: "destructive",
      });
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  return (
    <section id="chatSection" className="py-16">
      <div className="container mx-auto px-4">
        <h2 className="text-3xl font-light text-center mb-3">
          <span className="font-normal">Plan</span> Your Trip
        </h2>
        <p className="text-center text-gray-600 mb-8">วางแผนการเดินทางของคุณด้วยผู้ช่วย AI ส่วนตัว</p>
        
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div className="glass-card rounded-xl shadow-lg overflow-hidden">
            <div className="p-4 bg-tripPurple text-white">
              <h3 className="text-xl font-light">แชทกับผู้ช่วยวางแผนทริป</h3>
            </div>
            
            <ChatContainer messages={messages} isTyping={isTyping} />
            <ChatInput 
              newMessage={newMessage}
              setNewMessage={setNewMessage}
              handleSendMessage={handleSendMessage}
              handleKeyDown={handleKeyDown}
            />
          </div>
          
          <ChatRightPanel
            showItinerary={showItinerary}
            itinerary={[]}
            preferences={tripPreferences || {
              destination: null,
              duration: null,
              budget: null,
              style: null
            }}
            fullItinerary={null}
          />
        </div>
      </div>
      
      {/* Confirmation Dialog */}
      <Dialog open={showConfirmDialog} onOpenChange={setShowConfirmDialog}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>ยืนยันการวางแผนการเดินทาง</DialogTitle>
            <DialogDescription>
              {tripPreferences && (
                <div className="py-4">
                  <p className="mb-2"><strong>จุดหมายปลายทาง:</strong> {tripPreferences.destination}</p>
                  <p className="mb-2"><strong>ระยะเวลา:</strong> {tripPreferences.duration} วัน</p>
                  <p className="mb-2"><strong>งบประมาณ:</strong> {tripPreferences.budget}</p>
                  <p className="mb-2"><strong>สไตล์การท่องเที่ยว:</strong> {tripPreferences.style}</p>
                </div>
              )}
              <p>ข้อมูลถูกต้องหรือไม่? เราจะเริ่มวางแผนการเดินทางเลยมั้ย</p>
            </DialogDescription>
          </DialogHeader>
          <DialogFooter className="flex justify-end gap-3">
            <Button variant="outline" onClick={() => setShowConfirmDialog(false)}>
              แก้ไขข้อมูล
            </Button>
            <Button onClick={handleConfirmTrip}>
              เริ่มวางแผนการเดินทาง
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </section>
  );
};

export default ChatInterface;