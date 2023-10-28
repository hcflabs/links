import Image from "next/image";
async function getPost(link: string) {
  console.log("->RUN 3", link);
  const res = await fetch(
    `http://host.docker.internal:8080/api/v1/links/${link}`
  );

  return res.json();
}

export default async function Page({ params }: { params: { link: string } }) {
  console.log("->RUN 2", params);
  const data = await getPost(params.link);
  return (
    <>
      <span>Title:{data.result.title}</span>
    </>
  );
}
export async function generateStaticParams() {
  console.log("->RUN 1");
  const posts = await fetch("http://host.docker.internal:8080/api/v1/links/holdon").then(
    (res) => res.json()
  );
  return posts.data.map((post: { id: { toString: () => any } }) => ({
    id: post.id.toString(),
  }));
}
