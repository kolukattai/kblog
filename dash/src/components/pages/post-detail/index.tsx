import { PostService } from "@/services/posts";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

export const PostDetail = () => {

    const router = useRouter()
    const [data, setData] = useState([]);

    const init = () => {
        PostService.getOne(`${router.query.slug}`).then(res => {
            setData(res.data);
        }).catch(err => {
            console.error(err);
        })
    }

    useEffect(() => {
        init()
    }, [])

    return (
        <div>
            {JSON.stringify(data)}
        </div>
    )
}