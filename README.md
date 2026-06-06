# May-Tre-Ledgers Backend

Repository này chứa backend cho dự án May-Tre-Ledgers (Go).

## Yêu cầu

- Go 1.18+ đã được cài đặt
- PostgreSQL (nếu muốn chạy migration và dùng DB)

## Cài các gói (dependencies)

Từ thư mục gốc của repo, chạy:

```bash
go mod tidy       # thêm/loại bỏ dependency theo code
go mod download   # tải tất cả modules về local cache
```

Nếu muốn cài một package cụ thể (tool hoặc binary):

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Xây dựng và chạy

Xây tất cả package:

```bash
go build ./...
```

Chạy server (từ thư mục gốc):

```bash
cd cmd/server
go run .
```

## Chạy migration (nếu dùng `migrate`)

Ví dụ kết nối PostgreSQL:

```bash
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

Thay `user:pass`, `localhost:5432`, `dbname` bằng thông tin thực tế của bạn.


