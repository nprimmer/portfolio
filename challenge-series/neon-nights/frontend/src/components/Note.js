import React from 'react';
import './Note.css';

const Note = ({ note }) => {
  return (
    <div className="note">
      <div className="note-content">{note.description}</div>
    </div>
  );
};

export default Note;
