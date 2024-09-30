# Golang simple chat application

This project is a simple chat application built with Golang, PostgreSQL, and Websocket

## Services

### 1. Server

The Server is a simple chat application that allows you to chat with your friends

**Key Features:**
- Register, login, logout user
- Authentication with JWT 
- Using Websocket and Go channels to create room, send and receive message (just text for now and do not store message in database)

**Tech Stack:**
- Go
- Gin
- Gorm
- PostgreSQL
- Websocket
- JWT

### 2. Client

The Client is a simple Next.js application

**Tech Stack:**
- Next.js
- Websocket
- TailwindCSS

## How to run

1. Run server
in `server` directory
```bash
go run cmd/server/main.go
```

2. Run client
in `client` directory
```bash
npm run dev
```

3. Open your browser and navigate to `http://localhost:3000`

## How to use

1. Register a new account
2. Login to your account
3. Create a new room or join an existing room
4. Start chatting with another user in the room

## Improvement ideas
Although I have many ideas in mind, due to the urgency of the project at the company and limited time, I could only create a very simple chat app.

Here are some ideas to further develop this app

- Add message storage in the database (Key-value database like Redis or SQL database)
- Implement file sharing capabilities 
- Add support for multimedia messages (images, videos)
- Implement push notifications for new messages (notifications service like Firebase)
- Improve scalability to handle more users and messages
- Ensure reliability with fault-tolerant mechanisms 
- Enhance security with better encryption and data protection 

## Acknowledgements
I would like to express my gratitude to Mr. Hau for organizing and teaching the Golang course. I have learned a lot of basic concepts from Golang and applied some of them to this project. I hope he succeeds in his future courses. TYSM!