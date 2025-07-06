import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../../css/LoginUser.css";

function LoginUser() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const [error, setError] = useState("");

  const handleLogin = async (e) => {
    e.preventDefault();
    setError("");

    try {
      const response = await fetch("http://localhost:8080/api/login/user", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: username, password }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log("✅ Login berhasil:", data);
        navigate("/dashboard/user");
      } else {
        setError("Username atau Password salah.");
      }
    } catch (err) {
      console.error("Login gagal:", err);
      setError("Terjadi kesalahan saat login.");
    }
  };

  return (
    <div className="login-user-container">
      <h2>Login Sebagai User</h2>
      <form onSubmit={handleLogin} className="login-form">
        <input
          type="text"
          placeholder="Nama Pengguna"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />

        <input
          type="password"
          placeholder="Kata Sandi"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        <button type="submit">Masuk</button>
        {error && <p className="error">{error}</p>}
      </form>
    </div>
  );
}

export default LoginUser;
