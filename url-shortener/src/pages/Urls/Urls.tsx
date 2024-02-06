import { Box, Button, CircularProgress, Container, Grid, Typography } from '@material-ui/core';
import { AxiosError, AxiosResponse } from 'axios';
import MUIDataTable, { SelectableRows } from 'mui-datatables';
import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import { getAllEvents } from '../../services/events';
import { getUrlLists } from '../../services/shorturl';

interface IUrlData {
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  ID: number;
  Alias: string;
  URL: string;
  Created: string;
  Expiry: string;
  Active: boolean;
}

interface AnalyticsData {
  clickEvents: number;
  submitEvents: number;
  viewEvents: number;
}

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

const StyledMUIDataTable = styled(MUIDataTable)`
  width: 100%;
  height: 100%;
`;

const StyledBox = styled(Box)`
  display: inline-flex;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 10px;
  margin-right: 10px;
  margin-bottom: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
`;

const Urls = (): React.ReactElement => {
  const [urls, setUrls] = useState<IUrlData[] | null>();
  const [analytics, setAnalytics] = useState<AnalyticsData | null>();
  const [loadingData, setLoadingData] = useState<boolean>(false);
  const [loadingAnalyticsData, setLoadingAnalyticsData] = useState<boolean>(false);

  const loadData = () => {
    setLoadingData(true);
    getUrlLists()
      .then((response: AxiosResponse) => {
        const { urls } = response.data;
        setUrls(urls);
      })
      .catch((error: AxiosError) => {
        console.error(error);
      })
      .finally(() => {
        setLoadingData(false);
      });
    setLoadingAnalyticsData(true)
    getAllEvents()
      .then((response: AxiosResponse) => {
        const { events } = response.data;
        setAnalytics(events);
      })
      .catch((error: AxiosError) => {
        console.error(error);
      })
      .finally(() => {
        setLoadingAnalyticsData(false);
      });
  };


  useEffect(() => {
    loadData();
  }, []);

  const columns = [
    {
      name: "Alias",
      label: "Alias",
      options: {
        filter: true,
        sort: true,
      },
    },
    {
      name: "URL",
      label: "URL",
      options: {
        filter: true,
        sort: true,
      },
    },
    {
      name: "Created",
      label: "Created",
      options: {
        filter: true,
        sort: true,
      },
    },
    {
      name: "Expiry",
      label: "Expiry",
      options: {
        filter: true,
        sort: true,
      },
    },
  ];

  const options = {
    selectableRows: 'none' as SelectableRows,
  };

  return (
    <StyledContainer maxWidth="md">
      <StyledHeader>
        <Typography variant="h3">
          URL Shortener
        </Typography>
      </StyledHeader>
      <StyledNavigation>
        <Button variant="outlined" onClick={() => window.location.href = '/urlify/'}>Home</Button>
        <Button variant="outlined" onClick={() => window.location.href = '/urlify/urls'}>All URLS</Button>
      </StyledNavigation>
      {loadingAnalyticsData ? (
        <Grid item xs={12}>
          <CircularProgress />
        </Grid>
      ) : (
        analytics &&
        <Grid item xs={12}>
          <StyledBox>
            <Typography variant="h6">Click Events:</Typography>
            <Typography variant="h5" style={{ marginLeft: '5px' }}>
              {analytics.clickEvents}
            </Typography>
          </StyledBox>
          <StyledBox>
            <Typography variant="h6">View Events:</Typography>
            <Typography variant="h5" style={{ marginLeft: '5px' }}>{analytics.viewEvents}</Typography>
          </StyledBox>
          <StyledBox>
            <Typography variant="h6">Submit Events:</Typography>
            <Typography variant="h5" style={{ marginLeft: '5px' }}>{analytics.submitEvents}</Typography>
          </StyledBox>
        </Grid>
      )}
      {loadingData ? (
        <Grid item xs={12}>
          <CircularProgress />
        </Grid>
      ) : (
        <Grid item xs={12}>
          {urls && urls.length ? (
            <StyledMUIDataTable title={'All Short Urls'} data={urls} columns={columns} options={options} />
          ) : (
            <h3> Loading..!</h3>
          )}
        </Grid>
      )}
    </StyledContainer>
  );
};

export default Urls;
