import React from 'react';
import './DebitCard.css';

const DebitCard = ({ card }) => {
  const { name, number, issued_date, expiration_date } = card;
  const issuedDate = new Date(issued_date.year, issued_date.month - 1); // Months are 0-based in JS
  const expirationDate = new Date(expiration_date.year, expiration_date.month - 1);

  const formatMMYY = (date) => {
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const year = String(date.getFullYear()).slice(-2); // Get the last two digits of the year
    return `${month}/${year}`;
  };

  return (
    <div className="debit-card">
      <img src={`${process.env.PUBLIC_URL}/logo192.png`} alt="Bank Logo" className="logo" />
      <div className="debit-text">DEBIT</div>
      <div className="card-number">{number}</div>
      <div className="card-holder">
        <div>{name}</div>
        <div>{`${formatMMYY(issuedDate)} - ${formatMMYY(expirationDate)}`}</div>
      </div>
    </div>
  );
};

export default DebitCard;
