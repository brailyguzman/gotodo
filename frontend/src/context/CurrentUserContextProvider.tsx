import { useState } from 'react';
import CurrentUserContext, { type CurrentUser } from './current-user-context';

const CurrentUserContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [currentUser, setCurrentUser] = useState<CurrentUser | null>(null);

  const context = {
    currentUser,
    setCurrentUser,
  };

  return (
    <CurrentUserContext.Provider value={context}>
      {children}
    </CurrentUserContext.Provider>
  );
};

export default CurrentUserContextProvider;
