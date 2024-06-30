import Link from "next/link"
import { useRouter } from "next/router"

export const Drawer = () => {

    const links = [
        {label: "Home", path: "/"},
        {label: "Posts", path: "/posts"},
        {label: "Scheduled", path: "/scheduled"},
    ]

    const router = useRouter()

    return (
        <aside className="drawer">
            <ul className="pt-5">
                {links.map((e, i) => (
                    <li key={i}>
                        <Link 
                        className={`inline-block p-2 px-7 hover:bg-black hover:text-white w-full m-1 rounded-sm ${router.pathname == e.path ? 'bg-gray-800 text-white' : ''}`}
                        href={e.path}>{e.label}</Link>
                    </li>
                ))}
            </ul>
        </aside>
    )
}


