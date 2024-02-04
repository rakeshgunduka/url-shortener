import { CircularProgress, Container, Grid, Typography } from '@material-ui/core';
import { AxiosError, AxiosResponse } from 'axios';
import React, { useEffect, useState } from 'react';
import { getAllEvents } from '../../services/events';

interface AnalyticsData {
  clickEvents: number;
  submitEvents: number;
  viewEvents: number;
}

const Analytics = (): React.ReactElement => {
  const [analytics, setAnalytics] = useState<AnalyticsData | null>();
  const [loadingData, setLoadingData] = useState<boolean>(false);

  const loadData = () => {
    setLoadingData(true);
    getAllEvents()
      .then((response: AxiosResponse) => {
        const { events } = response.data;
        setAnalytics(events);
      })
      .catch((error: AxiosError) => {
        console.error(error);
      })
      .finally(() => {
        setLoadingData(false);
      });
  };

  useEffect(() => {
    loadData();
  }, []);

  return (
    <Container maxWidth="md">
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <Typography variant="h4" align="center" gutterBottom>
            URL Shortener
          </Typography>
        </Grid>
        {loadingData ? (
          <Grid item xs={12}>
            <CircularProgress />
          </Grid>
        ) : (
          analytics &&
          <Grid item xs={12}>
            <div>
              <p>Click Events: {analytics.clickEvents}</p>
              <p>View Events: {analytics.viewEvents}</p>
              <p>Submit Events: {analytics.submitEvents}</p>
            </div>
          </Grid>
        )}
      </Grid>
    </Container>
  );
};

export default Analytics;
