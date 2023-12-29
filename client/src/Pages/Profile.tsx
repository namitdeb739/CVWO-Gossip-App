import { useEffect, useState } from "react";
import "./Profile.css";
import * as React from "react";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import eye_icon from "./../Assets/Images/Eye.png";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(3),
  textAlign: "center",
  color: theme.palette.text.secondary,
  fontSize: 19,
}));

function Profile(props: { user: User }) {
  const userCreatedAt = new Date(props.user.createdAt.value);

  const formatDate = (date: Date) => {
    return date.toISOString().slice(0, 19).replace("T", " ");
  };

  const [username, setUsername] = useState(props.user.username.value);
  const [password, setPassword] = useState(props.user.password.value);
  const [moderatedSubforums, setModeratedSubforums] = useState(
    props.user.moderatedSubforums.value
  );
  const [posts, setPosts] = useState(props.user.posts.value);
  const [comments, setComments] = useState(props.user.comments.value);
  const [votes, setVotes] = useState(props.user.votes.value);
  const [createdAt, setCreatedAt] = useState(userCreatedAt);
  const [hidePassword, setHidePassword] = useState(true);

  const [recentPosts, setRecentPosts] = useState<any[]>([""]);

  useEffect(() => {
    setUsername(props.user.username.value);
    setPassword(props.user.password.value);
    setModeratedSubforums(props.user.moderatedSubforums.value);
    setPosts(props.user.posts.value);
    setComments(props.user.comments.value);
    setVotes(props.user.votes.value);
    setCreatedAt(new Date(props.user.createdAt.value));

    const sortedPosts =
      posts === undefined
        ? []
        : posts.sort((a, b) => b.CreatedAt - a.CreatedAt);

    // Set the most recent posts
    setRecentPosts(sortedPosts.slice(0, 3));
  }, [props.user]);

  return (
    <div className="container">
      <div className="header">
        <div className="text">{username}</div>
        <div className="underline"></div>
      </div>
      <Grid
        container
        spacing={2}
        justifyContent="center"
        alignItems="center"
        marginTop={2}
      >
        <Grid item xs={3.5}>
          <Item>
            <div className="dataContainer">
              <div className="label">Date Joined:</div>
              <div className="data">{formatDate(createdAt)}</div>
            </div>
          </Item>
        </Grid>
        <Grid item xs={3.5}>
          <Item>
            <div className="dataContainer">
              <div className="label">Password:</div>
              <div className="data">
                {hidePassword ? "•".repeat(password.length) : password}
              </div>
              <div className="icon">
                <a
                  href="#"
                  onClick={() => {
                    setHidePassword(!hidePassword);
                  }}
                >
                  <img src={eye_icon} className="eyeIcon" />
                </a>
              </div>
            </div>
          </Item>
        </Grid>
        <Grid item xs={12}></Grid>
        <Grid item xs={4}>
          {1 > 0 ? <div></div> :
            <Item>
              <div className="dataListContainer">
                <div className="label">Moderated Subforums</div>
                <Box sx={{
                  display: "grid",
                  gridTemplateRows: "repeat(3, 1fr)"
                }}>
                  {moderatedSubforums.map((subforum, index) => (
                    <Item key={index}>
                      <Box
                        className="dataContainer"
                        sx={{
                          display: "grid",
                          gridTemplateColumns: "repeat(2, 1fr)",
                        }}
                      >
                        <Item>
                          <div className="label">{subforum.Name}</div>
                        </Item>
                        <Item>
                          <div className="data">
                            {subforum.Description.length > 15
                              ? subforum.Description.slice(0, 15) + "..."
                              : subforum.Description}
                          </div>
                        </Item>
                      </Box>
                    </Item>
                  ))}
                </Box>
              </div>
            </Item>}
        </Grid>
        <Grid item xs={8}>
          <Item>
            <div className="dataListContainer">
              <div className="label">Recent Posts</div>
              <Box
                sx={{
                  display: "grid",
                  gridTemplateRows: "repeat(3, 1fr)",
                }}
              >
                {recentPosts.length === 0 ? (
                  <div></div>
                ) : (
                  recentPosts.map((post, index) => (
                    <Item key={index}>
                      <Box
                        className="dataContainer"
                        sx={{
                          display: "grid",
                          gridTemplateColumns: "repeat(3, 1fr)",
                        }}
                      >
                        <Item>
                          <div className="label">{post.Title}</div>
                        </Item>
                        <Item>
                          <div className="data">
                            {post.Body.length > 15
                              ? post.Body.slice(0, 15) + "..."
                              : post.Body}
                          </div>
                        </Item>
                        <Item>
                          <div className="data">
                            {post.CreatedAt.split("T")[0]}
                          </div>
                        </Item>
                      </Box>
                    </Item>
                  ))
                )}
              </Box>
            </div>
          </Item>
        </Grid>
      </Grid>
    </div>
  );
}

export default Profile;
