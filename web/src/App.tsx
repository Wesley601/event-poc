import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'

interface User {
  ID: number;
  Name: string;
  Status: string;
}

function App() {
  const [users, setUsers] = useState<User[]>([])

  useEffect(() => {
    const source = new EventSource('http://localhost:3001/progress', { withCredentials: false })
    source.addEventListener('userCreated', function(e) {
      console.log(e.data);
    }, false);

    source.addEventListener('open', function(e) {
      // successful connection.
    }, false);

    source.addEventListener('error', function(e) {
      // error occurred
    }, false);

    fetch('http://localhost:3000/users')
      .then(response => response.json())
      .then(setUsers)

    return () => {
      source.close();
    }
  }, [])

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <ul>
          {users.map(user => (<li key={user.ID}>{user.Name}</li>))}
        </ul>
      </div>
    </div>
  )
}

export default App
