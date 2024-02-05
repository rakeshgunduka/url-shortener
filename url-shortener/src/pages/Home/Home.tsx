
import { Box, Button, Container, TextField, Typography } from '@material-ui/core';
import { AxiosError, AxiosResponse } from 'axios';
import React, { useState } from 'react';
import styled from 'styled-components';
import { createShortUrl } from '../../services/shorturl';

const StyledContainer = styled(Container)`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
`;

const StyledHeader = styled(Box)`
  text-align: center;
  margin-bottom: 20px;
`;

const StyledNavigation = styled(Box)`
  text-align: center;
  margin-bottom: 20px;
  
`;

const StyledBox = styled(Box)`
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  margin: auto;

  @media (max-width: 600px) {
    max-width: 80%;
  }
`;


const Home = (): React.ReactElement => {
  const [url, setUrl] = useState<string>('');
  const [shortUrl, setShortUrl] = useState<string>('');
  const [shortUrlGenInProgress, setShortUrlGenInProgress] = useState<boolean>(false);

  const handleUrlChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  const handleSubmit = () => {
    setShortUrlGenInProgress(true);
    createShortUrl({ url })
      .then((response: AxiosResponse) => {
        const { alias } = response.data;
        setShortUrl(`${process.env.REACT_APP_BE_BASE_URL}/${alias}`);
      })
      .catch((error: AxiosError) => {
        console.error(error);
      })
      .finally(() => {
        setShortUrlGenInProgress(false);
      });
  };

  return (
    <StyledContainer maxWidth="md">
      <StyledHeader>
        <Typography variant="h3">
          URL Shortener
        </Typography>
      </StyledHeader>
      <StyledNavigation>
        <Button variant="outlined" onClick={() => window.location.href = '/'}>Home</Button>
        <Button variant="outlined" onClick={() => window.location.href = '/urls'}>Text</Button>
      </StyledNavigation>
      <StyledBox>
        <TextField
          label="URL"
          value={url}
          onChange={handleUrlChange}
          fullWidth
          margin="normal"
        />
        <Button
          variant="contained"
          color="primary"
          onClick={handleSubmit}
          disabled={shortUrlGenInProgress}
          fullWidth
          style={{ marginTop: '16px' }}
        >
          {shortUrlGenInProgress ? 'Generating...' : 'Submit'}
        </Button>
        {shortUrl && (
          <p style={{ marginTop: '16px' }}>
            Your shortened URL: <a target="_blank" href={shortUrl} rel="noreferrer">{shortUrl}</a>
          </p>
        )}
      </StyledBox>
    </StyledContainer>
  );
};

export default Home;
