# 📚 博学书城 API 文档

## 📋 概述

本文档详细描述了博学书城系统的所有API接口，包括请求方法、参数、响应格式和示例。

### 基础信息

- **Base URL**: `http://localhost:8080`
- **API Version**: `v1`
- **Content-Type**: `application/json`
- **认证方式**: JWT Token (Bearer Token)

### 响应格式

所有API响应都遵循统一的格式：

```json
{
  "code": 0,           // 状态码：0表示成功，-1表示失败
  "message": "success", // 响应消息
  "data": {}           // 响应数据
}
```

## 🔐 认证相关

### 用户注册

**接口**: `POST /api/v1/user/register`

**请求参数**:
```json
{
  "username": "testuser",
  "password": "123456",
  "confirm_password": "123456",
  "email": "test@example.com",
  "phone": "13800138000",
  "captcha_id": "captcha_id",
  "captcha_value": "1234"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "注册成功"
}
```

### 用户登录

**接口**: `POST /api/v1/user/login`

**请求参数**:
```json
{
  "username": "testuser",
  "password": "123456",
  "captcha_id": "captcha_id",
  "captcha_value": "1234"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "登录成功",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "phone": "13800138000",
      "avatar": "https://example.com/avatar.jpg",
      "is_admin": false
    }
  }
}
```

### 获取用户信息

**接口**: `GET /api/v1/user/profile`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取用户信息成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "phone": "13800138000",
    "avatar": "https://example.com/avatar.jpg",
    "is_admin": false,
    "created_at": "2024-01-01T10:00:00Z"
  }
}
```

## 📖 图书相关

### 获取图书列表

**接口**: `GET /api/v1/book/list`

**查询参数**:
- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 12)

**响应示例**:
```json
{
  "code": 0,
  "message": "获取图书列表成功",
  "data": {
    "books": [
      {
        "id": 1,
        "title": "三体",
        "author": "刘慈欣",
        "price": 5900,
        "discount": 80,
        "type": "科幻",
        "stock": 100,
        "cover_url": "https://example.com/cover.jpg",
        "description": "地球文明与三体文明的星际战争...",
        "isbn": "9787536692930",
        "publisher": "重庆出版社",
        "publish_date": "2008-01-01",
        "pages": 302,
        "language": "中文",
        "format": "平装"
      }
    ],
    "total": 20,
    "total_page": 2,
    "current_page": 1
  }
}
```

### 搜索图书

**接口**: `GET /api/v1/book/search`

**查询参数**:
- `q`: 搜索关键词
- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 12)

**响应示例**:
```json
{
  "code": 0,
  "message": "搜索成功",
  "data": {
    "books": [...],
    "total": 5,
    "total_page": 1,
    "current_page": 1
  }
}
```

### 获取图书详情

**接口**: `GET /api/v1/book/detail/{id}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取图书详情成功",
  "data": {
    "id": 1,
    "title": "三体",
    "author": "刘慈欣",
    "price": 5900,
    "discount": 80,
    "type": "科幻",
    "stock": 100,
    "cover_url": "https://example.com/cover.jpg",
    "description": "地球文明与三体文明的星际战争...",
    "isbn": "9787536692930",
    "publisher": "重庆出版社",
    "publish_date": "2008-01-01",
    "pages": 302,
    "language": "中文",
    "format": "平装"
  }
}
```

### 获取分类图书

**接口**: `GET /api/v1/book/category/{category}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取分类图书成功",
  "data": [
    {
      "id": 1,
      "title": "三体",
      "author": "刘慈欣",
      "price": 5900,
      "discount": 80,
      "type": "科幻",
      "stock": 100,
      "cover_url": "https://example.com/cover.jpg"
    }
  ]
}
```

### 获取热销图书

**接口**: `GET /api/v1/book/hot`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取热销图书成功",
  "data": [
    {
      "id": 1,
      "title": "三体",
      "author": "刘慈欣",
      "price": 5900,
      "discount": 80,
      "type": "科幻",
      "stock": 100,
      "cover_url": "https://example.com/cover.jpg"
    }
  ]
}
```

### 获取新书

