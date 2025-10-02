import { useEffect } from 'react'
import './App.css'
import { connect, sendMessage } from './api/index.js'

function App() {

  // Connect to the WS server once
  useEffect(() => {
    connect()
  }, [])

  const sendHello = () => {
    sendMessage('hello');
  }

  return (
    <div className="App">
      <button onClick={sendHello}>Send 'hello'</button>
    </div>
  )
}

export default App
