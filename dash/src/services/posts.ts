import { HttpClient } from "./http";

class Post extends HttpClient {
    constructor() {
        super();
    }
    async getMetaData() {
        return this.http.get("/posts")
    }
    
    async getOne(slug: string) {
        return this.http.get(`/posts/${slug}`)
    }
}




export const PostService = new Post()