**接口**: `GET /api/v1/book/new`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取新书成功",
  "data": [
    {
      "id": 1,
      "title": "三体",
      "author": "刘慈欣",
      "price": 5900,
      "discount": 80,
      "type": "科幻",
      "stock": 100,
      "cover_url": "https://example.com/cover.jpg"
    }
  ]
}
```

## 🛒 订单相关

### 创建订单

**接口**: `POST /api/v1/order/create`

**请求头**: `Authorization: Bearer {token}`

**请求参数**:
```json
{
  "items": [
    {
      "book_id": 1,
      "quantity": 2,
      "price": 4720
    }
  ]
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "创建订单成功",
  "data": {
    "id": 1,
    "user_id": 1,
    "order_no": "ORD202401010001",
    "total_amount": 9440,
    "status": 0,
    "is_paid": false,
    "created_at": "2024-01-01T10:00:00Z",
    "order_items": [
      {
        "id": 1,
        "book_id": 1,
        "quantity": 2,
        "price": 4720,
        "subtotal": 9440,
        "book": {
          "title": "三体",
          "author": "刘慈欣",
          "cover_url": "https://example.com/cover.jpg"
        }
      }
    ]
  }
}
```

### 获取订单列表

**接口**: `GET /api/v1/order/list`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取订单列表成功",
  "data": {
    "orders": [
      {
        "id": 1,
        "order_no": "ORD202401010001",
        "total_amount": 9440,
        "status": 1,
        "is_paid": true,
        "created_at": "2024-01-01T10:00:00Z",
        "payment_time": "2024-01-01T10:05:00Z",
        "order_items": [
          {
            "id": 1,
            "book_id": 1,
            "quantity": 2,
            "price": 4720,
            "subtotal": 9440,
            "book": {
              "title": "三体",
              "author": "刘慈欣",
              "cover_url": "https://example.com/cover.jpg"
            }
          }
        ]
      }
    ],
    "total": 1,
    "page": 1,
    "page_size": 10,
    "total_pages": 1
  }
}
```

### 支付订单

**接口**: `POST /api/v1/order/{id}/pay`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 0,
  "message": "支付成功"
}
```

## ❤️ 收藏相关

### 添加收藏

**接口**: `POST /api/v1/favorite/add`

**请求头**: `Authorization: Bearer {token}`

**请求参数**:
```json
{
  "book_id": 1
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "添加收藏成功"
}
```

### 取消收藏

**接口**: `DELETE /api/v1/favorite/remove`

**请求头**: `Authorization: Bearer {token}`

**请求参数**:
```json
{
  "book_id": 1
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "取消收藏成功"
}
```

### 获取收藏列表

**接口**: `GET /api/v1/favorite/list`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取收藏列表成功",
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "book_id": 1,
      "created_at": "2024-01-01T10:00:00Z",
      "book": {
        "id": 1,
        "title": "三体",
        "author": "刘慈欣",
        "price": 5900,
        "discount": 80,
        "type": "科幻",
        "stock": 100,
        "cover_url": "https://example.com/cover.jpg"
      }
    }
  ]
}
```

### 获取收藏数量

**接口**: `GET /api/v1/favorite/count`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取收藏数量成功",
  "data": {
    "count": 5
  }
}
```

## 🎨 轮播图相关

### 获取轮播图列表

**接口**: `GET /api/v1/carousel/list`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取轮播图成功",
  "data": [
    {
      "id": 1,
      "title": "精选好书推荐",
      "description": "发现更多精彩好书",
      "image_url": "https://example.com/carousel1.jpg",
      "link_url": "/category/文学",
      "sort_order": 1,
      "is_active": true
    }
  ]
}
```

## 📂 分类相关

### 获取分类列表

**接口**: `GET /api/v1/category/list`

**响应示例**:
```json
{
  "code": 0,
  "message": "获取分类列表成功",
  "data": [
    {
      "id": 1,
      "name": "科幻",
      "description": "科幻小说和未来科技",
      "icon": "🚀",
      "color": "#FF6B6B",
      "gradient": "linear-gradient(135deg, #FF6B6B, #FF8E8E)",
      "sort": 1,
      "is_active": true,
      "book_count": 4
    }
  ]
}
```

## 🔐 验证码相关

### 生成验证码

**接口**: `GET /api/v1/captcha/generate`

**响应示例**:
```json
{
  "code": 0,
  "message": "验证码生成成功",
  "data": {
    "captcha_id": "3JU9cHTsexsCQymMZvCb",
    "captcha_base64": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPAAAABQCAMAAAAQlwhOAAAA81BMVEUAAAABfHdm4..."
  }
}
```

## ⚠️ 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| -1 | 失败 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 📝 注意事项

1. **认证**: 需要认证的接口必须在请求头中包含有效的JWT Token
2. **价格**: 所有价格字段都以分为单位存储，前端显示时需要除以100
3. **分页**: 分页接口支持page和page_size参数
4. **时间格式**: 所有时间字段使用ISO 8601格式
5. **图片URL**: 图片URL需要完整的HTTPS地址

## 🔧 开发环境

- **后端服务**: http://localhost:8080
- **前端服务**: http://localhost:3000
- **数据库**: MySQL 8.0+
- **API测试工具**: Postman, curl

---

📖 更多详细信息请参考项目README文档。 