import httpClient from '../helpers/httpClient';

export const createShortUrl = (data: {
  url: string;
}) => {
  return httpClient.post("/urls/create", data);
};

export const getUrlLists = () => {
  return httpClient.get("/urls/");
}

export const getOriginalUrl = (shortUrl: string) => {
  return httpClient.get(`/urls/${shortUrl}`);
}
