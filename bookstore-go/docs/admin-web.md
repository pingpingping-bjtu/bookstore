## 🎯 功能特性

### 管理员功能

#### 1. 仪表盘
- 统计卡片展示
- 最近添加的图书
- 系统功能进度
- 快速操作按钮

#### 2. 图书管理
- 图书列表展示（支持搜索、筛选、分页）
- 新增/编辑图书（完整表单）
- 删除图书（确认提示）
- 上架/下架操作
- 图书封面展示

#### 3. 分类管理
- 分类列表展示
- 新增/编辑分类
- 删除分类（安全检查）
- 分类状态管理

#### 4. 订单管理（预留）
- 订单列表查看
- 订单状态管理
- 订单详情查看

#### 5. 用户管理（预留）
- 用户列表查看
- 用户状态管理
- 用户详情查看


## 🔗 API接口

### 管理员API (端口: 8081)

#### 图书管理
- `GET /api/v1/admin/books/list` - 获取图书列表
- `GET /api/v1/admin/books/:id` - 获取图书详情
- `POST /api/v1/admin/books/create` - 创建图书
- `PUT /api/v1/admin/books/:id` - 更新图书
- `DELETE /api/v1/admin/books/:id` - 删除图书
- `PUT /api/v1/admin/books/:id/status` - 更新图书状态

#### 分类管理
- `GET /api/v1/admin/categories/list` - 获取分类列表
- `POST /api/v1/admin/categories/create` - 创建分类
- `PUT /api/v1/admin/categories/:id` - 更新分类
- `DELETE /api/v1/admin/categories/:id` - 删除分类

#### 订单管理
- `GET /api/v1/admin/orders/list` - 获取订单列表
- `GET /api/v1/admin/orders/:id` - 获取订单详情
- `PUT /api/v1/admin/orders/:id/status` - 更新订单状态

#### 用户管理
- `GET /api/v1/admin/users/list` - 获取用户列表
- `GET /api/v1/admin/users/:id` - 获取用户详情
- `PUT /api/v1/admin/users/:id/status` - 更新用户状态

### 用户端API (端口: 8080)

保持原有的用户端API不变，继续为前端用户提供服务。

## 🎨 设计特色

### 圆角设计
- 所有按钮采用圆角设计 (`border-radius: 6px/8px`)
- 输入框、选择器统一圆角风格
- 卡片、模态框采用圆角设计
- 标签和进度条也采用圆角

### 现代化UI
- 渐变按钮效果
- 阴影和过渡动画
- 优雅的加载状态
- 友好的错误提示

### 用户体验
- 直观的操作界面
- 清晰的信息层级
- 友好的错误提示
- 流畅的交互体验

## 🔧 开发指南

### 添加新的管理员功能

1. **添加模型定义** (model/)
2. **添加业务逻辑** (service/)
3. **添加控制器** (web/controller/admin_*.go)
4. **添加路由** (web/router/admin_router.go)
5. **添加前端页面** (bookstore-admin-fronted/src/pages/)