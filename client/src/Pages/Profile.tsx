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
  padding: theme.spacing(1),
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

  useEffect(() => {
    setUsername(props.user.username.value);
    setPassword(props.user.password.value);
    setModeratedSubforums(props.user.moderatedSubforums.value);
    setPosts(props.user.posts.value);
    setComments(props.user.comments.value);
    setVotes(props.user.votes.value);
    setCreatedAt(new Date(props.user.createdAt.value));
  }, [props.user]);

  return (
    <div className="container">
      <div className="header">
        <div className="text">{username}</div>
        <div className="underline"></div>
      </div>
      <Grid container spacing={2} marginTop={2}>
        <Grid item xs={4}>
          <Item>
            <u>
              <b>Date joined:</b>
            </u>{" "}
            {formatDate(createdAt)}
          </Item>
        </Grid>
        <Grid item xs={8}>
          <Item>xs=4</Item>
        </Grid>
        <Grid item xs={4}>
          <Item className="passwordContainer">
            <u>
              <b>Password:</b>
            </u>{" "}
            {hidePassword ? "*".repeat(password.length) : password}
            <img className="eyeIcon" src={eye_icon} alt="" onClick={setHidePassword(!hidePassword)} />
          </Item>
        </Grid>
        <Grid item xs={8}>
          <Item>xs=8</Item>
        </Grid>
      </Grid>
    </div>
  );
}

export default Profile;
