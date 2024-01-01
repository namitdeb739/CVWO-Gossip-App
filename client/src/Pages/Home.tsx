import { useEffect, useState } from "react";
import "./Home.css";
import PostCard from "./PostCard";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Pagination from "@mui/material/Pagination";

function Home() {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8080/api/post", {
        headers: { "Content-Type": "application/json" },
      });

      const content = await response.json();

      const sortedPosts =
        content.data?.length > 0
          ? [...content.data].sort(
              (a: { CreatedAt: Date }, b: { CreatedAt: Date }) =>
                b.CreatedAt.valueOf() - a.CreatedAt.valueOf()
            )
          : [];
      setPosts(sortedPosts);
    })();
  });

  return (
    <div className="container">
      <Box sx={{ width: "100%" }}>
        <Stack spacing={0.5}>
          {posts === undefined ? (
            <div></div>
          ) : (
            posts.map((post, index) => <PostCard key={index} post={post} />)
          )}
        </Stack>
      </Box>

      <Pagination count={10} color="secondary" />
    </div>
  );
}

export default Home;
