import { Route, Routes } from 'react-router-dom';
import Login from './pages/Login';
import Signup from './pages/Signup';
import CurrentUserContext from './context/current-user-context';
import { useCallback, useContext, useEffect, useState } from 'react';
import axios from 'axios';
import GuestRoutes from './components/GuestRoutes';
import ProtectedRoutes from './components/ProtectedRoute';
import Home from './pages/Home';

const App = () => {
  const { setCurrentUser } = useContext(CurrentUserContext);
  const [loading, setLoading] = useState(true);

  const loadCurrentUser = useCallback(async () => {
    try {
      const response = await axios.get('/api/auth/me');
      setCurrentUser(response.data.user);
    } catch (err) {
      console.error('Failed to load current user:', err);
      setCurrentUser(null);
    }

    setLoading(false);
  }, [setCurrentUser]);

  useEffect(() => {
    loadCurrentUser();
  }, [loadCurrentUser]);

  if (loading) {
    return null;
  }

  return (
    <div id="app">
      <Routes>
        <Route element={<GuestRoutes />}>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
        </Route>

        <Route element={<ProtectedRoutes />}>
          <Route path="/" element={<Home />} />
        </Route>
      </Routes>
    </div>
  );
};

export default App;
