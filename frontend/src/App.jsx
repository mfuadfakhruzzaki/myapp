import { useState } from "react";
import LoginPage from "./LoginPage";
import RegisterPage from "./RegisterPage";
import Dashboard from "./Dashboard";

export default function App() {
  // Ambil token dari localStorage jika ada
  const [token, setToken] = useState(localStorage.getItem("token") || "");

  // Fungsi logout sederhana untuk menghapus token
  const handleLogout = () => {
    localStorage.removeItem("token");
    setToken("");
  };

  // Jika belum login, tampilkan halaman login dan register
  if (!token) {
    return (
      <div>
        <h1>Welcome to MyApp</h1>
        <LoginPage setToken={setToken} />
        <RegisterPage />
      </div>
    );
  }

  // Jika sudah login, tampilkan dashboard
  return (
    <div>
      <button onClick={handleLogout}>Logout</button>
      <Dashboard token={token} />
    </div>
  );
}
