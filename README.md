# Real-Time Chat App

This is a real-time chat application built using Go for the backend and React Native for the mobile app. The application allows users to engage in instant messaging with each other, providing a seamless and responsive communication experience.

## Features

- **Real-Time Messaging:** Enjoy instant messaging with real-time updates, ensuring quick and responsive communication.

- **User Authentication:** Users can sign up, log in, and securely authenticate themselves to use the chat app.

- **Private and Group Chats:** Create private chats for one-on-one conversations or group chats to communicate with multiple users simultaneously.

- **Push Notifications:** Receive push notifications for new messages and stay updated even when the app is in the background.

- **Cross-Platform Compatibility:** The React Native mobile app ensures compatibility across both iOS and Android platforms.

- **Simple and Intuitive User Interface:** The user interface is designed to be user-friendly, making it easy for users to navigate and use the app efficiently.

## Installation

### Backend (Go)

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/real-time-chat-app.git
   ```

2. Navigate to the server directory:
   ```bash
   cd real-time-chat-app/server
   ```

3. Install dependencies:
   ```bash
   go get -d ./...
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

The backend server should now be running at `http://localhost:8000`.

### Mobile App (React Native)

1. Navigate to the mobile app directory:
   ```bash
   cd real-time-chat-app/mobile
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the app:
   ```bash
   npx expo run:android
      # or
   npx expo run:ios
   ```

Ensure that your mobile device or emulator is connected and configured appropriately.

## Configuration

1. You need running instance of PostgreSql database - check out the Makefile
2. Update the mobile app .env to use the correct server ip and port

## Technologies Used

- **Backend:**
  - Go
  - Gorilla WebSocket (for real-time communication)
  - PostgreSql 

- **Mobile App:**
  - React Native
  - React Navigation
  - Redux  
  - WebSocket