import { useState } from "react";
import "./DashboardUser.css";

function DashboardUser() {
  const [activeTab, setActiveTab] = useState("services");

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
      </nav>

      <main className="dashboard-content">
        {renderContent()}
      </main>
    </div>
  );
}

export default DashboardUser;
