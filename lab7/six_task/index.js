const socket = new WebSocket('ws://localhost:8080/ws');
    const messagesDiv = document.getElementById('messages');
    const messageInput = document.getElementById('messageInput');
    const sendButton = document.getElementById('sendButton');

    // Обработка получения сообщения
    socket.onmessage = (event) => {
      const message = document.createElement('div');
      message.textContent = event.data;
      messagesDiv.appendChild(message);
    };

    // Обработка отправки сообщения
    sendButton.addEventListener('click', () => {
      const message = messageInput.value;
      if (message) {
        socket.send(message);
        messageInput.value = '';
      }
    });

    // Обработка открытия соединения
    socket.onopen = () => {
      console.log('Connected to WebSocket server');
    };

    // Обработка закрытия соединения
    socket.onclose = () => {
      console.log('Disconnected from WebSocket server');
    };

    // Обработка ошибок
    socket.onerror = (error) => {
      console.error(`WebSocket error: ${error}`);
    };