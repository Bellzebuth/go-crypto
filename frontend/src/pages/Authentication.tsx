import { useState } from "react";
import Login from "../components/Login";
import Register from "../components/Register";

const Authentication = () => {
  const [isLogin, setIsLogin] = useState(true);

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-2xl font-bold mb-4">
        {isLogin ? "Sign in" : "Sign up"}
      </h1>
      {isLogin ? <Login /> : <Register />}
      <button
        className="mt-4 text-blue-500 underline"
        onClick={() => setIsLogin(!isLogin)}
      >
        {isLogin ? "Sign up" : "Sign in"}
      </button>
    </div>
  );
};

export default Authentication;
