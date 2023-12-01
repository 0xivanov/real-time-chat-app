import React, { useState, useEffect, useCallback } from 'react';
import { GiftedChat, IMessage } from 'react-native-gifted-chat';

const userId = Math.floor(Math.random() * (1000000 - 1 + 1)) + 1;
const ChatApp: React.FC = () => {
  const [messages, setMessages] = useState<IMessage[]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  useEffect(() => {
    // Initialize WebSocket connection
    const newSocket = new WebSocket('ws://localhost:8000/ws');
    setSocket(newSocket);

    // Connection opened
    newSocket.onopen = () => {
      const helloWebsocket = [{
        _id: Math.floor(Math.random() * (1000000 - 1 + 1)) + 1,
        text: 'Hello, WebSocket!',
        createdAt: new Date(),
        user: {
          _id: userId,
          name: 'React Native',
          avatar: 'https://placeimg.com/140/140/any',
        },
      }]
      newSocket.send(JSON.stringify(helloWebsocket));
    };

    // Message received
    newSocket.onmessage = (event) => {
      const message: IMessage[] = JSON.parse(event.data);
      if (message[0].user._id != userId) {
        setMessages((prevMessages) => GiftedChat.append(prevMessages, message));
      }
    };

    // Error occurred
    newSocket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    // Connection closed
    newSocket.onclose = (event) => {
      console.log('WebSocket connection closed:', event.code, event.reason);
    };

    // Cleanup WebSocket on component unmount
    return () => {
      newSocket.close();
    };
  }, []);

  const onSend = useCallback((messages: IMessage[]) => {
    console.log(userId)
    if (!socket) return;
    console.log(messages)
    // Send the new message to the WebSocket server
    socket.send(JSON.stringify(messages));
    setMessages(previousMessages =>
      GiftedChat.append(previousMessages, messages)
    );
  }, [socket]);

  useEffect(() => {
    setMessages([
      {
        _id: 1,
        text: 'Hello developer',
        createdAt: new Date(),
        user: {
          _id: 2,
          name: 'React Native',
          avatar: 'https://placeimg.com/140/140/any',
        },
      },
    ]);
  }, []);

  return (
    <GiftedChat
      messages={messages}
      onSend={(newMessages) => onSend(newMessages)}
      user={{
        _id: userId
      }}
    />
  );
};

export default ChatApp;
