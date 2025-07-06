// src/App.js
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Register from "./auth/Register";
import Login from "./auth/Login";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </Router>
  );
}

export default App;
