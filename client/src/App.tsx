import "./App.css";
import Home from "./pages/Home";
import LoginRegister from "./pages/LoginRegister";
import Nav from "./components/Nav";
import theme from "./Theme";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ThemeProvider } from "@emotion/react";
import { useEffect, useState } from "react";
import Profile from "./pages/Profile";
export const ENDPOINT = "http://localhost:8080";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [moderatedSubforums, setModeratedSubforums] = useState([""]);
  const [posts, setPosts] = useState([""]);
  const [comments, setComments] = useState([""]);
  const [votes, setVotes] = useState([""]);
  const [createdAt, setCreatedAt] = useState(Date.now);

  const user: User = {
    username: { value: username, setter: setUsername },
    password: { value: password, setter: setPassword },
    moderatedSubforums: {
      value: moderatedSubforums,
      setter: setModeratedSubforums,
    },
    posts: { value: posts, setter: setPosts },
    comments: { value: comments, setter: setComments },
    votes: { value: votes, setter: setVotes },
    createdAt: { value: createdAt, setter: setCreatedAt },
  };

  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8080/api/authuser", {
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      const content = await response.json();

      user.username.setter(content.data.Username);
      user.password.setter(content.data.Password);
      user.moderatedSubforums.setter(content.data.moderatedSubforums);
      user.posts.setter(content.data.Posts);
      user.comments.setter(content.data.Comments);
      user.votes.setter(content.data.Votes);
      user.votes.setter(content.data.createdAt);
    })();
  });

  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Nav
          username={user.username.value}
          setUsername={user.username.setter}
        />
        <main>
          <BrowserRouter>
            <Routes>
              <Route
                path="/"
                element={<Home username={user.username.value} />}
              />
              <Route
                path="/login"
                element={
                  <LoginRegister
                    type="Login"
                    setUsername={user.username.setter}
                  />
                }
                key="login"
              />
              <Route
                path="/register"
                element={<LoginRegister type="Register" />}
                key="register"
              />
              <Route path="/profile" element={<Profile user={user} />} />
            </Routes>
          </BrowserRouter>
        </main>
      </div>
    </ThemeProvider>
  );
}

export default App;
