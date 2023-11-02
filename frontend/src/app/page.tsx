'use client'
import { useState, useEffect } from "react";

export default function Home() {
  const [data, setData] = useState(null);
  const [isLoading, setLoading] = useState(true);

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        Hello World
        {/* <h1>{data}</h1> */}
        <p>{data}</p>
      </div>
    </main>
  );
}
