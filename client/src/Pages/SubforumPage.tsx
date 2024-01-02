import { Box, Stack, Button, Typography, Pagination } from "@mui/material";
import React, { useEffect, useState } from "react";
import { ENDPOINT } from "../App";
import PostCard from "../components/PostCard";
import { useLocation, useParams } from "react-router-dom";

function SubforumPage() {
  const { ID } = useLocation().pathname==="/" ? {ID: ""} : useParams();
  const [subforum, setSubforum] = useState<Subforum>({
    Name: "",
    Description: "",
    Moderators: [],
    Posts: [],
    CreatedAt: new Date(),
  });

  const [currentPage, setCurrentPage] = useState<number>(1);
  const postsPerPage = 5;

  useEffect(() => {
    (async () => {
      const response = await fetch(ENDPOINT + "/api/subforum/" + ID, {
        headers: { "Content-Type": "application/json" },
      });

      const content = await response.json();

      setSubforum(content.data);
    })();
  }, [ID]);

  const indexOfLastPost = currentPage * postsPerPage;
  const indexOfFirstPost = indexOfLastPost - postsPerPage;
  const currentPosts = subforum.Posts.slice(indexOfFirstPost, indexOfLastPost);

  const handlePageChange = (
    _event: React.ChangeEvent<unknown>,
    value: number
  ) => {
    setCurrentPage(value);
  };

  return (
    <div className="container">
      <div className="header">
        <div className="text">NUS Forum / {subforum.Name}</div>
        <div className="underline"></div>
      </div>
      <Box sx={{ width: "100%", marginTop: "20px"}}>
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
          {subforum.Posts.map((post, index) => (
            <PostCard key={index} post={post} />
          ))}
        </Stack>
      </Box>

      <Pagination
        count={Math.ceil(subforum.Posts.length / postsPerPage)}
        page={currentPage}
        color="secondary"
        onChange={handlePageChange}
      />
    </div>
  );
}

export default SubforumPage;
