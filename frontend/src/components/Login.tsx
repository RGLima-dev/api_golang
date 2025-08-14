import React, { useState } from "react";

const Login = () => {
  // Estado para o input
  const [Email, setEmail] = useState("");
  const [Password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      const response = await fetch("http://localhost:5000/login", { // URL da sua API
        method: "POST", // método HTTP
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ Email, Password }), // dados enviados
      });

      if (!response.ok) {
        throw new Error("Erro no login");
      }

      const data = await response.json(); // resposta da API
      console.log("Login bem-sucedido:", data);
      alert(`Bem-vindo, ${data.name}`);
    } catch (error) {
      console.error(error);
      alert("Falha no login");
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <input
        type="text"                 // tipo do input
        value={Email}               // valor do estado
        onChange={(e) => setEmail(e.target.value)} // atualiza o estado ao digitar
      />

      <input
        type="text"
        value={Password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handleLogin}>Login</button>
    </div>
  );
};

export default Login;