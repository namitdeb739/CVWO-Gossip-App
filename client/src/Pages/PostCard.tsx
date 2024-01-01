import "./PostCard.css";
import Card from "@mui/material/Card";
import CardHeader from "@mui/material/CardHeader";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import Avatar from "@mui/material/Avatar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";
import ThumbDownIcon from "@mui/icons-material/ThumbDown";
import { useState } from "react";
import Chip from "@mui/material/Chip";
import getDataFromID from "../helpers/getDataFromID";
import { Box } from "@mui/material";

function PostCard(props: { post: Post }) {
  const [userID] = useState(props.post.UserID);
  const [subforumID] = useState(props.post.SubforumID);
  const [title] = useState(props.post.Title);
  const [body] = useState(props.post.Body);
  const [votes] = useState(props.post.Votes);
  const [tags] = useState(props.post.Tags);
  const [createdAt] = useState(props.post.CreatedAt);

  const user = getDataFromID<User>(userID, "user");
  const subforum = getDataFromID<Subforum>(subforumID, "subforum");

  const formatDate = (date: Date) => {
    return date.toString().slice(0, 19).replace("T", " ");
  };

  return (
    <div className="container">
      <Card className="card" sx={{ maxWidth: 750 }}>
        <CardHeader
          avatar={<Avatar sx={{ bgcolor: "var(--nusorange)" }}></Avatar>}
          title={
            <b>
              Posted to <u>{subforum ? subforum.Name : ""}</u> by{" "}
              <u>{user ? user.Username : ""}</u>
            </b>
          }
          subheader={<p>{formatDate(createdAt)}</p>}
        />
        <CardContent sx={{ paddingBottom: 0 }}>
          <Typography
            variant="h5"
            fontWeight="bold"
            color="text.secondary"
            sx={{ marginBottom: 2 }}
          >
            {title}
          </Typography>
          <Typography
            variant="body1"
            color="text.secondary"
            sx={{ marginBottom: 3 }}
          >
            {body.length > 50 ? body.slice(0, 15) + "..." : body}
          </Typography>
          <Box sx={{ marginTop: 2 }}>
            {tags && tags.length > 0 ? (
              tags.map((tag, index) => (
                <Chip
                  label={tag.Name}
                  variant="outlined"
                  key={index}
                  sx={{ marginRight: 1, marginBottom: 1 }}
                />
              ))
            ) : (
              <div />
            )}
          </Box>
        </CardContent>
        <CardActions disableSpacing sx={{ paddingTop: 0 }}>
          <IconButton>
            <Typography variant="h6" color="text.secondary">
              {votes && votes.length > 0
                ? votes.filter((vote) => vote.Type === true).length
                : 0}
            </Typography>
            <ThumbUpIcon className="icon" />
          </IconButton>
          <IconButton>
            <Typography variant="h6" color="text.secondary">
              {votes && votes.length > 0
                ? votes.filter((vote) => vote.Type === false).length
                : 0}
            </Typography>
            <ThumbDownIcon className="icon" />
          </IconButton>
        </CardActions>
      </Card>
    </div>
  );
}

export default PostCard;
