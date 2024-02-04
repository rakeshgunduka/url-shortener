import { CircularProgress, Container, Grid, Typography } from '@material-ui/core';
import { AxiosError, AxiosResponse } from 'axios';
import MUIDataTable, { SelectableRows } from 'mui-datatables';
import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
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

const StyledMUIDataTable = styled(MUIDataTable)`
  width: 100%;
  height: 100%;
`;


const Urls = (): React.ReactElement => {
  const [urls, setUrls] = useState<IUrlData[] | null>();
  const [loadingData, setLoadingData] = useState<boolean>(false);

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
    {
      name: "Active",
      label: "Active",
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
          <Grid item xs={12}>
            {urls && urls.length ? (
              <StyledMUIDataTable title={'Your Reportees'} data={urls} columns={columns} options={options} />
            ) : (
              <h3> Loading..!</h3>
            )}
          </Grid>
        )}
      </Grid>
    </Container>
  );
};

export default Urls;
