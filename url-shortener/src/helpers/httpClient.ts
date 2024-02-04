import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

export const eHttpMethod = {
  GET: 'get',
  POST: 'post',
  PUT: 'put',
  DELETE: 'delete',
  PATCH: 'patch',
};

export interface AxiosParams {
  path: string;
  method: string;
  body?: unknown;
  params?: unknown;
}

class HttpClient {
  BE_BASE_URL = process.env.REACT_APP_BE_BASE_URL + '/app/api';

  async get(path: string, params?: unknown) {
    return this.axiosBeRequest({ path, method: eHttpMethod.GET, params });
  }

  async post(path: string, body: unknown, params?: unknown) {
    return this.axiosBeRequest({ path, method: eHttpMethod.POST, body, params });
  }

  async delete(path: string, params?: unknown) {
    return this.axiosBeRequest({ path, method: eHttpMethod.DELETE, params });
  }

  async put(path: string, body: unknown, params?: unknown) {
    return this.axiosBeRequest({ path, method: eHttpMethod.PUT, body, params });
  }

  async patch(path: string, body: unknown, params?: unknown) {
    return this.axiosBeRequest({ path, method: eHttpMethod.PATCH, body, params });
  }

  private async axiosBeRequest({ path, method, body, params }: AxiosParams) {
    try {
      const url = this.BE_BASE_URL + path;
      const headers = {};

      const config: AxiosRequestConfig = {
        method,
        url,
        data: body,
        params,
        headers,
        timeout: 30000,
      };
      const response: AxiosResponse = await axios(config);
      return response;
    } catch (error) {
      throw error;
    }
  }
}

const httpClient = new HttpClient();
export default httpClient;
