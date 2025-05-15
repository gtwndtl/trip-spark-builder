
import React from 'react';
import { Message } from '@/types';

type ChatMessageProps = {
  message: Message;
};

const ChatMessage = ({ message }: ChatMessageProps) => {
  return (
    <div className={`message-bubble ${message.sender === 'user' ? 'user-message' : 'bot-message'}`}>
      {message.text.split('\n').map((text, i) => (
        <p key={i} className={i > 0 ? 'mt-2' : ''}>{text}</p>
      ))}
    </div>
  );
};

export default ChatMessage;
