import axios, { AxiosHeaders, AxiosInstance, AxiosRequestConfig } from "axios";
import { Link, NewLink } from "@/defs/link_pb";

export class LinkApiv1 {
  prefix = "/api/v1";
  constructor(private readonly baseUrl: string) {}

  public async getLink(
    link: string = "",
    params?: any,
    signal?: AbortSignal
  ): Promise<Link | undefined> {
    try {
      const client = this.createClient(params);
      const response = await client.get(`${this.prefix}/links/${link}`, { signal });
      return response.data;
    } catch (error: any) {
      this.handleError(error);
    }
  }

  public async postNewLink(link: NewLink, signal?: AbortSignal): Promise<any> {
    try {
      const client = this.createClient();
      const response = await client.post(
        `${this.prefix}/links/${link.getShortUrl}`,
        link,
        { signal }
      );
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  public async deleteLink(link: string, signal?: AbortSignal): Promise<any> {
    try {
      const client = this.createClient();
      const response = await client.delete(
        `${this.prefix}/links/${link}`,
        { signal }
      );
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  private createClient(params: object = {}): AxiosInstance {
    const config: AxiosRequestConfig = {
      baseURL: this.baseUrl,
      headers: new AxiosHeaders(),
      params: { params },
    };
    // if (this.authToken) {
    //   config.headers = {
    //     Authorization: `Bearer ${this.authToken}`,
    //   };
    // }
    return axios.create(config);
  }

  private handleError(error: any): never {
    if (!error.response) {
      throw new Error(error.message);
    } else if (error.response.status === 401) {
      throw new Error(error.response.data);
    } else if (error.response.status === 403) {
      throw new Error(error.response.data);
    } else {
      throw error;
    }
  }
}
