# AI Writer 开发规范文档

本文档为AI写作辅助系统提供统一的开发规范，涵盖前后端目录结构、命名规范、API定义规范以及数据库设计规范。

---

## 1. 项目目录结构规范

### 1.1 前端目录结构

```
frontend/
├── src/
│   ├── api/                    # API 接口层
│   │   ├── index.ts            # axios 实例配置
│   │   ├── auth.ts            # 认证相关接口
│   │   ├── works.ts           # 作品相关接口
│   │   ├── chapters.ts        # 章节相关接口
│   │   ├── optimization.ts    # AI优化相关接口
│   │   └── publish.ts         # 发布相关接口
│   │
│   ├── assets/                 # 静态资源
│   │   ├── images/             # 图片资源
│   │   ├── icons/             # 图标资源
│   │   └── fonts/             # 字体资源
│   │
│   ├── components/            # 公共组件
│   │   ├── common/            # 通用组件（按钮、输入框、模态框等）
│   │   │   ├── Button/
│   │   │   ├── Modal/
│   │   │   └── Loading/
│   │   ├── layout/            # 布局组件
│   │   │   ├── Header/
│   │   │   ├── Sidebar/
│   │   │   └── MainLayout/
│   │   └── works/             # 作品相关组件
│   │       ├── WorkCard/
│   │       └── ChapterTree/
│   │
│   ├── hooks/                 # 自定义 Hooks
│   │   ├── useAuth.ts         # 认证状态管理
│   │   ├── useWorks.ts       # 作品列表管理
│   │   └── useOptimistic.ts  # AI优化相关
│   │
│   ├── pages/                 # 页面组件
│   │   ├── Login/             # 登录页
│   │   ├── Register/          # 注册页
│   │   ├── WorksList/        # 作品列表页
│   │   ├── WorkEditor/        # 作品编辑页
│   │   │   ├── BasicInfo/     # 基本信息子页
│   │   │   ├── VolumeInfo/    # 分卷信息子页
│   │   │   └── ChapterInfo/   # 章节信息子页
│   │   ├── Optimization/      # AI优化页
│   │   └── Publish/           # 一键发布页
│   │
│   ├── store/                 # 状态管理
│   │   ├── index.ts           # store 入口
│   │   ├── userSlice.ts       # 用户状态
│   │   ├── worksSlice.ts     # 作品状态
│   │   └── uiSlice.ts        # UI状态
│   │
│   ├── styles/                # 全局样式
│   │   ├── variables.less     # 样式变量
│   │   ├── mixins.less       # 混合宏
│   │   └── global.less       # 全局样式
│   │
│   ├── types/                 # TypeScript 类型定义
│   │   ├── index.ts          # 全局类型导出
│   │   ├── user.ts           # 用户相关类型
│   │   ├── work.ts           # 作品相关类型
│   │   └── api.ts           # API 响应类型
│   │
│   ├── utils/                 # 工具函数
│   │   ├── storage.ts        # 本地存储封装
│   │   ├── format.ts         # 格式化工具
│   │   └── validation.ts    # 校验工具
│   │
│   ├── router/                # 路由配置
│   │   └── index.ts         # 路由入口
│   │
│   ├── App.tsx               # 根组件
│   └── main.tsx              # 入口文件
│
├── public/                    # 公共静态资源
├── package.json
├── tsconfig.json
├── vite.config.ts
└── .eslintrc.js
```

**目录职责说明：**

| 目录 | 职责 |
|------|------|
| `api/` | 封装所有 API 请求，统一管理接口路径和请求配置 |
| `components/` | 可复用的 UI 组件，按功能模块分类 |
| `pages/` | 页面级组件，对应路由页面 |
| `hooks/` | 封装可复用的状态逻辑和副作用 |
| `store/` | 全局状态管理（Zustand/Redux） |
| `types/` | TypeScript 类型定义和接口 |
| `utils/` | 纯函数工具库 |
| `styles/` | 全局样式和样式变量 |
| `router/` | 路由配置和导航守卫 |

