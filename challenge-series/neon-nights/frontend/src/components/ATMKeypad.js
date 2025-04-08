import React from 'react';
import './ATMKeypad.css';

const ATMKeypad = ({
  input,
  setInput,
  setFlash,
  setMessage,
  setEndpoint,
  invalidAttempts,
  setInvalidAttempts,
  backendHostname,
  endpoint
}) => {
  const handleNumberClick = (number) => {
    if (input.length < 4) {
      setInput([...input, number]);
    }
  };

  const handleDeleteClick = () => {
    setInput(input.slice(0, -1));
  };

  const handleOkClick = async () => {
    if (input.length !== 4) {
      setMessage('ONLY 4 DIGIT PINS ACCEPTED');
      setFlash(true);
      setTimeout(() => setMessage(''), 3000);
      setInput([]); // Clear the input array
      return;
    }

    try {
      const response = await fetch(`${backendHostname}${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ pin: input.join('') })
      });

      const data = await response.json();
      console.log('Response status:', response.status);
      console.log('Response data:', data);

      if (response.status === 401 && data.error === 'Invalid PIN') {
        setInvalidAttempts((prev) => ({
          ...prev,
          [endpoint]: (prev[endpoint] || 0) + 1
        }));
        setMessage(data.error);
        setFlash(true);
        setTimeout(() => setMessage(''), 3000);
      } else if (response.ok && data.next_key) {
        setEndpoint(data.next_key);
        setMessage('');
      } else {
        console.error('Unexpected response:', response.status, data);
        throw new Error('Unexpected response');
      }
    } catch (error) {
      console.error('Error submitting pin:', error);
    } finally {
      setInput([]); // Clear the input array
    }
  };

  return (
    <div className="atm-keypad-container">
      <div className="card-slot"></div>
      <div className="atm-keypad">
        {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((number) => (
          <button key={number} className="keypad-button" onClick={() => handleNumberClick(number)}>
            {number}
          </button>
        ))}
        <button className="keypad-button" onClick={handleDeleteClick}>DEL</button>
        <button className="keypad-button" onClick={() => handleNumberClick(0)}>0</button>
        <button className="keypad-button" onClick={handleOkClick}>OK</button>
      </div>
    </div>
  );
};

export default ATMKeypad;
