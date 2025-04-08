import React from 'react';
import './CreditCard.css';

const CreditCard = ({ card }) => {
  const { card_type, name, number, issued_date, expiration_date } = card;

  let backgroundColor = 'navy';
  let textColor = 'gold';
  if (card_type.includes('Silver')) {
    backgroundColor = 'silver';
    textColor = 'darkgrey';
  } else if (card_type.includes('Gold')) {
    backgroundColor = 'gold';
    textColor = 'white';
  }

  let logoSrc = '';
  const lowerCardType = card_type.toLowerCase();
  if (lowerCardType.includes('visa')) {
    logoSrc = `${process.env.PUBLIC_URL}/logo-visa.png`;
  } else if (lowerCardType.includes('mastercard')) {
    logoSrc = `${process.env.PUBLIC_URL}/logo-mc.png`;
  } else if (lowerCardType.includes('american express')) {
    logoSrc = `${process.env.PUBLIC_URL}/logo-ae.png`;
  } else if (lowerCardType.includes('discover')) {
    logoSrc = `${process.env.PUBLIC_URL}/logo-disc.png`;
  } else if (lowerCardType.includes('diners club')) {
    logoSrc = `${process.env.PUBLIC_URL}/logo-dc.png`;
  }

  const formatMMYY = (date) => {
    const month = String(date.month).padStart(2, '0');
    const year = String(date.year).slice(-2);
    return `${month}/${year}`;
  };

  return (
    <div className="credit-card" style={{ backgroundColor, color: textColor }}>
      <img src={logoSrc} alt="Card Logo" className="logo" />
      <div className="card-number">{number}</div>
      <div className="card-holder">
        <div>{name}</div>
        <div>{`${formatMMYY(issued_date)} - ${formatMMYY(expiration_date)}`}</div>
      </div>
    </div>
  );
};

export default CreditCard;
