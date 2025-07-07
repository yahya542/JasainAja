// src/App.js
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginSelect from "./pages/LoginSelect";
import LoginUser from "./pages/user/LoginUser";
import LoginProvider from "./pages/provider/LoginProvider";
import RegisterUser from "./pages/user/RegisterUser";
import RegisterProvider from "./pages/provider/RegisterProvider";
import DashboardUser from "./pages/user/DashboardUser";
import DashboardProvider from "./pages/provider/DashboardProvider"; 

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginSelect />} />
        <Route path="/login/user" element={<LoginUser />} />
        <Route path="/login/provider" element={<LoginProvider />} />
        <Route path="/register/user" element={<RegisterUser />} />
        <Route path="/register/provider" element={<RegisterProvider />} />
        <Route path="/dashboard/user" element={<DashboardUser />} />
        <Route path="/dashboard/provider" element={<DashboardProvider />} /> 
      </Routes>
    </Router>
  );
}

export default App;
