import { useEffect, useState } from 'react'
import './App.css'
import { connect, sendMessage } from './api/index.js'

function App() {

  const [messages, setMessages] = useState([])
  const [inputMsg, setInputMsg] = useState('')

  // Connect to the WS server once
  useEffect(() => {
    connect(addMsg)
  }, [])

  const addMsg = (msg) => {
    setMessages(prevMsgs => prevMsgs.concat(msg))
  }

  const handleSend = (event) => {
    event.preventDefault()
    if (inputMsg.trim() === '') return // empty input
    sendMessage(inputMsg)
    setInputMsg('')
  }

  return (
    <div className="App">
      <div className="chat-container">
        <div className="chat-box">
          {messages.map((msg, i) => (
            <div key={i} className="chat-message">
              {msg}
            </div>
          ))}
        </div>

        <form className="chat-input-area" onSubmit={handleSend}>
          <input
            type="text"
            className="chat-input"
            placeholder="Type a message..."
            value={inputMsg}
            onChange={(e) => setInputMsg(e.target.value)}
          />
          <button className="send-btn" type="submit">
            Send
          </button>
        </form>
      </div>
    </div>
  )
}

export default App
