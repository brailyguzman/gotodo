import { useContext } from 'react';
import CurrentUserContext from '../context/current-user-context';

const Home = () => {
  const { currentUser } = useContext(CurrentUserContext);
  return (
    <div>
      <h1>Hello {currentUser?.name}</h1>
      <p>Welcome to the GoTodo application!</p>
    </div>
  );
};

export default Home;
