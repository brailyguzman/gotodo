import { Navigate, Outlet } from 'react-router-dom';
import { useContext } from 'react';
import CurrentUserContext from '../context/current-user-context';

const ProtectedRoutes = () => {
  const { currentUser } = useContext(CurrentUserContext);
  return currentUser ? <Outlet /> : <Navigate to={'/login'} replace />;
};

export default ProtectedRoutes;