---

### 1.2 后端目录结构

```
backend/
├── cmd/                       # 入口程序
│   └── server/
│       └── main.go            # 服务入口
│
├── internal/                  # 内部包（不可被外部导入）
│   ├── config/                # 配置管理
│   │   └── config.go
│   │
│   ├── middleware/            # 中间件
│   │   ├── auth.go           # JWT 认证中间件
│   │   ├── cors.go           # 跨域中间件
│   │   ├── logger.go          # 日志中间件
│   │   └── recovery.go       # 异常恢复中间件
│   │
│   ├── handler/               # 控制器层（处理请求）
│   │   ├── auth.go
│   │   ├── works.go
│   │   ├── volumes.go
│   │   ├── chapters.go
│   │   ├── optimization.go
│   │   └── publish.go
│   │
│   ├── service/               # 业务逻辑层
│   │   ├── auth_service.go
│   │   ├── works_service.go
│   │   ├── volume_service.go
│   │   ├── chapter_service.go
│   │   ├── optimization_service.go
│   │   └── publish_service.go
│   │
│   ├── repository/            # 数据访问层（与数据库交互）
│   │   ├── user_repo.go
│   │   ├── work_repo.go
│   │   ├── volume_repo.go
│   │   ├── chapter_repo.go
│   │   └── ...
│   │
│   ├── model/                 # 数据模型
│   │   ├── user.go
│   │   ├── work.go
│   │   ├── volume.go
│   │   ├── chapter.go
│   │   └── ...
│   │
│   └── dto/                   # 数据传输对象
│       ├── request/           # 请求参数结构
│       │   ├── auth_req.go
│       │   ├── works_req.go
│       │   └── ...
│       │
│       └── response/          # 响应结构
│           ├── auth_resp.go
│           ├── works_resp.go
│           └── ...
│
├── pkg/                       # 可被外部导入的包
│   ├── utils/                 # 工具函数
│   │   ├── hash.go
│   │   ├── jwt.go
│   │   └── response.go
│   │
│   └── errors/                # 错误定义
│       └── errors.go
│
├── api/                       # API 路由定义
│   └── v1/
│       ├── router.go
│       ├── auth.go
│       ├── works.go
│       └── ...
│
├── migrations/                # 数据库迁移
│   └── 001_init.sql
│
├── configs/                   # 配置文件
│   └── .env.example
│
├── go.mod
├── go.sum
└── main.go                    # 入口文件软链接
```

**分层架构说明：**

```
┌─────────────────────────────────────┐
│           Handler (Controller)      │  处理请求参数、调用Service、返回响应
├─────────────────────────────────────┤
│           Service                   │  业务逻辑处理、事务管理
├─────────────────────────────────────┤
│           Repository                │  数据CRUD操作、缓存处理
├─────────────────────────────────────┤
│           Model                     │  数据结构定义
└─────────────────────────────────────┘
```

---

## 2. 文件命名规范

### 2.1 前端文件命名

| 类型 | 命名规则 | 示例 |
|------|----------|------|
| 组件文件 | `组件名.tsx` / `组件名.ts` | `WorkCard.tsx`, `useWorks.ts` |
| 组件样式 | `组件名.module.css` | `WorkCard.module.css` |
| 组件目录 | `PascalCase` | `WorkCard/`, `ChapterTree/` |
| 类型定义 | `类型名.type.ts` / `类型名.ts` | `user.type.ts`, `api.ts` |
| 工具模块 | `功能名.utils.ts` | `format.utils.ts` |
| Hooks | `use功能名.ts` | `useAuth.ts`, `useWorks.ts` |
| 配置文件 | `kebab-case` | `vite.config.ts`, `.env` |
| 常量文件 | `kebab-case` | `constants.ts`, `api-config.ts` |

**组件文件组织示例：**
```
WorkCard/
├── WorkCard.tsx        # 组件实现
├── WorkCard.module.css # 组件样式
└── index.ts            # 统一导出
```

---

### 2.2 后端文件命名

