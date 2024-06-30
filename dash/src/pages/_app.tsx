import { Drawer } from "@/components/drawer";
import "@/styles/globals.scss";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  const render = (children: any) => {
    let ComponentA: any = Component
    switch (ComponentA) {
      case "plain":
        return children
      default:
        return <main className="default-layout">
          <Drawer />
          <div className="main">
            {children}
          </div>
        </main>
    }
  }

  return render(<Component {...pageProps} />);
}
