import React from 'react';
export default function TaskFilter({filter, setFilter}){
  return (
    <div style={{display:'flex', gap:8}}>
      <button className='btn' onClick={()=>setFilter('all')}>All</button>
      <button className='btn' onClick={()=>setFilter('pending')}>Pending</button>
      <button className='btn' onClick={()=>setFilter('completed')}>Completed</button>
    </div>
  )
}
