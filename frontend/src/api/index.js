var socket = new WebSocket('ws://localhost:8080/ws')

// Connects to WS and listens to messages
let connect = () => {
    console.log('Attempting connection...')

    socket.onopen = () => {
        console.log('Successfully connected')
    }

    socket.onmessage = (event) => {
        console.log('Message received: ' + event.data)
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

