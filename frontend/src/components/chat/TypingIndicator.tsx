
import React from 'react';

const TypingIndicator = () => {
  return (
    <div className="message-bubble bot-message">
      <div className="typing-indicator">
        <span className="typing-dot"></span>
        <span className="typing-dot"></span>
        <span className="typing-dot"></span>
      </div>
    </div>
  );
};

export default TypingIndicator;
