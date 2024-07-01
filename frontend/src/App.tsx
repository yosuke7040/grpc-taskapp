import { FC, useState } from 'react';
import { MantineProvider, Text, Title } from '@mantine/core';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import { Client } from './Client';
import { HomeRoute } from "./routes/HomeRoute";
import { LoginRoute } from "./routes/LoginRoute";
import { LoginResponse } from "./gen/rpc/auth/v1/auth_pb";

function App(): FC {
  // const [count, setCount] = useState(0)

  const [user, setUser] = useState<LoginResponse | null>(null);
  const handleSubmit = (res: LoginResponse) => {
    setUser(res);
  }

  return (
    <Client baseUrl={"http://localhost:8080"} token={user?.token}>
      <MantineProvider withGlobalStyles withNormalizeCSS>
        <header>
          <Title order={1}>connect-list</Title>
          {user != null && <Text>Logged in</Text>}
        </header>
        <main>
          {user != null ? (
            <HomeRoute />
          ) : (
            <LoginRoute onSubmit={handleSubmit} />
          )}
        </main>

        <footer>
          <Text mt="lg">Footer Links</Text>
        </footer>
      </MantineProvider>
    </Client>
  )
}

export default App
