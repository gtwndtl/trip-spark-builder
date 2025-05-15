
import React, { useRef, useEffect } from 'react';
import ChatMessage from './ChatMessage';
import TypingIndicator from './TypingIndicator';
import { Message } from '@/types/chat';

type ChatContainerProps = {
  messages: Message[];
  isTyping: boolean;
};

const ChatContainer = ({ messages, isTyping }: ChatContainerProps) => {
  const chatEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages, isTyping]);

  return (
    <div className="chat-container p-4">
      {messages.map((message) => (
        <ChatMessage key={message.id} message={message} />
      ))}
      
      {isTyping && <TypingIndicator />}
      
      <div ref={chatEndRef}></div>
    </div>
  );
};

export default ChatContainer;
