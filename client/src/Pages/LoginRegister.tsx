import "./LoginRegister.css";
import user_icon from "./../Assets/Images/User.png";
import key_icon from "./../Assets/Images/Key.png";
import eye_icon from "./../Assets/Images/Eye.png";
import { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";

function LoginRegister({ type }: { type: string }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [hidePassword, setHidePassword] = useState(true);
  const [registerRedirect, setRegisterRedirect] = useState(false);
  const [loginRedirect, setLoginRedirect] = useState(false);

  const isUsernameValid = /^[^\s]+$/.test(username) || username === "";
  const isPasswordValid =
    /^(?=.*[a-z])(?=.*[A-Z])[^\s]{8,}$/.test(password) || password === "";

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    await fetch(
      "http://localhost:8080/api/" + type,
      type === "Register"
        ? {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              username,
              password,
            }),
          }
        : {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "http:localhost:8080/api/login",
            },
            credentials: "include",
            body: JSON.stringify({
              username,
              password,
            }),
          }
    );

    type === "Register" ? setRegisterRedirect(true) : setLoginRedirect(true);
  };

  if (registerRedirect) {
    return <Navigate to={"/login"} replace={true} />;
  }
  if (loginRedirect) {
    return <Navigate to={"/"} replace={true} />;
  }

  return (
    <div className="container">
      <div className="header">
        <div className="text">{type}</div>
        <div className="underline"></div>
      </div>
      <div className="inputs">
        <div className={isUsernameValid ? "input" : "input invalid"}>
          <img src={user_icon} className="inputIcon" />
          <input
            type="Username"
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className={isPasswordValid ? "input" : "input invalid"}>
          <img src={key_icon} className="inputIcon" />
          <input
            className={password.length >= 8 ? "" : "invalid"}
            value={password}
            type={hidePassword ? "password" : "text"}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
          />
          <a
            href="#"
            className="toggle-btn"
            onClick={() => {
              setHidePassword(!hidePassword);
            }}
          >
            <img src={eye_icon} className="inputIcon" />
          </a>
        </div>
      </div>
      <div className="errorMessage">
        {isUsernameValid ? (
          <p></p>
        ) : (
          <p>Invalid Username: Username cannot contain spaces</p>
        )}
      </div>
      <div className="errorMessage">
        {isPasswordValid ? (
          <p></p>
        ) : (
          <p>
            Invalid Password: Password must be at least 8 characters, must have
            upper and lower case letters, cannot contain spaces
          </p>
        )}
      </div>
      <div className="submit-container">
        <button
          className={
            !isUsernameValid || !isPasswordValid ? "disabled" : "submit"
          }
          onClick={submit}
          disabled={!isUsernameValid || !isPasswordValid}
        >
          {type}
        </button>
      </div>
      <div>
        {type === "Login" ? (
          <p>
            No account?{" "}
            <a href="/register">
              <u>Register here</u>
            </a>
          </p>
        ) : (
          <p>
            Already have an account?{" "}
            <a href="/login">
              <u>Log in here</u>
            </a>
          </p>
        )}
      </div>
    </div>
  );
}

export default LoginRegister;
