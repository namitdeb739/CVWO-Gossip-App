import './App.css';
import Login from "./pages/Login";
import Home from "./pages/Home";
import Register from "./pages/Register";
import Nav from "./components/Nav";
import theme from './Theme';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ThemeProvider } from "@emotion/react";

export const ENDPOINT = 'http://localhost:8080'

function App() {
  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Nav />

        <main>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
            </Routes>
          </BrowserRouter>
        </main>
      </div>
    </ThemeProvider>
  );
}

export default App;