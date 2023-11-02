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
export default function MyComponent() {
  return <h1>Hello</h1>;
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
