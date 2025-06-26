import { createContext } from 'react';

export interface CurrentUser {
  id: number;
  name: string;
  email: string;
}

export interface CurrentUserContextType {
  currentUser: CurrentUser | null;
  setCurrentUser: React.Dispatch<React.SetStateAction<CurrentUser | null>>;
}

const CurrentUserContext = createContext<CurrentUserContextType>({
  currentUser: null,
  setCurrentUser: () => {},
});

export default CurrentUserContext;
