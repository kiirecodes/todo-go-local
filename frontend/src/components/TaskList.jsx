import React from 'react';

export default function TaskList({tasks, onToggle, onDelete}){
  if (!tasks.length) return <div className='small'>No tasks</div>;
  return (
    <div style={{marginTop:12}}>
      {tasks.map(t=>(
        <div key={t.id} className={'task' + (t.completed ? ' completed' : '')}>
          <div style={{display:'flex', gap:12, alignItems:'center'}}>
            <input type='checkbox' checked={t.completed} onChange={()=>onToggle(t)} />
            <div>
              <div style={{fontWeight:600}}>{t.title}</div>
              <div className='small'>{t.priority} {t.due_date? ' · due '+t.due_date : ''}</div>
            </div>
          </div>
          <div>
            <button className='btn' onClick={()=>onDelete(t.id)}>Delete</button>
          </div>
        </div>
      ))}
    </div>
  )
}
