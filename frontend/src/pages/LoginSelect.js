import { Link } from "react-router-dom";
import "../css/LoginSelect.css"; // Import your CSS file for styling

function LoginSelect() {
  return (
    <div className="login-select-container">
      <h1>Selamat Datang di JasainAja! 👋</h1>
      <p>Mau masuk sebagai Apa?</p>
      <div className="role-buttons">
        <Link to="/login/user" className="btn user">Login sebagai User</Link>
        <Link to="/login/provider" className="btn provider">Login sebagai Provider</Link>
      </div>
      <p className="reg">Belum punya akun? <Link to="/register/user">Daftar di sini</Link></p>
    </div>
  );
}

export default LoginSelect;
