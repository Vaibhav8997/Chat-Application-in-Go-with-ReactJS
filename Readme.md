# Chat Application in Go with ReactJS

This is a **real-time, full-stack chat application** built using **Go** for the backend and **ReactJS** for the frontend.  
It uses **WebSockets** to enable bi-directional communication, providing a dynamic and responsive user experience.

---

## Features
- **Real-time Messaging** – Instantly send and receive messages with all connected users.  
- **Go Backend** – High-performance and concurrent backend powered by Go, handling WebSocket connections and broadcasting messages.  
- **ReactJS Frontend** – Modern single-page app with a clean and intuitive UI.  
- **Scalable Architecture** – Efficient concurrency model of Go handles multiple clients simultaneously.  

---

## Technologies Used
### Backend
- **Go** – Core programming language  
- **WebSockets** – Real-time, two-way communication  

### Frontend
- **ReactJS** – UI development  
- **Sass** – Component styling  

---

## Getting Started

Follow these steps to set up and run the project locally.

### Prerequisites
- [Go](https://go.dev/dl/) (version **1.16+**)  
- [Node.js](https://nodejs.org/) (version **14+**)  
- npm (comes with Node.js)  

---

## Clone the Repository
```bash
git clone https://github.com/Vaibhav8997/Chat-Application-in-Go-with-ReactJS
cd Chat-Application-in-Go-with-ReactJS
```

---

### Backend Setup

Navigate to the backend directory, initialize the Go module, and download dependencies:
```
cd Backend
go mod init
go mod tidy
```

---

### Frontend Setup

Navigate to the frontend directory and install dependencies:
```
cd ../frontend
npm install
```

---

## Run the Application

Start the backend and frontend in separate terminals.

▶️ Start Backend
```
cd Backend
go run main.go
```

▶️ Start Frontend
```
cd frontend
npm start
```

The application will now be running :>
Access the frontend at: http://localhost:3000

---

## 📂 Project Structure
```text
Chat-Application-in-Go-with-ReactJS/
│── Backend/ # Go server (WebSocket handling, broadcasting)
│── frontend/ # ReactJS frontend
│── README.md # Project documentation
``` 