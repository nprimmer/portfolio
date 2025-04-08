import React from 'react';
import './DriversLicense.css';

const DriversLicense = ({ license, endpoint }) => {
  const { name, number, expiration_date, address, birthdate } = license;
  return (
    <div className="drivers-license">
      <div className="license-header">
        <img src={`${process.env.PUBLIC_URL}/headshot-${endpoint}.png`} alt="headshot" className="headshot" />
        <div className="license-info">
          <div className="license-title">DRIVER'S LICENSE</div>
          <div className="license-number">License No: {number}</div>
        </div>
      </div>
      <div className="license-body">
        <div className="license-section">
          <div className="label">Name:</div>
          <div className="value">{name}</div>
        </div>
        <div className="license-section">
          <div className="label">DOB:</div>
          <div className="value">{`${birthdate.month}/${birthdate.day}/${birthdate.year}`}</div>
        </div>
        <div className="license-section">
          <div className="label address-label">Address:</div>
        </div>
        <div className="license-section">
          <div className="address">{`${address.number} ${address.street} ${address.street_type},`}</div>
        </div>
        <div className="license-section">
          <div className="address">{`${address.city}, ${address.state} ${address.zip}`}</div>
        </div>
      </div>
    </div>
  );
};

export default DriversLicense;
