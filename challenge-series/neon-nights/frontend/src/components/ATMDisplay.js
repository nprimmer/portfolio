import React, { useState, useEffect } from 'react';
import './ATMDisplay.css';

const ATMDisplay = ({ input, message, flash }) => {
  const [flashBackground, setFlashBackground] = useState(false);

  useEffect(() => {
    if (flash) {
      setFlashBackground(true);
      setTimeout(() => setFlashBackground(false), 300);
    }
  }, [flash]);

  return (
    <div className={`atm-display ${flashBackground ? 'flash' : ''}`}>
      <p>ENTER 4 DIGIT PIN:</p>
      <p>{'*'.repeat(input.length)}</p>
      {message && <p className="error-message">{message}</p>}
    </div>
  );
};

export default ATMDisplay;
