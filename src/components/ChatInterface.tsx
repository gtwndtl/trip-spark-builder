
import React, { useState } from 'react';
import { Message } from '@/types/chat';
import ChatContainer from './chat/ChatContainer';
import ChatInput from './chat/ChatInput';
import ChatRightPanel from './chat/ChatRightPanel';

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
  
  // Basic handler for sending messages - replace with your LLM integration
  const handleSendMessage = () => {
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
    
    // Simulate typing indicator then response
    // Replace this with your actual LLM integration
    setTimeout(() => {
      setIsTyping(false);
      // This is where you would call your LLM and handle the response
    }, 1000);
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
            preferences={{
              destination: null,
              duration: null,
              budget: null,
              style: null
            }}
            fullItinerary={null}
          />
        </div>
      </div>
    </section>
  );
};

export default ChatInterface;
