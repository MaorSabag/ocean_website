import React, { useEffect, useState } from 'react';
import "./index.css";

const NotFoundPage = () => {
  const [text, setText] = useState('');
  const [index, setIndex] = useState(0);

  const fullText = "404 Page Not Found\nTry to find another way!";
  const typingDelay = 100; // Delay between each character typing

  useEffect(() => {
    const timer = setTimeout(() => {
      if (index < fullText.length) {
        setText(fullText.substring(0, index + 1));
        setIndex(index + 1);
      }
    }, typingDelay);

    return () => clearTimeout(timer);
  }, [index, fullText]);

  return (
    <div className="not-found-container">
      <div className="typing-machine">
        <span className="typing-text">{text}</span>
        <span className="typing-cursor" />
      </div>
    </div>
  );
};

export default NotFoundPage;
