import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css"; // Jika Anda ingin menambahkan styling global

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
