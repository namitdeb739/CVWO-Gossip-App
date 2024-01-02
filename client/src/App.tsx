import "./App.css";
import Home from "./pages/Home";
import LoginRegister from "./pages/LoginRegister";
import Nav from "./components/Nav";
import theme from "./Theme";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ThemeProvider } from "@emotion/react";
import { useEffect, useState } from "react";
import Profile from "./pages/Profile";
import UserPage from "./pages/UserPage";
import SubforumPage from "./pages/SubforumPage";
export const ENDPOINT = "http://localhost:8080";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [moderatedSubforums, setModeratedSubforums] = useState<Subforum[]>([]);
  const [posts, setPosts] = useState<Post[]>([]);
  const [comments, setComments] = useState<Comment[]>([]);
  const [votes, setVotes] = useState<Vote[]>([]);
  const [createdAt, setCreatedAt] = useState(new Date);

  const user: User = {
    Username: username,
    Password: password,
    ModeratedSubforums: moderatedSubforums,
    Posts: posts,
    Comments: comments,
    Votes: votes,
    CreatedAt: createdAt,
  };

  useEffect(() => {
    (async () => {
      const response = await fetch(ENDPOINT + "/api/authuser", {
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      const content = await response.json();

      setUsername(content.data.Username);
      setPassword(content.data.Password);
      setModeratedSubforums(content.data.Moderated_Subforums);
      setPosts(content.data.Posts);
      setComments(content.data.Comments);
      setVotes(content.data.Votes);
      setCreatedAt(new Date(content.data.CreatedAt));
    })();
  });

  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Nav username={user.Username} setUsername={setUsername} />
        <main>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/profile" element={<Profile user={user} />} />
              <Route path="/user/:ID" element={<UserPage />} />
              <Route path="/subforum/:ID" element={<SubforumPage />} />
              <Route
                path="/login"
                element={
                  <LoginRegister type="Login" setUsername={setUsername} />
                }
                key="login"
              />
              <Route
                path="/register"
                element={<LoginRegister type="Register" />}
                key="register"
              />
            </Routes>
          </BrowserRouter>
        </main>
      </div>
    </ThemeProvider>
  );
}

export default App;
