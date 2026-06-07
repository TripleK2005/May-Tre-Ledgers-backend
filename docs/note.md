# Project Documentation

## 1. Cấu trúc thư mục (Project Structure)

```text
.
├── cmd/
│   └── server/
│       └── main.go             # Điểm khởi đầu của ứng dụng (Entry point)
├── docs/                       # Tài liệu dự án (API contract, schema, notes)
├── internal/
│   ├── core/                   # Các thành phần cốt lõi, dùng chung toàn hệ thống
│   │   ├── config/             # Quản lý cấu hình (env vars)
│   │   ├── database/           # Kết nối và cấu hình cơ sở dữ liệu
│   │   └── response/           # Chuẩn hóa phản hồi API (Success, Error)
│   ├── modules/                # Chứa logic nghiệp vụ theo từng module (ví dụ user/, partner/, ...)
│   └── router/                 # Cấu hình định tuyến (Routing)
├── migrations/                 # Các file script khởi tạo và thay đổi database
├── .env.example                # File mẫu cấu hình biến môi trường
├── go.mod                      # Quản lý dependencies
└── README.md                   # Hướng dẫn tổng quan dự án
```

## 2. Quy tắc đặt tên Migration (Migration Naming Convention)

Dự án sử dụng công cụ migration (như `golang-migrate`) với quy tắc đặt tên như sau:

**Định dạng:** `XXXXXX_description.dir.sql`

Trong đó:
- `XXXXXX`: Là số thứ tự gồm 6 chữ số (ví dụ: `000001`, `000002`). Số này phải tăng dần để đảm bảo thứ tự thực thi.
- `description`: Mô tả ngắn gọn nội dung của bản migration, nối với nhau bằng dấu gạch dưới `_` (ví dụ: `create_users`).
- `dir`: Hướng của migration:
    - `up`: Chứa các câu lệnh SQL để áp dụng thay đổi (tạo bảng, thêm cột...).
    - `down`: Chứa các câu lệnh SQL để hoàn tác thay đổi (xóa bảng, xóa cột...).

**Ví dụ:**
- `000001_create_roles.up.sql`
- `000001_create_roles.down.sql`
