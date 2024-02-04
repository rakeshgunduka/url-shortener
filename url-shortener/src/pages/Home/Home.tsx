import { Button, Container, Grid, TextField, Typography } from '@material-ui/core';
import { AxiosError, AxiosResponse } from 'axios';
import React, { useState } from 'react';
import { createShortUrl } from '../../services/shorturl';

const Home = (): React.ReactElement => {
  const [url, setUrl] = useState<string>('');
  const [shortUrl, setShortUrl] = useState<string>('');
  const [shortUrlGenInProgress, setShortUrlGenInProgress] = useState<boolean>(false);

  const handleUrlChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  const handleSubmit = () => {
    setShortUrlGenInProgress(true)
    createShortUrl({ url }).then((response: AxiosResponse) => {
      const { alias } = response.data;
      setShortUrl(`${process.env.REACT_APP_BE_BASE_URL}/${alias}`);
    }).catch((error: AxiosError) => {
      console.error(error);
    }).finally(() => {
      setShortUrlGenInProgress(false)
    });
  };

  return (
    <Container maxWidth="md">
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <Typography variant="h4" align="center" gutterBottom>
            URL Shortener
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <TextField
            label="URL"
            value={url}
            onChange={handleUrlChange}
          />
          <Button variant="contained" color="primary" onClick={handleSubmit} disabled={shortUrlGenInProgress}>
            {shortUrlGenInProgress ? 'Generating...' : 'Submit'}
          </Button>
            {shortUrl && <p>Your shortened URL: <a target="_blank" href={shortUrl} rel="noreferrer">{shortUrl}</a></p>}
        </Grid>
      </Grid>
    </Container>
  );
};

export default Home;
