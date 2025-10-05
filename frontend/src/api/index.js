var socket = new WebSocket('ws://localhost:8080/ws')

// Connects to WS and listens to messages
let connect = (callback) => {
    console.log('Attempting connection...')

    socket.onopen = () => {
        console.log('Successfully connected')
    }

    socket.onmessage = (event) => {
        console.log('Message received: ' + event.data)
        callback(event.data) // Passes the message to the callback in the frontend which renders it
    }

    socket.onclose = (event) => {
        console.log('Socket closed connection: ', event)
    }

    socket.onerror = (error) => {
        console.log('Socket error: ', error)
    }
}

let sendMessage = (message) => {
    console.log('Sending message: ' + message)
    socket.send(message)
}

export { connect, sendMessage }

