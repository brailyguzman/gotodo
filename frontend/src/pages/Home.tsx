import { useContext, useEffect, useState } from 'react';
import CurrentUserContext from '../context/current-user-context';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

interface Todo {
  ID: number;
  text: string;
  done: boolean;
}

const Home = () => {
  const { currentUser, setCurrentUser } = useContext(CurrentUserContext);
  const navigate = useNavigate();
  const [newTodo, setNewTodo] = useState('');
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const res = await axios.get('/api/todos/');
        setTodos(res.data.todos);
      } catch (err) {
        console.error('Failed to fetch todos:', err);
      } finally {
        setLoading(false);
      }
    };
    fetchTodos();
  }, []);

  const handleAddTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTodo.trim()) return;
    try {
      const res = await axios.post('/api/todos/', { text: newTodo });
      setTodos([...todos, { ID: res.data.id, text: newTodo, done: false }]);
      setNewTodo('');
    } catch (err) {
      console.error('Failed to add todo:', err);
    }
  };

  const handleToggleDone = async (id: number, done: boolean) => {
    try {
      await axios.patch(`/api/todos/${id}`, {
        done,
        text: todos.find((t) => t.ID === id)?.text,
      });
      setTodos(
        todos.map((todo) => (todo.ID === id ? { ...todo, done } : todo))
      );
    } catch (err) {
      console.error('Failed to toggle todo:', err);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await axios.delete(`/api/todos/${id}`);
      setTodos(todos.filter((todo) => todo.ID !== id));
    } catch (err) {
      console.error('Failed to delete todo:', err);
    }
  };

  const handleLogout = async () => {
    try {
      await axios.post('/api/auth/logout');
      setCurrentUser(null);
      navigate('/login');
    } catch (err) {
      console.error('Failed to logout:', err);
    }
  };

  if (loading) return null;

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 pt-16 rounded shadow-md w-full max-w-2xl h-dvh relative">
        <button
          onClick={handleLogout}
          className="absolute top-6 right-8 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 text-base font-semibold"
        >
          Logout
        </button>
        <h1 className="text-3xl font-bold mb-6 text-center text-blue-700">
          Hey {currentUser?.name.split(' ')[0]}, what do you want to do today?
        </h1>
        <form onSubmit={handleAddTodo} className="flex gap-2 mb-8">
          <input
            value={newTodo}
            onChange={(e) => setNewTodo(e.target.value)}
            className="flex-1 px-4 py-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 text-lg"
            placeholder="Add a new todo"
          />
          <button
            type="submit"
            className="bg-blue-600 text-white px-6 py-3 rounded hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 text-lg font-semibold"
          >
            Add
          </button>
        </form>
        <ul className="space-y-4 overflow-y-auto h-[calc(100vh-300px)]">
          {todos.length === 0 && (
            <li className="text-center text-gray-400 text-xl">
              No todos yet. Add one!
            </li>
          )}
          {todos.map((todo) => (
            <li
              key={todo.ID}
              className="flex items-center justify-between bg-gray-50 rounded px-4 py-4 shadow hover:bg-blue-50 transition group"
            >
              <div className="flex items-center gap-4">
                <input
                  type="checkbox"
                  checked={todo.done}
                  onChange={() => handleToggleDone(todo.ID, !todo.done)}
                  className="w-6 h-6 accent-blue-500"
                />
                <span
                  className={`text-xl ${
                    todo.done ? 'line-through text-gray-400' : 'text-gray-800'
                  }`}
                >
                  {todo.text}
                </span>
              </div>
              <button
                onClick={() => handleDelete(todo.ID)}
                className="text-red-500 text-lg font-bold px-4 py-2 rounded hover:bg-red-100 transition"
                title="Delete"
              >
                ✕
              </button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Home;
