import { useState } from "react";
import "../../css/db_user.css"; 
import { useNavigate } from 'react-router-dom';

function DashboardUser() {
  const [activeTab, setActiveTab] = useState("services");
  const navigate = useNavigate();
  const handleLogout = () => {
    // Hapus data login dari localStorage
    localStorage.removeItem('user');
    // Redirect ke halaman login
    navigate('/login/user');
  };

  const renderContent = () => {
    switch (activeTab) {
      case "services":
        return <p>🛠️ Menampilkan daftar jasa yang tersedia...</p>;
      case "request":
        return <p>📝 Form request jasa akan tampil di sini...</p>;
      case "history":
        return <p>📜 Riwayat request jasa akan tampil di sini...</p>;
      default:
        return null;
    }
  };

  return (
    <div className="dashboard-user">
      <header>
        <h1>👋 Selamat datang, User!</h1>
        <p>Kelola aktivitas jasa kamu di sini</p>
      </header>

      <nav className="dashboard-menu">
        <button
          className={activeTab === "services" ? "active" : ""}
          onClick={() => setActiveTab("services")}
        >
          🔍 Lihat Services
        </button>
        <button
          className={activeTab === "request" ? "active" : ""}
          onClick={() => setActiveTab("request")}
        >
          📝 Request Jasa
        </button>
        <button
          className={activeTab === "history" ? "active" : ""}
          onClick={() => setActiveTab("history")}
        >
          📜 Riwayat
        </button>
        <button
          onClick={handleLogout}
          className="px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg shadow"
        >
          🔓 Logout
        </button>
      </nav>

      <main className="dashboard-content">
        {renderContent()}
      </main>
    </div>
  );
}

export default DashboardUser;
