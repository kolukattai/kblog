import Image from "next/image";
import { Inter } from "next/font/google";
import { HomePage } from "@/components/pages/home";
import { useEffect } from "react";
import { PostService } from "@/services/posts";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {


  return (
    <div>
      <HomePage />
    </div>
  );
}
