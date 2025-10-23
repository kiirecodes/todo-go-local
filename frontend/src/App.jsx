import React, {useEffect, useState} from 'react';
import Header from './components/Header';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';
import TaskFilter from './components/TaskFilter';
import Footer from './components/Footer';
import * as api from './services/api';

/*
  Main App: uses auto theme (prefers-color-scheme), simple search/filter UI
*/
export default function App(){
  const [tasks, setTasks] = useState([]);
  const [filter, setFilter] = useState('all');
  const [query, setQuery] = useState('');

  useEffect(()=>{ load(); }, []);

  async function load(){
    try{
      const res = await api.fetchTodos(query, filter==='completed'?'true': filter==='pending'?'false':'');
      setTasks(res);
    }catch(e){
      console.error('load error', e);
    }
  }

  async function handleCreate(payload){
    await api.createTodo(payload);
    await load();
  }

  async function handleToggle(task){
    const updated = {...task, completed: !task.completed};
    await api.updateTodo(task.id, updated);
    await load();
  }

  async function handleDelete(id){
    if (!confirm('Delete?')) return;
    await api.deleteTodo(id);
    await load();
  }

  	const filtered = tasks?.filter(t => {
		if (query) return (t.title || '').toLowerCase().includes(query.toLowerCase());
		return true;
	}) || [];


  return (
    <div className='app'>
      <div className='card'>
        <Header title='Unique To-Do (React + Gin)'>
          <div className='controls'>
            						<div className='small'>
							Completed: {tasks?.filter(t => t.completed).length || 0}/{tasks?.length || 0}
						</div>

        </Header>

        <div style={{display:'flex', gap:8, marginTop:12}}>
          <input className='input search' placeholder='Search title...' value={query} onChange={e=>setQuery(e.target.value)} />
          <TaskFilter filter={filter} setFilter={setFilter} />
        </div>

        <TaskForm onCreate={handleCreate} />
        <TaskList tasks={filtered} onToggle={handleToggle} onDelete={handleDelete} />
        <Footer />
      </div>
    </div>
  )
}
