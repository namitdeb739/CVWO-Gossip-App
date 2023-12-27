import './App.css';
import Home from "./pages/Home";
import LoginRegister from "./pages/LoginRegister";
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
              <Route path="/"   element={<Home />} />
              <Route path="/login" element={<LoginRegister type="Login" />} />
              <Route path="/register" element={<LoginRegister type="Register" />} />
            </Routes>
          </BrowserRouter>
        </main>
      </div>
    </ThemeProvider>
  );
}

export default App;