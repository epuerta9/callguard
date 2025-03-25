import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
} from "react";

// Local storage hooks
export const getFromLocalStorage = (key: string, initialValue: any) => {
  const storedValue = localStorage.getItem(key);
  return storedValue ? JSON.parse(storedValue) : initialValue;
};

const setToLocalStorage = (key: string, value: any) => {
  localStorage.setItem(key, JSON.stringify(value));
};

const removeFromLocalStorage = (key: string) => {
  localStorage.removeItem(key);
};

interface User {
  user_id: string;
  account_id: string;
  access_token: string;
}

interface StoreContextProps {
  user: User | null;
  loader: boolean;
  setUser: (user: User | null) => void;
  setLoader: (loading: boolean) => void;
}

const initialState: StoreContextProps = {
  user: null,
  loader: false,
  setUser: () => {},
  setLoader: () => {},
};

const StoreContext = createContext<StoreContextProps>(initialState);

export const StoreProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<User | null>(() =>
    getFromLocalStorage("user", null)
  );
  const [loader, setLoader] = useState<boolean>(false);

  useEffect(() => {
    setToLocalStorage("user", user);
  }, [user]);

  return (
    <StoreContext.Provider value={{ user, loader, setUser, setLoader }}>
      {children}
    </StoreContext.Provider>
  );
};

export const useStore = () => {
  const context = useContext(StoreContext);
  if (!context) {
    throw new Error("useStore must be used within a StoreProvider");
  }
  return context;
};

export default StoreContext;
