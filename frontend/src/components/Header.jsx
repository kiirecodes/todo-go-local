import React from 'react';
export default function Header({title, children}){
  return (
    <div className='header'>
      <div>
        <h1 style={{margin:0}}>{title}</h1>
        <div className='small'>Modern React + Gin To-Do</div>
      </div>
      <div>{children}</div>
    </div>
  )
}
