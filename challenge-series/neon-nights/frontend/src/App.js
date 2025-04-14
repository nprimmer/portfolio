import React, { useEffect, useState } from 'react';
import './App.css';
import ATMDisplay from './components/ATMDisplay';
import ATMKeypad from './components/ATMKeypad';
import WalletContents from './components/WalletContents';

const BACKEND_URL = 'https://first-miami-backend.toolchest.app/';

const App = () => {
  const [currentUser, setCurrentUser] = useState(null);
  const [input, setInput] = useState([]);
  const [message, setMessage] = useState('');
  const [flash, setFlash] = useState(false);
  const [endpoint, setEndpoint] = useState('tyler-ponce');
  const [invalidAttempts, setInvalidAttempts] = useState({});

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(BACKEND_URL + endpoint);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setCurrentUser(data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [endpoint]);

  return (
    <div className="app-container">
      <div className="left-column">
        {currentUser && <WalletContents contents={currentUser.wallet_contents} endpoint={endpoint} />}
      </div>
      <div className="center-column">
        <div className="top-row">Miami First Federal Bank of Miami</div>
        <div className="middle-row">
          <ATMDisplay text="Welcome to Miami First Federal" input={input} message={message} flash={flash} />
          <ATMKeypad
            input={input}
            setInput={setInput}
            setFlash={setFlash}
            setMessage={setMessage}
            setEndpoint={setEndpoint}
            invalidAttempts={invalidAttempts}
            setInvalidAttempts={setInvalidAttempts}
            backendHostname={BACKEND_URL}
            endpoint={endpoint}
          />
        </div>
        <div className="bottom-row">
          <div className="vent">
            {Array.from({ length: 10 }).map((_, index) => (
              <div key={index} className="vent-slit"></div>
            ))}
          </div>
        </div>
      </div>
      <div className="right-column"></div>
    </div>
  );
};

export default App;
