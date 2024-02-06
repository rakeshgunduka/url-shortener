import React from 'react';
import { Link } from 'react-router-dom';

const UrlNotFound = (): React.ReactElement => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
      <h1 style={{ fontSize: '4rem', marginBottom: '2rem' }}>404</h1>
      <p style={{ fontSize: '2rem', marginBottom: '2rem' }}>Oops! The page you're looking for doesn't exist.</p>
      <Link to="/" style={{ fontSize: '1.5rem', textDecoration: 'none', color: 'blue' }}>Go back to home</Link>
    </div>
  );
};

export default UrlNotFound;
