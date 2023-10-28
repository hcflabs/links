import { LinksApiClient } from "./LinksClient";


export class ApiClientFactory {
    constructor(
        private readonly baseUrl: string,
        // private readonly headers: Headers = {}
    ) {}

    // public createClient(): LinksApiClient {
    //     return new LinksApiClient(this.baseUrl, this.headers);
    // }

    // public createAuthorizedClient(authToken: string): LinksApiClient {
    //     return new LinksApiClient(this.baseUrl, this.headers, authToken);
    // }
}