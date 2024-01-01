import { useEffect, useState } from "react";
import "./Home.css";
import PostCard from "./PostCard";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Pagination from "@mui/material/Pagination";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

function Home() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const postsPerPage = 5;

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
  }, []); // Add an empty dependency array to useEffect

  const indexOfLastPost = currentPage * postsPerPage;
  const indexOfFirstPost = indexOfLastPost - postsPerPage;
  const currentPosts = posts.slice(indexOfFirstPost, indexOfLastPost);

  const handlePageChange = (
    _event: React.ChangeEvent<unknown>,
    value: number
  ) => {
    setCurrentPage(value);
  };

  return (
    <div className="container">
      <Box sx={{ width: "100%" }}>
        <Stack spacing={1}>
          <div
            style={{
              display: "flex",
              justifyContent: "center",
              width: "100%",
              marginBottom: "30px",
            }}
          >
            <Button
              variant="contained"
              className="submit"
              color="secondary"
              sx={{ width: "750px" }}
            >
              <Typography
                variant="h3"
                noWrap
                // component="a"
                // href="/"
                sx={{
                  flexGrow: 0,
                  fontSize: 20,
                  fontWeight: 900,
                  letterSpacing: "0rem",
                  color: "#000",
                  textDecoration: "none",
                }}
              >
                Make Post
              </Typography>
            </Button>
          </div>
          {currentPosts.map((post, index) => (
            <PostCard key={index} post={post} />
          ))}
        </Stack>
      </Box>

      <Pagination
        count={Math.ceil(posts.length / postsPerPage)}
        page={currentPage}
        color="secondary"
        onChange={handlePageChange}
      />
    </div>
  );
}

export default Home;