| 类型 | 命名规则 | 示例 |
|------|----------|------|
| Go 源文件 | `snake_case.go` | `user_service.go`, `work_handler.go` |
| 包名 | `snake_case` | `package handler`, `package service` |
| 测试文件 | `源文件_test.go` | `user_service_test.go` |
| 配置文件 | `snake_case` | `config.go`, `.env.example` |
| 迁移文件 | `版本_描述.sql` | `001_create_users.sql` |

**包名与目录名关系：**
- 目录名：`handler`
- 包名：`handler`（可与目录名一致）

---

## 3. 变量、方法命名规范

### 3.1 前端命名

| 类型 | 命名规则 | 示例 |
|------|----------|------|
| 变量 | `camelCase` | `const userName = 'Tom'` |
| 函数 | `camelCase` | `function getUserInfo()` |
| React 组件 | `PascalCase` | `function WorkCard() {}` |
| Hooks | `camelCase`，以 `use` 开头 | `useAuth()`, `useWorks()` |
| 常量 | `UPPER_SNAKE_CASE` | `const API_BASE_URL = '/api'` |
| TypeScript 接口 | `PascalCase`，以 `I` 开头或描述性名词 | `interface IUser`, `interface WorkInfo` |
| TypeScript 类型 | `PascalCase` | `type UserRole = 'admin' \| 'user'` |
| 枚举 | `PascalCase` | `enum UserRole { Admin, User }` |
| 事件处理 | `handle` 开头 | `handleSubmit()`, `handleClick()` |
| 事件 props | `on` 开头 | `onSubmit`, `onChange` |

**TypeScript 接口命名建议：**
```typescript
// 推荐：描述性名词
interface User {
  id: number;
  name: string;
}

// 推荐：Request/Response 后缀
interface LoginRequest {
  username: string;
  password: string;
}

interface LoginResponse {
  token: string;
  user: User;
}

// 可选：I 前缀（团队统一即可）
interface IUser {
  id: number;
}
```

---

### 3.2 后端命名

| 类型 | 命名规则 | 示例 |
|------|----------|------|
| 导出变量 | `PascalCase` | `var JWTKey = "secret"` |
| 未导出变量 | ` camelCase` | `var jwtKey = "secret"` |
| 导出函数 | `PascalCase` | `func Login()` |
| 未导出函数 | `camelCase` | `func login()` |
| 结构体 | `PascalCase` | `type User struct {}` |
| 结构体方法 | `PascalCase` | `func (u *User) GetID()` |
| 接口 | `PascalCase`，以 `er` 结尾 | `type UserService interface {}` |
| 常量 | `PascalCase` 或 `UPPER_SNAKE_CASE` | `const StatusOK = 200` |
| 错误变量 | `Err` 开头 | `var ErrNotFound = errors.New("not found")` |
| 数据库模型 | `PascalCase`，与表名对应 | `type Work struct {}` |

**导出/未导出规则：**
```go
// 导出（公开）
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
}

// 未导出（私有）
type user struct {
    id       int
    password string
}

// 导出的函数
func GetUserByID(id int) (*User, error) {}

// 未导出的函数
func validatePassword(password string) bool {}
```

**接口命名惯例：**
```go
// Service 层接口
type UserService interface {
    GetByID(id int) (*User, error)
    Create(user *User) error
}

// Repository 层接口
type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
}
```

---

## 4. API 接口定义规范

### 4.1 URL 设计

| 规范 | 说明 |
|------|------|
| 版本管理 | 使用 URL 版本号 `/api/v1/` |
| 资源命名 | 使用复数名词 `/api/v1/works` |
| 路径参数 | 使用 `:param` 格式 `/api/v1/works/:id` |
| 查询参数 | 使用小写加连字符 `?page=1&sort_by=created_at` |
| 嵌套资源 | 体现层级关系 `/api/v1/works/:workId/volumes` |

