// 'use client'
// import { useState, useEffect } from "react";

// export default function Home() {
//   const [data, setData] = useState(null);
//   const [isLoading, setLoading] = useState(true);

//   useEffect(() => {
//     fetch("/api/v1/links")
//       .then((res) => res.json())
//       .then((data) => {
//         setData(data);
//         setLoading(false);
//       });
//   }, []);

//   if (isLoading) return <p>Loading...</p>;
//   if (!data) return <p>No Links Yet?</p>;

//   return (
//     <main className="flex min-h-screen flex-col items-center justify-between p-24">
//       <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
//         Hello World
//         {/* <h1>{data}</h1> */}
//         <p>{data}</p>
//       </div>
//     </main>
//   );
// }

async function getLinkMetadata(link: string) {
  const res = await fetch('http://localhost:8081/api/v1/links' + link)
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }

  return res.json()
}
export default async function MyComponent() {
  const data = await getLinkMetadata("hcf")

  return <h1>{data}</h1>;
}

// export default function NewLinkPage() {
//   const [formData, setFormData] = useState({name: "",email: "",message: ""});

//   const handleChange = (event) => {
//     const { name, value } = event.target;
//     setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
//   };

//   const handleSubmit = (event) => {
//     event.preventDefault();
//     alert(`Name: ${formData.name}, Email: ${formData.email}, Message: ${formData.message}`
//     );
// };

//   return (
//     <form onSubmit={handleSubmit}>
//       <label htmlFor="name">Name:</label>
//       <input type="text" id="name" name="name" value={formData.name} onChange={handleChange}/>

//       <label htmlFor="email">Email:</label>
//       <input type="email" id="email" name="email" value={formData.email} onChange={handleChange}/>

//       <label htmlFor="message">Message:</label>
//       <textarea id="message" name="message" value={formData.message} onChange={handleChange}/>

//       <button type="submit">Submit</button>
//     </form>
//   );
// }
