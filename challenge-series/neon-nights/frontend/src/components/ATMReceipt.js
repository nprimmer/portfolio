import React from 'react';
import './ATMReceipt.css';

const ATMReceipt = ({ transactionId }) => {
  return (
    <div className="atm-receipt">
      <div className="receipt-header">
        <p>Miami First Federal Bank of Miami</p>
        <p>ATM Withdrawal Receipt</p>
      </div>
      <div className="receipt-body">
        <p>Transaction ID:</p>
        <p>{transactionId}</p>
        <p>Date: 07/08/1984</p>
        <p>Time: 12:34 PM</p>
        <p>Withdrawal Amount: $200.00</p>
        <p>Account Balance: $1,800.00</p>
      </div>
      <div className="receipt-footer">
        <p>Thank you for banking with us!</p>
      </div>
    </div>
  );
};

export default ATMReceipt;
