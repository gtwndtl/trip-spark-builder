import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { SendHorizontal } from 'lucide-react';
import { useToast } from '@/hooks/use-toast';
import { useIsMobile } from '@/hooks/use-mobile';
import RecommendedTrips from './RecommendedTrips';
import LongdoMap from './LongdoMap';
import { FullItinerary, TripPreferences } from '@/types';

type Message = {
  id: string;
  text: string;
  sender: 'user' | 'bot';
  timestamp: Date;
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
  const [fullItinerary, setFullItinerary] = useState<FullItinerary | null>(null);
  
  const chatEndRef = useRef<HTMLDivElement>(null);
  const { toast } = useToast();
  const isMobile = useIsMobile();
  const navigate = useNavigate();
  
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
          // Sample itinerary data
          const sampleItinerary = [
            {
              day: 1,
              date: "15 พฤษภาคม 2568",
              activities: [
                {
                  time: "09:00",
                  title: "เช็คอินที่โรงแรม Grande Centre Point สุขุมวิท 55",
                  description: "เริ่มต้นทริปของคุณด้วยการเช็คอินที่โรงแรมย่านสุขุมวิท ใกล้ BTS ทองหล่อ"
                },
                {
                  time: "11:30",
                  title: "ทานอาหารกลางวันที่ร้าน สมบูรณ์โภชนา",
                  description: "ลิ้มลองอาหารไทยต้นตำรับที่เปิดมายาวนานกว่า 50 ปี"
                },
                {
                  time: "14:00",
                  title: "เที่ยวชมวัดพระแก้ว",
                  description: "สัมผัสความงดงามของสถาปัตยกรรมไทยและพระแก้วมรกตที่เป็นสิ่งศักดิ์สิทธิ์คู่บ้านคู่เมือง"
                },
                {
                  time: "17:00",
                  title: "ล่องเรือชมวิวแม่น้ำเจ้าพระยา",
                  description: "ชมวิวทิวทัศน์สองฝั่งแม่น้ำเจ้าพระยา พร้อมรับประทานอาหารบนเรือ"
                },
                {
                  time: "20:00",
                  title: "เดินเที่ยวที่เยาวราช",
                  description: "สัมผัสบรรยากาศไชน่าทาวน์ของกรุงเทพฯ และลิ้มลองอาหารจีนแท้ๆ"
                }
              ],
              routes: [
                { from: "โรงแรม", to: "ร้านสมบูรณ์โภชนา", transport: "แท็กซี่", duration: "25 นาที" },
                { from: "ร้านสมบูรณ์โภชนา", to: "วัดพระแก้ว", transport: "แท็กซี่", duration: "30 นาที" },
                { from: "วัดพระแก้ว", to: "ท่าเรือ", transport: "เดิน", duration: "15 นาที" },
                { from: "ท่าเรือ", to: "เยาวราช", transport: "รถไฟฟ้า MRT", duration: "25 นาที" }
              ]
            },
            {
              day: 2,
              date: "16 พฤษภาคม 2568",
              activities: [
                {
                  time: "08:00",
                  title: "อาหารเช้าที่โรงแรม",
                  description: "เริ่มต้นวันด้วยบุฟเฟ่ต์อาหารเช้านานาชาติที่โรงแรม"
                },
                {
                  time: "10:00",
                  title: "ตลาดน้ำดำเนินสะดวก",
                  description: "สัมผัสประสบการณ์ตลาดน้ำขึ้นชื่อของไทย"
                },
                {
                  time: "14:00",
                  title: "อุทยานประวัติศาสตร์พระนครศรีอยุธยา",
                  description: "เยี่ยมชมอดีตเมืองหลวงเก่าของไทยที่เป็นมรดกโลก"
                },
                {
                  time: "18:30",
                  title: "กลับกรุงเทพ - ทานอาหารเย็นที่ Terminal 21",
                  description: "ช้อปปิ้งและทานอาหารที่ห้างสรรพสินค้าธีมท่องเที่ยวรอบโลก"
                }
              ],
              routes: [
                { from: "โรงแรม", to: "ตลาดน้ำดำเนินสะดวก", transport: "รถตู้", duration: "1 ชั่วโมง 30 นาที" },
                { from: "ตลาดน้ำดำเนินสะดวก", to: "อยุธยา", transport: "รถตู้", duration: "1 ชั่วโมง 45 นาที" },
                { from: "อยุธยา", to: "Terminal 21", transport: "รถตู้", duration: "1 ชั่วโมง 15 นาที" },
                { from: "Terminal 21", to: "โรงแรม", transport: "BTS", duration: "10 นาที" }
              ]
            },
            {
              day: 3,
              date: "17 พฤษภาคม 2568",
              activities: [
                {
                  time: "09:00",
                  title: "อาหารเช้าที่โรงแรม",
                  description: "เริ่มต้นวันด้วยบุฟเฟ่ต์อาหารเช้านานาชาติที่โรงแรม"
                },
                {
                  time: "11:00",
                  title: "ช้อปปิ้งที่สยามพารากอน",
                  description: "เพลิดเพลินกับการช้อปปิ้งสินค้าแบรนด์เนมและสินค้าไทย"
                },
                {
                  time: "14:30",
                  title: "พิพิธภัณฑ์ศิลปะไทยร่วมสมัย (MOCA)",
                  description: "ชมงานศิลปะร่วมสมัยที่ใหญ่ที่สุดในประเทศไทย"
                },
                {
                  time: "18:00",
                  title: "อาหารเย็นที่ร้าน เจ๊ไฝ ทองหล่อ",
                  description: "ปิดท้ายทริปด้วยอาหารไทย-จีนรสเด็ด"
                },
                {
                  time: "21:00",
                  title: "เครื่องดื่มที่ Octave Rooftop Bar",
                  description: "ชมวิวกรุงเทพฯ ยามค่ำคืนที่บาร์บนดาดฟ้า"
                }
              ],
              routes: [
                { from: "โรงแรม", to: "สยามพารากอน", transport: "BTS", duration: "20 นาที" },
                { from: "สยามพารากอน", to: "MOCA", transport: "แท็กซี่", duration: "25 นาที" },
                { from: "MOCA", to: "ร้านเจ๊ไฝ", transport: "แท็กซี่", duration: "30 นาที" },
                { from: "ร้านเจ๊ไฝ", to: "Octave Rooftop Bar", transport: "เดิน", duration: "15 นาที" },
                { from: "Octave Rooftop Bar", to: "โรงแรม", transport: "เดิน", duration: "10 นาที" }
              ]
            }
          ];
          
          setItinerary(sampleItinerary);
          
          const completeItinerary: FullItinerary = {
            preferences: {
              destination: updatedPreferences.destination,
              duration: updatedPreferences.duration,
              budget: updatedPreferences.budget, 
              style: updatedPreferences.style
            },
            days: sampleItinerary
          };
          
          setFullItinerary(completeItinerary);
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

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

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
                  className="flex-1 p-3 border border-gray-200 rounded-l-md focus:outline-none focus:ring-2 focus:ring-tripPurple"
                  rows={1}
                ></textarea>
                <button
                  onClick={handleSendMessage}
                  className="apple-button p-3 rounded-r-md"
                >
                  <SendHorizontal size={20} />
                </button>
              </div>
            </div>
          </div>
          
          <div className="glass-card rounded-xl shadow-lg overflow-hidden">
            {!showItinerary ? (
              <div className="flex flex-col space-y-4">
                <RecommendedTrips />
                <LongdoMap />
              </div>
            ) : (
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
            )}
          </div>
        </div>
      </div>
    </section>
  );
};

export default ChatInterface;
