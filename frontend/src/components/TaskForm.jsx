import React, {useState} from 'react';

export default function TaskForm({onCreate}){
  const [title, setTitle] = useState('');
  const [priority, setPriority] = useState('Medium');

  const submit = (e)=>{
    e.preventDefault();
    const t = title.trim();
    if (t.length < 1) return alert('Enter title');
    onCreate({title:t, priority});
    setTitle('');
    setPriority('Medium');
  }

  return (
    <form onSubmit={submit} style={{display:'flex', gap:8, marginTop:12}}>
      <input className='input' value={title} onChange={e=>setTitle(e.target.value)} placeholder='Task title...' />
      <select className='select' value={priority} onChange={e=>setPriority(e.target.value)}>
        <option>Low</option><option>Medium</option><option>High</option>
      </select>
      <button className='btn' type='submit'>Add</button>
    </form>
  )
}
