# 🔐 Auth CMS Backend (Golang)

**Auth CMS Backend** là một dịch vụ xác thực và phân quyền người dùng được xây dựng bằng **Golang + Echo**, phục vụ cho hệ thống quản lý lớp học thông minh. Hệ thống cung cấp các tính năng quản lý tài khoản, xác thực JWT, phân quyền theo vai trò (role-based access control), và khôi phục mật khẩu.

---

## 🚀 Tính năng chính

### ✅ Public Endpoints:
| Method | Endpoint | Mô tả |
|--------|----------|-------|
| GET | `/auth/health` | Kiểm tra trạng thái API |
| POST | `/auth/register` | Đăng ký người dùng mới |
| POST | `/auth/login` | Đăng nhập, trả về access token & refresh token |
| POST | `/auth/forgot-password` | Gửi yêu cầu khôi phục mật khẩu |
| POST | `/auth/refresh-token` | Làm mới access token từ refresh token |

---

### 🔒 Protected Endpoints (yêu cầu JWT):

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| GET | `/auth/me` | Lấy thông tin người dùng hiện tại |
| POST | `/auth/change-password` | Đổi mật khẩu |
| PATCH | `/auth/update-profile` | Cập nhật thông tin cá nhân |

---

### 🔐 Role-based Access Control (RBAC):

| Role | Endpoint | Mô tả |
|------|----------|-------|
| admin | `/auth/admin-only` | Chỉ admin mới truy cập được |
| lecturer | `/auth/lecturer-only` | Chỉ giảng viên mới truy cập được |
| student | `/auth/student-only` | Chỉ sinh viên mới truy cập được |

---

## 🛠 Cài đặt & Chạy server

### 1. Clone project

```bash
git clone https://github.com/yourusername/auth-cms-backend.git
cd auth-cms-backend
```

### 2. Khởi tạo module

```bash
go mod tidy
```

### 3. Chạy server

```bash
go run main.go
```

API sẽ chạy tại: `http://localhost:8080`

---

## 🔐 Cách hoạt động xác thực

- ✅ Hệ thống sử dụng **JWT** để xác thực
- ✅ Sử dụng middleware `JWTAuthMiddleware` để bảo vệ route
- ✅ Sử dụng middleware `RoleMiddleware("role")` để kiểm tra quyền truy cập

### Ví dụ header yêu cầu:

```http
Authorization: Bearer <access_token>
```

---

## 🧩 Phân loại vai trò (role)

- `"admin"` – quản trị viên hệ thống
- `"lecturer"` – giảng viên
- `"student"` – sinh viên

---

## 📦 Công nghệ sử dụng

- Language: **Go**
- Framework: **Echo**
- Auth: **JWT-based**
- Middleware: Custom auth/role middlewares
- Storage: có thể dùng PostgreSQL, MySQL hoặc MongoDB (tùy bạn triển khai)

---

## 👤 Tác giả

- **Tên**: Vũ Bá Đông  
- 📩 Email: [vubadong071102@gmail.com](mailto:vubadong071102@gmail.com)

---

## 📄 License

MIT License