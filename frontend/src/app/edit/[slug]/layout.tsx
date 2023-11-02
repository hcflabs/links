// import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function editLayout({
  children, // will be a page or nested layout
}: {
  children: React.ReactNode;
}) {
  return <div>{children}</div>;
}
