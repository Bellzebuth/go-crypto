import { createContext, useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "js-cookie";

interface AuthContextType {
  user: string | null;
  login: (username: string, password: string) => Promise<void>;
  register: (username: string, password: string) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<string | null>(
    Cookies.get("session_token") || null
  );
  const navigate = useNavigate();

  const login = async (username: string, password: string) => {
    const res = await fetch("http://localhost:8080/login", {
      method: "POST",
      credentials: "include", // mandatory for cookies
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });

    if (res.ok) {
      // cookie is defined on server side
      const sessionToken = Cookies.get("session_token");
      if (sessionToken) {
        setUser(sessionToken);
        navigate("/");
      }
    } else {
      console.error("❌ Connection failed");
    }
  };

  const register = async (username: string, password: string) => {
    const res = await fetch("http://localhost:8080/register", {
      method: "POST",
      credentials: "include",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });

    if (res.ok) {
      await login(username, password); // connect user after register
    } else {
      console.error("Registration failed");
    }
  };

  const logout = () => {
    fetch("http://localhost:8080/logout", {
      method: "GET",
      credentials: "include",
    })
      .then(() => {
        Cookies.remove("session_token");
        setUser(null);
        navigate("/auth");
      })
      .catch(console.error);
  };

  return (
    <AuthContext.Provider value={{ user, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext)!;
