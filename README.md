# AI Writer - AI写作辅助系统

一个功能完善的AI写作辅助系统，帮助作者高效管理创作流程，提升作品质量。

## 技术栈

### 前端
- React 18 + TypeScript + Vite
- 状态管理: React Context + useReducer / Zustand
- 路由: React Router v6
- UI组件: Ant Design / 自定义组件
- 样式: CSS Modules / Tailwind CSS
- HTTP客户端: axios
- 文本对比: react-diff-viewer

### 后端
- Go 1.20+
- 框架: Gin / Echo
- 对象存储: AWS S3 / MinIO
- 元数据存储: SQLite (本地轻量级)
- 认证: JWT
- AI集成: OpenAI API / 自研模型接口

## 核心功能

### 作品管理
- 作品列表展示（网格/列表视图）
- 作品筛选与排序
- 新增作品向导
- 作品删除

### 作品创作
- 基本信息编辑
- 分卷信息管理
- 章节信息管理
- 场景列表编辑

### AI优化
- 优化步骤列表
- 步骤详情展示
- 文本对比视图
- 自定义步骤
- 优化历史记录

### 一键发布
- 作品与章节选择
- 字数重划分
- 章节名自动生成
- 多平台发布

## 项目结构

```
aiwriter/
├── frontend/          # React前端应用
│   ├── src/
│   └── package.json
├── backend/           # Go后端服务
│   ├── main.go
│   └── go.mod
├── feature_list.json  # 功能测试列表
├── init.sh          # 开发环境初始化脚本
└── README.md
```

## 快速开始

### 环境要求

- Node.js 16+
- Go 1.20+
- AWS S3 兼容存储 (AWS S3 / MinIO / 阿里云OSS)
- pnpm 或 npm

### 安装步骤

1. 克隆项目
```bash
git clone <repository-url>
cd aiwriter
```

2. 使用初始化脚本
```bash
bash init.sh
```

3. 手动初始化（如果需要）

**前端:**
```bash
cd frontend
npm install
npm run dev
```

**后端:**
```bash
cd backend
go mod download
go run main.go
```

### 配置

1. 复制环境配置文件
```bash
cp backend/.env.example backend/.env
```

2. 编辑 `backend/.env`，配置 S3 存储连接和AI服务API密钥

## API端点

### 认证
- POST /api/auth/register - 用户注册
- POST /api/auth/login - 用户登录
- POST /api/auth/logout - 登出
- GET /api/auth/profile - 获取个人信息

### 作品
- GET /api/works - 获取作品列表
- POST /api/works - 创建作品
- GET /api/works/:id - 获取作品详情
- PUT /api/works/:id - 更新作品
- DELETE /api/works/:id - 删除作品

### 分卷
- GET /api/works/:workId/volumes - 获取分卷列表
- POST /api/works/:workId/volumes - 创建分卷

### 章节
- GET /api/works/:workId/chapters - 获取章节列表
- POST /api/works/:workId/chapters - 创建章节

### AI优化
- GET /api/optimization/steps - 获取优化步骤
- POST /api/optimization/execute - 执行优化
- GET /api/optimization/records - 获取优化历史

### 发布
- GET /api/publish/platforms - 获取支持的平台
- POST /api/publish/preview - 预览重划分
- POST /api/publish/execute - 执行发布

## 开发指南

### 代码规范
- 前端使用 ESLint + Prettier
- 后端遵循 Go 代码规范
- 使用 TypeScript 强类型

### 测试
运行测试:
```bash
# 前端测试
cd frontend
npm run test

# 后端测试
cd backend
go test ./...
```

### 构建

```bash
# 前端构建
cd frontend
npm run build

# 后端构建
cd backend
go build -o aiwriter
```

## 功能列表

详见 [feature_list.json](feature_list.json)，包含200+详细测试用例。

## 贡献

欢迎提交 Issue 和 Pull Request。

## 许可证

MIT License
