import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../../css/RegisterUser.css";

function RegisterUser() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData({...formData, [e.target.username]: e.target.value});
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage("");

    try {
      const response = await fetch("http://localhost:8080/api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({...formData, time_balance: 0}),
      });

      if (response.ok) {
        setMessage("✅ Berhasil daftar!");
        setTimeout(() => navigate("/login/user"), 1500);
      } else {
        setMessage("❌ Gagal daftar. Periksa kembali.");
      }
    } catch (err) {
      console.error("Register error:", err);
      setMessage("❌ Terjadi kesalahan.");
    }
  };

  return (
    <div className="register-container">
      <h2>Daftar Akun User</h2>
      <form onSubmit={handleSubmit} className="register-form">
        <input type="text" name="username" placeholder="Nama Pengguna" required onChange={handleChange} />
        <input type="email" name="email" placeholder="Email" required onChange={handleChange} />
        <input type="password" name="password" placeholder="Password" required onChange={handleChange} />
        <button type="submit">Daftar</button>
        {message && <p className="msg">{message}</p>}
      </form>
    </div>
  );
}

export default RegisterUser;
