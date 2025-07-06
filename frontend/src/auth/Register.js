// src/pages/Register.js
import { useState } from "react";

function Register() {
  const [form, setForm] = useState({
    username: "",
    email: "",
    password: ""
  });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.username]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8080/api/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          username: form.username,
          email: form.email,
          password: form.password,
          time_balance: 0 // default balance
        })
      });

      if (res.ok) {
        alert("✅ Berhasil register!");
      } else {
        alert("❌ Gagal register");
      }
    } catch (error) {
      console.error("Register error:", error);
      alert("❌ Error saat register");
    }
  };

  return (
    <div>
      <h2>Register User</h2>
      <form onSubmit={handleSubmit}>
        <input username="username" placeholder="Nama" onChange={handleChange} required />
        <input username="email" placeholder="Email" onChange={handleChange} required />
        <input username="password" type="password" placeholder="Password" onChange={handleChange} required />
        <button type="submit">Daftar</button>
      </form>
    </div>
  );
}

export default Register;
