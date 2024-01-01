import { useEffect, useState } from "react";
import "./Profile.css";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import VisibilityIcon from "@mui/icons-material/Visibility";

const Item = styled(Paper)(({ theme }) => ({
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  fontSize: 19,
  boxShadow: "none",
  backgroundColor: "var(--linkwater)",
  transition: "background-color 0.1s",
}));

function Profile(props: { user: User }) {
  const [username, setUsername] = useState(props.user.Username);
  const [password, setPassword] = useState(props.user.Password);
  const [moderatedSubforums, setModeratedSubforums] = useState(
    props.user.ModeratedSubforums
  );
  const [posts, setPosts] = useState(props.user.Posts);
  const [votes, setVotes] = useState(props.user.Votes);
  const [createdAt, setCreatedAt] = useState(props.user.CreatedAt);
  const [hidePassword, setHidePassword] = useState(true);

  const [recentPosts, setRecentPosts] = useState<Post[]>([]);

  useEffect(() => {
    (async () => {
      setUsername(props.user.Username);
      setPassword(props.user.Password);
      setModeratedSubforums(props.user.ModeratedSubforums);
      setPosts(props.user.Posts);
      setVotes(props.user.Votes);
      setCreatedAt(props.user.CreatedAt);

      const sortedPosts =
        props.user.Posts && props.user.Posts.length > 0
          ? props.user.Posts.sort(
              (a, b) =>
                new Date(b.CreatedAt.toString()).valueOf() -
                new Date(a.CreatedAt.toString()).valueOf()
            )
          : [];

      setRecentPosts(sortedPosts.slice(0, 10));
    })();
  }, [props.user]);

  const formatDate = (date: Date) => {
    return date.toISOString().slice(0, 19).replace("T", " ");
  };

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
          <div className="dateJoinedBox">
            <div className="dataContainer">
              <div className="label">Date Joined:</div>
              <div className="data">{formatDate(createdAt)}</div>
            </div>
          </div>
        </Grid>
        <Grid item xs={3.5}>
          <div className="passwordBox">
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
                  <VisibilityIcon />
                </a>
              </div>
            </div>
          </div>
        </Grid>
        <Grid item xs={3.5}>
          <div className="voteBox">
            <div className="dataContainer">
              <div className="label">Vote Count:</div>
              <div className="data">
                {votes && votes.length > 0 ? (
                  <div>
                    {votes.filter((vote) => vote.Type === true).length} Up /{" "}
                    {votes.filter((vote) => vote.Type === false).length} Down
                  </div>
                ) : (
                  <div>0 Up / 0 Down</div>
                )}
              </div>
            </div>
          </div>
        </Grid>
        <Grid item xs={12}></Grid>
        <Grid item xs={5}>
          <Item>
            <div className="dataListContainer">
              <div className="label">Moderated Subforums</div>
              <Box
                sx={{
                  display: "grid",
                  gridTemplateRows: "repeat(10, 1fr)",
                }}
              >
                {moderatedSubforums && moderatedSubforums.length > 0 ? (
                  moderatedSubforums.map((subforum, index) => (
                    <Item key={index} className="custom-hover-effect">
                      <Box
                        sx={{
                          display: "grid",
                          gridTemplateColumns: "repeat(2, 1fr)",
                        }}
                        className="hoverBox"
                      >
                        <Item className="hoverBox">
                          <div className="dataHeader">{subforum.Name}</div>
                        </Item>
                        <Item className="hoverBox">
                          <div className="data">
                            {subforum.Description === undefined
                              ? ""
                              : subforum.Description.length > 15
                              ? subforum.Description.slice(0, 15) + "..."
                              : subforum.Description}
                          </div>
                        </Item>
                      </Box>
                    </Item>
                  ))
                ) : (
                  <div />
                )}
              </Box>
            </div>
          </Item>
        </Grid>
        <Grid item xs={7}>
          <Item>
            <div className="dataListContainer">
              <div className="label">Recent Posts</div>
              <Box
                sx={{
                  display: "grid",
                  gridTemplateRows: "repeat(10, 1fr)",
                }}
              >
                {recentPosts && recentPosts.length > 0 ? (
                  recentPosts.map((post, index) => (
                    <Item key={index} className="custom-hover-effect">
                      <Box
                        className="hoverBox"
                        sx={{
                          display: "grid",
                          gridTemplateColumns: "repeat(3, 1fr)",
                        }}
                      >
                        <Item className="hoverBox">
                          <div className="dataHeader">{post.Title}</div>
                        </Item>
                        <Item className="hoverBox">
                          <div className="data">
                            {post.Body && post.Body.length > 15
                              ? post.Body.slice(0, 15) + "..."
                              : post.Body}
                          </div>
                        </Item>
                        <Item className="hoverBox">
                          <div className="dataDate">
                            {post.CreatedAt
                              ? post.CreatedAt.toString().split("T")[0]
                              : ""}
                          </div>
                        </Item>
                      </Box>
                    </Item>
                  ))
                ) : (
                  <div />
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
