import { useState } from "react";

// eslint-disable-next-line react/prop-types
export default function LoginPage({ setToken }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8081/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      if (res.ok) {
        const data = await res.json();
        if (data.token) {
          setToken(data.token);
          localStorage.setItem("token", data.token);
        } else {
          alert("Token tidak ditemukan di respons");
        }
      } else {
        const errorText = await res.text();
        alert("Login gagal: " + errorText);
      }
    } catch (err) {
      console.error("Error saat login:", err);
      alert("Terjadi error saat login.");
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            Username:
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </label>
        </div>
        <div>
          <label>
            Password:
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </label>
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