**URL 示例：**
```
GET    /api/v1/works                    # 获取作品列表
POST   /api/v1/works                    # 创建作品
GET    /api/v1/works/:id                # 获取作品详情
PUT    /api/v1/works/:id                # 更新作品
DELETE /api/v1/works/:id                # 删除作品
GET    /api/v1/works/:workId/volumes    # 获取作品的分卷列表
POST   /api/v1/works/:workId/chapters   # 创建章节
```

---

### 4.2 请求方法

| 方法 | 使用场景 | 示例 |
|------|----------|------|
| GET | 获取资源 | `GET /api/v1/works` 获取作品列表 |
| POST | 创建资源 | `POST /api/v1/works` 创建作品 |
| PUT | 完整更新资源 | `PUT /api/v1/works/:id` 完整更新作品 |
| PATCH | 部分更新资源 | `PATCH /api/v1/works/:id` 更新部分字段 |
| DELETE | 删除资源 | `DELETE /api/v1/works/:id` 删除作品 |

---

### 4.3 请求与响应格式

**请求 Header：**
```
Content-Type: application/json
Authorization: Bearer <token>
```

**成功响应格式：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "我的作品",
    "wordCount": 50000
  }
}
```

**分页响应格式：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [...],
    "pagination": {
      "page": 1,
      "pageSize": 10,
      "total": 100,
      "totalPages": 10
    }
  }
}
```

**错误响应格式：**
```json
{
  "code": 1001,
  "message": "作品不存在",
  "detail": "The work with id 123 does not exist"
}
```

**通用查询参数：**
| 参数 | 说明 | 示例 |
|------|------|------|
| `page` | 页码 | `?page=1` |
| `page_size` | 每页数量 | `?page_size=20` |
| `sort_by` | 排序字段 | `?sort_by=created_at` |
| `order` | 排序方向 | `?order=desc` |
| `keyword` | 搜索关键词 | `?keyword=主角` |

---

### 4.4 状态码

**HTTP 状态码：**
| 状态码 | 说明 |
|--------|------|
| 200 | OK，请求成功 |
| 201 | Created，创建成功 |
| 204 | No Content，删除成功 |
| 400 | Bad Request，请求参数错误 |
| 401 | Unauthorized，未登录或token无效 |
| 403 | Forbidden，无权限 |
| 404 | Not Found，资源不存在 |
| 409 | Conflict，资源冲突（如用户名已存在） |
| 422 | Unprocessable Entity，参数验证失败 |
| 429 | Too Many Requests，请求过于频繁 |
| 500 | Internal Server Error，服务器内部错误 |

**业务错误码：**
| 错误码 | 说明 |
|--------|------|
| 1000 | 通用错误 |
| 1001 | 资源不存在 |
| 1002 | 参数错误 |
| 1003 | 权限不足 |
| 2001 | 用户名或密码错误 |
| 2002 | Token过期 |
| 2003 | Token无效 |
| 3001 | 作品不存在 |
| 3002 | 章节不存在 |
| 4001 | AI服务调用失败 |
| 5001 | 发布失败 |

---

### 4.5 认证方式

**JWT Token 传递：**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Token 包含信息（Payload）：**
```json
{
  "user_id": 1,
  "username": "writer",
  "exp": 1700000000,
  "iat": 1699900000
}
```

---

## 5. 数据库表结构、字段定义规范

### 5.1 表命名

| 规范 | 说明 |
|------|------|
| 命名风格 | 小写加下划线 `snake_case` |
| 复数名词 | 使用复数形式 `users`, `works`, `chapters` |
| 前缀 | 不使用表前缀（如需区分可用业务前缀） |

**表名示例：**
| 表名 | 说明 |
|------|------|
| `users` | 用户表 |
| `works` | 作品表 |
| `work_categories` | 作品类别表 |
| `volumes` | 分卷表 |
| `chapters` | 章节表 |
| `scenes` | 场景表 |
| `optimization_steps` | 优化步骤表 |
| `optimization_records` | 优化记录表 |
| `publish_tasks` | 发布任务表 |
| `notifications` | 通知表 |

---

### 5.2 字段命名

