import httpClient from '../helpers/httpClient';

export const getAllEvents = () => {
  return httpClient.get('/events/');
};

export enum EventName {
  Click = "click",
  View = "view",
  Submit = "submit",
}


export const storeEvent = (eventName: EventName) => {
  return httpClient.post(`/events/${eventName}`, {});
}
