# IronOps

IronOps is a modern, comprehensive IT operations and monitoring platform designed to streamline infrastructure management. It provides real-time monitoring, alerting, user management, and audit logging in a unified interface.

## 🚀 Features

- **Dashboard**: Real-time overview of system status, service health, and active alerts using WebSockets.
- **Service Management**: Track services and their running instances across different environments (Dev, Test, Prod).
- **Monitoring**: Real-time CPU and Memory usage tracking for registered instances.
- **Alerting System**:
  - **Rules**: valid configurable threshold-based rules (e.g., CPU > 80%).
  - **Channels**: Support for Feishu, DingTalk, and generic Webhooks.
  - **Notifications**: Instant alerts when thresholds are breached.
- **User Management**: Role-Based Access Control (RBAC) with Admin, Ops, and Viewer roles.
- **Audit Logs**: Comprehensive logging of all critical user actions and system events.

## 🛠️ Tech Stack

### Backend
- **Language**: Go (1.23+)
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: MySQL
- **Real-time**: WebSocket
- **System Stats**: [gopsutil](https://github.com/shirou/gopsutil)

### Frontend
- **Framework**: Vue 3
- **Build Tool**: Vite
- **UI Library**: [Element Plus](https://element-plus.org/)
- **State Management**: Pinia
- **Routing**: Vue Router
- **Visualization**: ECharts
- **HTTP Client**: Axios

## 📋 Prerequisites

- **Go**: v1.23 or higher
- **Node.js**: v16 or higher
- **MySQL**: v5.7 or v8.0

## ⚡ Getting Started

### 1. Database Setup

Ensure you have a MySQL instance running and create a database named `ironops`.

```sql
CREATE DATABASE ironops CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 2. Backend Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/IronOps.git
   cd IronOps
   ```

2. Configure the database connection. You can set the `MYSQL_DSN` environment variable or use the default (for local dev).
   
   **Default DSN**: `root:12345678@tcp(10.21.32.13:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local`

   To override:
   ```bash
   # Windows PowerShell
   $env:MYSQL_DSN="user:password@tcp(localhost:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local"
   ```

3. Run the backend server:
   ```bash
   go run cmd/server/main.go
   ```
   The server will start on `http://localhost:8080`.
   *Note: The application will automatically migrate database schemas and seed initial data if needed.*

### 3. Frontend Setup

1. Navigate to the web directory:
   ```bash
   cd web
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```
   The frontend will be available at `http://localhost:5173`.

## 📂 Project Structure

```
IronOps/
├── cmd/
│   ├── server/          # Main application entry point
│   └── seed/            # Database seeding scripts (optional)
├── internal/
│   ├── database/        # Database initialization and config
│   ├── handler/         # HTTP request handlers
│   ├── middleware/      # Auth and logging middleware
│   ├── model/           # Database models
│   └── service/         # Business logic (Alert Engine, Notifier, etc.)
├── web/                 # Vue 3 Frontend application
│   ├── src/
│   │   ├── api/         # API integration
│   │   ├── components/  # Reusable UI components
│   │   ├── views/       # Page views
│   │   └── ...
│   └── ...
└── ...
```

## 🔐 Default Credentials

(For testing/seeding purposes if using the seed script)
- **Admin**: `admin` / `123456`
- **Ops**: `ops_lead` / `123456`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License.
