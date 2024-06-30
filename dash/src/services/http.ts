import axios, { AxiosInstance } from "axios";

export class HttpClient {
    protected http!: AxiosInstance
    constructor() {
        this.http = axios.create({
            baseURL: "http://localhost:8080/api"
        })
    }
}