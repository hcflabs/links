// 'use client'
import { useState, useEffect } from "react";
import {Link, LinksMetadata} from "@/defs/link_pb";

async function getData() {
    const res = await fetch('http://localhost:8081/api/v1/links')
    // The return value is *not* serialized
    // You can return Date, Map, Set, etc.

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default async function Home() {
  // const [data, setData] = useState(null);
  // const [isLoading, setLoading] = useState(true);
    const data = await getData()
    console.log(JSON.stringify(data, null, 2))
    const links: LinksMetadata[] = JSON.parse(data)

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <p>{links.map((link) => {
            return (
                <div key={link}>
                    <p>{link}</p>
                    {/* {JSON.stringify(post)} */}
                </div>

            )
        })}</p>
      </div>
    </main>
  );
}
