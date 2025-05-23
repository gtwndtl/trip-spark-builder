
import React from 'react';
import { SendHorizontal } from 'lucide-react';

type ChatInputProps = {
  newMessage: string;
  setNewMessage: (message: string) => void;
  handleSendMessage: () => void;
  handleKeyDown: (e: React.KeyboardEvent) => void;
};

const ChatInput = ({ 
  newMessage, 
  setNewMessage, 
  handleSendMessage,
  handleKeyDown 
}: ChatInputProps) => {
  return (
    <div className="p-4 border-t">
      <div className="flex">
        <textarea
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="บอกฉันเกี่ยวกับทริปที่คุณต้องการ เช่น สถานที่ งบประมาณ จำนวนวัน และสไตล์การท่องเที่ยว"
          className="flex-1 p-3 border border-gray-200 rounded-l-md focus:outline-none focus:ring-2 focus:ring-tripPurple"
          rows={2}
        ></textarea>
        <button
          onClick={handleSendMessage}
          className="apple-button p-3 rounded-r-md"
          disabled={!newMessage.trim()}
        >
          <SendHorizontal size={20} />
        </button>
      </div>
    </div>
  );
};

export default ChatInput;
