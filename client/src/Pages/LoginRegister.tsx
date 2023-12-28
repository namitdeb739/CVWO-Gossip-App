import "./LoginRegister.css";
import user_icon from "./../Assets/Images/User.png";
import key_icon from "./../Assets/Images/Key.png";
import eye_icon from "./../Assets/Images/Eye.png";
import { SyntheticEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

function LoginRegister(props: {
  setUsername?: (name: string) => void;
  type: string;
}) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [attemptErrorMessage, setAttemptErrorMessage] = useState("");
  const [hidePassword, setHidePassword] = useState(true);
  const [registerRedirect, setRegisterRedirect] = useState(false);
  const [loginRedirect, setLoginRedirect] = useState(false);

  const isUsernameValid = /^[^\s]+$/.test(username) || username === "";
  const isPasswordValid =
    /^(?=.*[a-z])(?=.*[A-Z])[^\s]{8,}$/.test(password) || password === "";

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    const response = await fetch(
      "http://localhost:8080/api/" + props.type,
      props.type === "Register"
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
            },
            credentials: "include",
            body: JSON.stringify({
              username,
              password,
            }),
          }
    );

    const content = await response.json();

    if (content.status !== "success") {
      setAttemptErrorMessage(content.message)
    } else {
      setAttemptErrorMessage("");
      if (props.setUsername != undefined) {
        props.setUsername(content.data.Username);
      }

      props.type === "Register"
        ? setRegisterRedirect(true)
        : setLoginRedirect(true);
    }
  };

  const navigate = useNavigate();
  if (registerRedirect) {
    navigate("/login");
  }
  if (loginRedirect) {
    navigate("/");
  }

  return (
    <div className="container">
      <div className="header">
        <div className="text">{props.type}</div>
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
        {attemptErrorMessage===""?<p></p>:<p>{attemptErrorMessage}; try again</p>}
      </div>
      <div className="submit-container">
        <button
          className={
            !isUsernameValid || !isPasswordValid ? "disabled" : "submit"
          }
          onClick={submit}
          disabled={!isUsernameValid || !isPasswordValid}
        >
          {props.type}
        </button>
      </div>
      <div>
        {props.type === "Login" ? (
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
