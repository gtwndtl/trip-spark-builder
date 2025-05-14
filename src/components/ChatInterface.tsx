
import React, { useState, useRef, useEffect } from 'react';
import { SendHorizontal, Globe } from 'lucide-react';
import { useToast } from '@/hooks/use-toast';

type Message = {
  id: string;
  text: string;
  sender: 'user' | 'bot';
  timestamp: Date;
};

type TripPreferences = {
  destination: string | null;
  duration: number | null;
  budget: string | null;
  style: string | null;
};

const initialMessages: Message[] = [
  {
    id: '1',
    text: 'สวัสดีค่ะ! ฉันคือผู้ช่วยวางแผนการท่องเที่ยวของคุณ บอกฉันเกี่ยวกับทริปที่คุณต้องการได้เลยค่ะ เช่น คุณอยากไปที่ไหน? จำนวนกี่วัน? มีงบประมาณเท่าไหร่? และสไตล์การท่องเที่ยวแบบไหนที่คุณชอบ?',
    sender: 'bot',
    timestamp: new Date()
  }
];

const ChatInterface = () => {
  const [messages, setMessages] = useState<Message[]>(initialMessages);
  const [newMessage, setNewMessage] = useState('');
  const [isTyping, setIsTyping] = useState(false);
  const [preferences, setPreferences] = useState<TripPreferences>({
    destination: null,
    duration: null,
    budget: null,
    style: null
  });
  const [showItinerary, setShowItinerary] = useState(false);
  const [itinerary, setItinerary] = useState<any[]>([]);
  
  const chatEndRef = useRef<HTMLDivElement>(null);
  const { toast } = useToast();
  
  // Sample itinerary data (in real app, this would come from the API)
  const sampleItinerary = [
    {
      day: 1,
      date: "15 May 2023",
      activities: [
        {
          time: "09:00",
          title: "เช็คอินที่โรงแรม Sukhumvit 11",
          description: "เริ่มต้นทริปของคุณด้วยการเช็คอินที่โรงแรมย่านสุขุมวิท"
        },
        {
          time: "11:30",
          title: "เที่ยวชมวัดพระแก้ว",
          description: "สัมผัสความงดงามของสถาปัตยกรรมไทยและพระแก้วมรกตที่เป็นสิ่งศักดิ์สิทธิ์คู่บ้านคู่เมือง"
        },
        {
          time: "15:00",
          title: "ล่องเรือชมวิวแม่น้ำเจ้าพระยา",
          description: "ชมวิวทิวทัศน์สองฝั่งแม่น้ำเจ้าพระยา พร้อมรับประทานอาหารบนเรือ"
        },
        {
          time: "19:00",
          title: "เดินเที่ยวที่เยาวราช",
          description: "สัมผัสบรรยากาศไชน่าทาวน์ของกรุงเทพฯ และลิ้มลองอาหารจีนแท้ๆ"
        }
      ],
      routes: [
        { from: "โรงแรม", to: "วัดพระแก้ว", transport: "แท็กซี่", duration: "30 นาที" },
        { from: "วัดพระแก้ว", to: "ท่าเรือ", transport: "เดิน", duration: "15 นาที" },
        { from: "ท่าเรือ", to: "เยาวราช", transport: "รถไฟฟ้า MRT", duration: "25 นาที" }
      ]
    },
    {
      day: 2,
      date: "16 May 2023",
      activities: [
        {
          time: "08:30",
          title: "ตลาดน้ำดำเนินสะดวก",
          description: "ตื่นเช้าเพื่อสัมผัสประสบการณ์ตลาดน้ำขึ้นชื่อของไทย"
        },
        {
          time: "13:00",
          title: "อุทยานประวัติศาสตร์พระนครศรีอยุธยา",
          description: "เยี่ยมชมอดีตเมืองหลวงเก่าของไทยที่เป็นมรดกโลก"
        },
        {
          time: "18:00",
          title: "กลับกรุงเทพ - Terminal 21",
          description: "ช้อปปิ้งที่ห้างสรรพสินค้าธีมท่องเที่ยวรอบโลก"
        }
      ],
      routes: [
        { from: "โรงแรม", to: "ตลาดน้ำดำเนินสะดวก", transport: "รถตู้", duration: "1 ชั่วโมง 30 นาที" },
        { from: "ตลาดน้ำดำเนินสะดวก", to: "อยุธยา", transport: "รถตู้", duration: "1 ชั่วโมง 45 นาที" },
        { from: "อยุธยา", to: "Terminal 21", transport: "รถตู้", duration: "1 ชั่วโมง 15 นาที" }
      ]
    }
  ];

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  // Mock function to analyze message and extract preferences
  const analyzeMessage = (message: string) => {
    const updatedPreferences = { ...preferences };
    
    // Very simplified destination detection
    if (message.includes('กรุงเทพ') || message.includes('bangkok')) {
      updatedPreferences.destination = 'กรุงเทพ';
    } else if (message.includes('เชียงใหม่') || message.includes('chiangmai')) {
      updatedPreferences.destination = 'เชียงใหม่';
    } else if (message.includes('ภูเก็ต') || message.includes('phuket')) {
      updatedPreferences.destination = 'ภูเก็ต';
    }
    
    // Simplified duration detection
    const durationMatch = message.match(/(\d+)\s*วัน/);
    if (durationMatch) {
      updatedPreferences.duration = parseInt(durationMatch[1]);
    }
    
    // Simplified budget detection
    if (message.includes('ประหยัด') || message.includes('งบน้อย')) {
      updatedPreferences.budget = 'ประหยัด';
    } else if (message.includes('หรูหรา') || message.includes('luxury')) {
      updatedPreferences.budget = 'หรูหรา';
    } else {
      const budgetMatch = message.match(/งบ\s*(\d+)/);
      if (budgetMatch) {
        updatedPreferences.budget = budgetMatch[1] + ' บาท';
      }
    }
    
    // Simplified style detection
    if (message.includes('ธรรมชาติ') || message.includes('nature')) {
      updatedPreferences.style = 'ธรรมชาติ';
    } else if (message.includes('วัฒนธรรม') || message.includes('culture')) {
      updatedPreferences.style = 'วัฒนธรรม';
    } else if (message.includes('ช้อปปิ้ง') || message.includes('shopping')) {
      updatedPreferences.style = 'ช้อปปิ้ง';
    }
    
    return updatedPreferences;
  };

  const handleSendMessage = () => {
    if (!newMessage.trim()) return;
    
    // Add user message
    const userMsg: Message = {
      id: Date.now().toString(),
      text: newMessage,
      sender: 'user',
      timestamp: new Date()
    };
    
    setMessages([...messages, userMsg]);
    setNewMessage('');
    setIsTyping(true);
    
    // Analyze message to extract preferences
    const updatedPreferences = analyzeMessage(newMessage);
    setPreferences(updatedPreferences);
    
    // Simulate bot thinking and response
    setTimeout(() => {
      let botResponse = '';
      let allPrefsFound = true;
      
      // Check if all preferences are set
      for (const [key, value] of Object.entries(updatedPreferences)) {
        if (value === null) {
          allPrefsFound = false;
          break;
        }
      }
      
      if (allPrefsFound) {
        botResponse = `ขอบคุณค่ะ! ฉันได้รวบรวมความต้องการของคุณแล้ว:
- จุดหมายปลายทาง: ${updatedPreferences.destination}
- ระยะเวลา: ${updatedPreferences.duration} วัน
- งบประมาณ: ${updatedPreferences.budget}
- สไตล์การท่องเที่ยว: ${updatedPreferences.style}

ฉันกำลังวางแผนทริปให้คุณ รอสักครู่นะคะ...`;

        // After 2 seconds, show itinerary
        setTimeout(() => {
          setItinerary(sampleItinerary);
          setShowItinerary(true);
        }, 2000);
      } else {
        // Ask for missing information
        if (!updatedPreferences.destination) {
          botResponse = 'คุณอยากไปเที่ยวที่ไหนคะ?';
        } else if (!updatedPreferences.duration) {
          botResponse = `สำหรับทริปที่ ${updatedPreferences.destination} คุณวางแผนจะไปกี่วันคะ?`;
        } else if (!updatedPreferences.budget) {
          botResponse = 'คุณมีงบประมาณสำหรับทริปนี้เท่าไหร่คะ?';
        } else if (!updatedPreferences.style) {
          botResponse = 'คุณชอบการท่องเที่ยวแบบไหนคะ? เช่น ธรรมชาติ, วัฒนธรรม, ช้อปปิ้ง?';
        }
      }
      
      const botMsg: Message = {
        id: (Date.now() + 1).toString(),
        text: botResponse,
        sender: 'bot',
        timestamp: new Date()
      };
      
      setMessages(prev => [...prev, botMsg]);
      setIsTyping(false);
    }, 1500);
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
        <h2 className="text-3xl font-bold text-center mb-8">Plan Your Trip</h2>
        
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div className="bg-white rounded-xl shadow-lg overflow-hidden">
            <div className="p-4 bg-tripPurple text-white">
              <h3 className="text-xl font-semibold">แชทกับผู้ช่วยวางแผนทริป</h3>
            </div>
            
            <div className="chat-container p-4">
              {messages.map((message) => (
                <div 
                  key={message.id} 
                  className={`message-bubble ${message.sender === 'user' ? 'user-message' : 'bot-message'}`}
                >
                  {message.text.split('\n').map((text, i) => (
                    <p key={i} className={i > 0 ? 'mt-2' : ''}>{text}</p>
                  ))}
                </div>
              ))}
              
              {isTyping && (
                <div className="message-bubble bot-message">
                  <div className="typing-indicator">
                    <span className="typing-dot"></span>
                    <span className="typing-dot"></span>
                    <span className="typing-dot"></span>
                  </div>
                </div>
              )}
              
              <div ref={chatEndRef}></div>
            </div>
            
            <div className="p-4 border-t">
              <div className="flex">
                <textarea
                  value={newMessage}
                  onChange={(e) => setNewMessage(e.target.value)}
                  onKeyDown={handleKeyDown}
                  placeholder="พิมพ์ข้อความของคุณที่นี่..."
                  className="flex-1 p-3 border border-gray-300 rounded-l-md focus:outline-none focus:ring-2 focus:ring-tripPurple"
                  rows={1}
                ></textarea>
                <button
                  onClick={handleSendMessage}
                  className="bg-tripPurple text-white p-3 rounded-r-md hover:bg-opacity-90 transition-colors"
                >
                  <SendHorizontal size={20} />
                </button>
              </div>
            </div>
          </div>
          
          <div className="bg-white rounded-xl shadow-lg overflow-hidden">
            {!showItinerary ? (
              <div className="flex flex-col items-center justify-center h-full p-6 text-center">
                <div className="text-8xl mb-6 text-gray-300">
                  <Globe className="w-24 h-24 mx-auto animate-pulse-slow" />
                </div>
                <h3 className="text-2xl font-semibold text-tripDark mb-2">แผนการท่องเที่ยวของคุณ</h3>
                <p className="text-gray-500">
                  แชทกับผู้ช่วยของเราเพื่อรับแผนการท่องเที่ยวที่ปรับแต่งให้เหมาะกับคุณ
                </p>
              </div>
            ) : (
              <>
                <div className="p-4 bg-tripPurple text-white">
                  <h3 className="text-xl font-semibold">แผนการเดินทางของคุณที่ {preferences.destination}</h3>
                  <p className="text-sm">
                    {preferences.duration} วัน • {preferences.budget} • {preferences.style}
                  </p>
                </div>
                
                <div className="p-6 overflow-auto" style={{ maxHeight: 'calc(100vh - 16rem)' }}>
                  {itinerary.map((day, index) => (
                    <div key={index} className="mb-8">
                      <h4 className="text-xl font-bold mb-4">วันที่ {day.day} - {day.date}</h4>
                      
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
                  
                  <button 
                    className="w-full py-3 bg-tripOrange text-white rounded-md hover:bg-opacity-90 transition-colors mt-4"
                    onClick={() => {
                      toast({
                        title: "แผนการเดินทางถูกบันทึกแล้ว",
                        description: "คุณสามารถเข้าถึงแผนการเดินทางได้ในหน้าโปรไฟล์ของคุณ",
                      });
                    }}
                  >
                    บันทึกแผนการเดินทาง
                  </button>
                </div>
              </>
            )}
          </div>
        </div>
      </div>
    </section>
  );
};

export default ChatInterface;
