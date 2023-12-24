import React from 'react'
import logo from './../assets/Images/Logo.png'
function Header() {
  return (
    <div className='flex'>
      <img src={logo} className='w-[180px]'/>
      <ul className='flex gap-4 md:gap-14'>
        <li>Home</li>
        <li>About</li>
        <li>Browse Subforums</li>
        <li>Profile</li>
      </ul>
      <button>Log In/Out</button>
    </div>
  )
}

export default Header