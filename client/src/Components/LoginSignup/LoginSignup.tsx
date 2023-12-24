// import React, { useState } from 'react'
import './LoginSignup.css'
import user_icon from './../../Assets/Images/User.png'
import key_icon from './../../Assets/Images/Key.png'
import eye_icon from './../../Assets/Images/Eye.png'
import { useState } from 'react';


export const LoginSignup = () => {
    const [password, setPassword] = useState("");
    const [hidePassword, setHidePassword] = useState(true);

    const handlePassword = (passwordValue: string) => {
        let validLength = false;

        validLength = passwordValue.length >= 8;

        setPassword(passwordValue);
     }
    
  return (
      <div className='container'>
          <div className='header'>
              <div className='text'>Sign Up / Log In</div>
              <div className='underline'></div>
          </div>
          <div className='inputs'>
              <div className='input'>
                  <img src={user_icon} className='inputIcon'/>
                  <input type="Username" placeholder='Username'/>
              </div>
              <div className={password.length >= 8 ? 'input' : 'input invalid'} >
                  <img src={key_icon} className='inputIcon' />
                  <input
                      className={password.length >= 8 ? '' : 'invalid'}
                      value={password}
                      type={hidePassword ? "password" : "text"}
                      onChange={({ target }) => { 
                          handlePassword(target.value);
                      }}
                      placeholder='Password'
                  />
                  <a
                      href='#'
                      className='toggle-btn'
                      onClick={() => {
                          setHidePassword(!hidePassword);
                      }}
                  >
                      <img src={eye_icon} className='inputIcon' />
                  </a>
              </div>
          </div>
          <div className='errorMessage'>
              {password.length >= 8 ? <p></p> : <p>Invalid Password: Length must be at least 8 characters</p>}
          </div>
          <div className='submit-container'>
              <div className='submit'>
                  Sign Up
              </div>
              <div className='submit'>
                  Log In
              </div>
          </div>
    </div>
  )
}
