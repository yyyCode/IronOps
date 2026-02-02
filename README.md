# IronOps

IronOps 是一个现代化、综合性的 IT 运维与监控平台，旨在简化基础设施管理。它在一个统一的界面中提供实时监控、告警、用户管理和审计日志功能。

## 🚀 功能特性

- **仪表盘**：利用 WebSocket 技术，实时概览系统状态、服务健康状况和活动告警。
- **服务管理**：跨不同环境（Dev、Test、Prod）追踪服务及其运行实例。
- **监控**：实时追踪已注册实例的 CPU 和内存使用率。
- **告警系统**：
  - **规则**：可配置的基于阈值的规则（例如：CPU > 80%）。
  - **渠道**：支持飞书、钉钉和通用 Webhook。
  - **通知**：当阈值被触发时即时发送告警通知。
- **用户管理**：基于角色的访问控制（RBAC），包含管理员（Admin）、运维（Ops）和观察者（Viewer）角色。
- **审计日志**：全面记录所有关键的用户操作和系统事件。

## 🛠️ 技术栈

### 后端 (Backend)
- **语言**：Go (1.23+)
- **框架**：[Gin](https://github.com/gin-gonic/gin)
- **ORM**：[GORM](https://gorm.io/)
- **数据库**：MySQL
- **实时通信**：WebSocket
- **系统统计**：[gopsutil](https://github.com/shirou/gopsutil)

### 前端 (Frontend)
- **框架**：Vue 3
- **构建工具**：Vite
- **UI 库**：[Element Plus](https://element-plus.org/)
- **状态管理**：Pinia
- **路由**：Vue Router
- **可视化**：ECharts
- **HTTP 客户端**：Axios

## 📋 前置要求

- **Go**: v1.23 或更高版本
- **Node.js**: v16 或更高版本
- **MySQL**: v5.7 或 v8.0

## ⚡ 快速开始

### 1. 数据库设置

确保你有一个正在运行的 MySQL 实例，并创建一个名为 `ironops` 的数据库。

```sql
CREATE DATABASE ironops CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 2. 后端设置

1. 克隆仓库：
   ```bash
   git clone https://github.com/yourusername/IronOps.git
   cd IronOps
   ```

2. 配置数据库连接。你可以设置 `MYSQL_DSN` 环境变量，或使用默认值（用于本地开发）。
   
   **默认 DSN**: `root:12345678@tcp(10.21.32.13:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local`

   如果需要覆盖默认配置：
   ```bash
   # Windows PowerShell
   $env:MYSQL_DSN="user:password@tcp(localhost:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local"
   
   # Linux/Mac
   export MYSQL_DSN="user:password@tcp(localhost:3306)/ironops?charset=utf8mb4&parseTime=True&loc=Local"
   ```

3. 运行后端服务器：
   ```bash
   go run cmd/server/main.go
   ```
   服务器将在 `http://localhost:8080` 启动。
   *注意：应用程序会自动迁移数据库架构，并在需要时播种初始数据。*

### 3. 前端设置

1. 进入 web 目录：
   ```bash
   cd web
   ```

2. 安装依赖：
   ```bash
   npm install
   ```

3. 启动开发服务器：
   ```bash
   npm run dev
   ```
   前端访问地址为 `http://localhost:5173`。

## 📂 项目结构

```
IronOps/
├── cmd/
│   ├── server/          # 应用程序主入口
│   └── seed/            # 数据库播种脚本（可选）
├── internal/
│   ├── database/        # 数据库初始化和配置
│   ├── handler/         # HTTP 请求处理器
│   ├── middleware/      # 认证和日志中间件
│   ├── model/           # 数据库模型
│   └── service/         # 业务逻辑（告警引擎、通知器等）
├── web/                 # Vue 3 前端应用
│   ├── src/
│   │   ├── api/         # API 集成
│   │   ├── components/  # 可复用 UI 组件
│   │   ├── views/       # 页面视图
│   │   └── ...
│   └── ...
└── ...
```

## 🔐 默认凭据

(如果使用播种脚本进行测试/演示)
- **管理员 (Admin)**: `admin` / `123456`
- **运维 (Ops)**: `ops_lead` / `123456`

## 🤝 贡献

欢迎贡献代码！请随时提交 Pull Request。

## 📄 许可证

本项目采用 MIT 许可证。
