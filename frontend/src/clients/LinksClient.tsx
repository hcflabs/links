import axios, { AxiosInstance, AxiosRequestConfig } from "axios";
// import { Forbidden, HttpError, Unauthorized } from "../errors";
// import { Headers } from "../types";

export class LinksApiClient {
  constructor(
    private readonly baseUrl: string,
    private readonly headers: Headers,
    private readonly authToken: string = ""
  ) {}

  // public async getLink(
  //   endpoint: string = "",
  //   params?: any,
  //   signal?: AbortSignal
  // ): Promise<any> {
  //   try {
  //     const client = this.createClient(params);
  //     const response = await client.get(endpoint, { signal });
  //     return response.data;
  //   } catch (error: any) {
  //     this.handleError(error);
  //   }
  // }

  // public async postNewLink(
  //   endpoint: string = "",
  //   data?: any,
  //   signal?: AbortSignal
  // ): Promise<any> {
  //   try {
  //     const client = this.createClient();
  //     const response = await client.post(endpoint, data, { signal });
  //     return response.data;
  //   } catch (error) {
  //     this.handleError(error);
  //   }
  // }

  // public async uploadFile(
  //   endpoint: string = "",
  //   formData: FormData
  // ): Promise<any> {
  //   try {
  //     const client = this.createClient();
  //     const response = await client.post(endpoint, formData, {
  //       headers: {
  //         "Content-Type": "multipart/form-data",
  //       },
  //     });
  //     return response.data;
  //   } catch (error) {
  //     this.handleError(error);
  //   }
  // }

  // private createClient(params: object = {}): AxiosInstance {
  //   const config: AxiosRequestConfig = {
  //     baseURL: this.baseUrl,
  //     headers: this.headers,
  //     params: params,
  //   };
  //   if (this.authToken) {
  //     config.headers = {
  //       Authorization: `Bearer ${this.authToken}`,
  //     };
  //   }
  //   return axios.create(config);
  // }

  // private handleError(error: any): never {
  //   // if (!error.response) {
  //   //   throw new HttpError(error.message);
  //   // } else if (error.response.status === 401) {
  //   //   throw new Unauthorized(error.response.data);
  //   // } else if (error.response.status === 403) {
  //   //   throw new Forbidden(error.response.data);
  //   // } else {
  //   //   throw error;
  //   // }
  // }
}
