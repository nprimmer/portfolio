import React from 'react';
import './WalletContents.css';
import DebitCard from './DebitCard';
import DriversLicense from './DriversLicense';
import CreditCard from './CreditCard';
import Note from './Note';
import ATMReceipt from './ATMReceipt';

const WalletContents = ({ contents, endpoint }) => {
  const sortedContents = [...contents].sort((a, b) => {
    const order = ['debit_card', 'driver_license', 'credit_card', 'note', 'atm_receipt'];
    const aIndex = order.indexOf(a.type);
    const bIndex = order.indexOf(b.type);

    if (aIndex === -1 && bIndex === -1) return 0; // Both are not in the predefined order
    if (aIndex === -1) return 1; // a is not in the predefined order
    if (bIndex === -1) return -1; // b is not in the predefined order
    return aIndex - bIndex; // Both are in the predefined order
  });

  return (
    <div className="wallet-contents">
      {sortedContents.map((item, index) => (
        <div key={index} className="wallet-item">
          {item.type === 'debit_card' ? (
            <div className="debit-card-container">
              <DebitCard card={item} />
            </div>
          ) : item.type === 'driver_license' ? (
            <div className="drivers-license-container">
              <DriversLicense license={item} endpoint={endpoint} />
            </div>
          ) : item.type === 'credit_card' ? (
            <div className="credit-card-container">
              <CreditCard card={item} />
            </div>
          ) : item.type === 'note' ? (
            <div className="note-container">
              <Note note={item} />
            </div>
          ) : item.type === 'atm_receipt' ? (
            <div className="atm-receipt-container">
              <ATMReceipt transactionId={item.transaction_id} />
            </div>
          ) : (
            <pre>{JSON.stringify(item, null, 2)}</pre>
          )}
        </div>
      ))}
    </div>
  );
};

export default WalletContents;
