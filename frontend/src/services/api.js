const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

export async function fetchTodos(q='', completed='', priority='') {
  let url = `${API_URL}/todos`;
  const params = [];
  if (q) params.push(`q=${encodeURIComponent(q)}`);
  if (completed) params.push(`completed=${encodeURIComponent(completed)}`);
  if (priority) params.push(`priority=${encodeURIComponent(priority)}`);
  if (params.length) url += `?${params.join('&')}`;
  const res = await fetch(url, {credentials: 'include'});
  if (!res.ok) throw new Error('Network response was not ok');
  return res.json();
}

export async function createTodo(payload) {
  const res = await fetch(`${API_URL}/todos`, {
    method: 'POST',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(payload),
    credentials: 'include'
  });
  if (!res.ok) throw new Error('create failed');
  return res.json();
}

export async function updateTodo(id, payload) {
  const res = await fetch(`${API_URL}/todos/${id}`, {
    method: 'PUT',
    headers: {'Content-Type':'application/json'},
    body: JSON.stringify(payload),
    credentials: 'include'
  });
  if (!res.ok) throw new Error('update failed');
  return res.json();
}

export async function deleteTodo(id) {
  const res = await fetch(`${API_URL}/todos/${id}`, {
    method: 'DELETE',
    credentials: 'include'
  });
  if (!res.ok && res.status !== 204) throw new Error('delete failed');
  return;
}