| 字段类型 | 命名规则 | 示例 |
|----------|----------|------|
| 主键 | `id` | `id BIGINT PRIMARY KEY` |
| 外键 | `单数表名_id` | `user_id`, `work_id` |
| 时间戳 | `created_at`, `updated_at` | `created_at TIMESTAMP` |
| 软删除 | `deleted_at` | `deleted_at TIMESTAMP NULL` |
| 布尔值 | `is_xxx` 格式 | `is_default`, `is_public` |
| 计数 | `xxx_count` 格式 | `chapter_count`, `word_count` |

**通用字段：**
```sql
created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
deleted_at  TIMESTAMP NULL
```

---

### 5.3 数据类型选择

| 场景 | 推荐类型 | 说明 |
|------|----------|------|
| 主键自增 | `BIGINT UNSIGNED` | 使用无符号大整数避免溢出 |
| 用户名 | `VARCHAR(50)` | 限制最大长度 |
| 邮箱 | `VARCHAR(255)` | 邮箱地址较长 |
| 密码哈希 | `VARCHAR(255)` | bcrypt 哈希后长度 |
| 手机号 | `VARCHAR(20)` | 考虑国际号码 |
| 简介/描述 | `TEXT` | 长文本 |
| JSON 数据 | `JSON` | MySQL 5.7+ 支持 |
| 时间 | `TIMESTAMP` | 带时区，范围有限 |
| 时间 | `DATETIME` | 不带时区，范围更大 |
| 状态枚举 | `TINYINT` | 0-255 的整数 |

**字段定义示例：**
```sql
CREATE TABLE users (
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    username    VARCHAR(50) NOT NULL COMMENT '用户名',
    password    VARCHAR(255) NOT NULL COMMENT '密码哈希',
    email       VARCHAR(255) NOT NULL COMMENT '邮箱',
    avatar      VARCHAR(500) NULL COMMENT '头像URL',
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uk_username (username),
    UNIQUE KEY uk_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---

### 5.4 索引规范

| 索引类型 | 命名规则 | 示例 |
|----------|----------|------|
| 主键索引 | `PRIMARY` | 主键自动为 PRIMARY |
| 唯一索引 | `uniq_表名_字段` | `uniq_users_email` |
| 普通索引 | `idx_表名_字段` | `idx_works_user_id` |
| 组合索引 | `idx_表名_字段1_字段2` | `idx_chapters_work_volume` |

**索引创建原则：**
- WHERE 条件中频繁使用的字段
- ORDER BY 排序的字段
- JOIN 操作的关联字段
- 避免在区分度低的字段上建索引（如性别、状态）
- 组合索引注意字段顺序（区分度高的放前面）

---

### 5.5 其他规范

**字符集与排序规则：**
```sql
DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

说明：
- `utf8mb4` 支持完整的 UTF-8 字符（包括 emoji）
- `utf8mb4_unicode_ci` 排序规则对中文友好

**引擎选择：**
```sql
ENGINE=InnoDB
```

**外键约束：**
- 推荐使用外键约束保证数据一致性
- 重要业务表建议添加外键
- 注意：频繁更新的表慎用外键

---

## 6. 附录

### 6.1 目录结构快速创建脚本

**前端：**
```bash
mkdir -p src/{api,assets,components/{common,layout,works},hooks,pages/{Login,Register,WorksList,WorkEditor/{BasicInfo,VolumeInfo,ChapterInfo},Optimization,Publish},store,styles,types,utils,router}
```

**后端：**
```bash
mkdir -p cmd/server internal/{config,middleware,handler,service,repository,model,dto/{request,response}} pkg/{utils,errors} api/v1 migrations configs
```

---

### 6.2 常用配置

**前端环境变量：**
```
VITE_API_BASE_URL=/api/v1
VITE_UPLOAD_URL=/api/v1/upload
```

**后端环境变量：**
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=aiwriter
JWT_SECRET=your-secret-key
AI_API_KEY=your-api-key
```

---

*本文档将随项目迭代持续更新*
