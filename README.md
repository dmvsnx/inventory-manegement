# Inventory Management API

Simple Inventory Management REST API built with **Golang**, **Fiber**, **GORM**, and **MySQL** using **Repository Pattern** architecture.

## Features

- ✅ Product Management (CRUD)
- ✅ Stock In
- ✅ Stock Out
- ✅ Automatic Stock Update
- ✅ Low Stock Report
- ✅ Clean Architecture
- ✅ Repository Pattern
- ✅ Transaction Support (Stock Movement)

---

# Tech Stack

- Go
- Fiber v2
- GORM
- MySQL
- UUID/Auto Increment ID
- Repository Pattern

---

# Project Structure

```
.
├── cmd/
│    └── main.go
│
├── internal/
│   │── app/
│   ├── configs/
│   │
│   ├── delivery/
│   │   └── handlers/
│   │   └── routes/
│   │
│   ├── repository/
│   │
│   ├── usecase/
│   │
│   ├── model/
│   │
│   └── dto/
│
├── .env.sample
├── go.sum
├── go.mod
└── README.md
```

---

# Installation

Clone repository

```bash
git clone https://github.com/yourusername/inventory-management.git
```

Go to project

```bash
cd inventory-management
```

Install dependency

```bash
go mod tidy
```

Create `.env`

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=inventory_db
```

Run application

```bash
go run cmd/server/main.go
```

---

# API Base URL

```
http://localhost:8080/api/v1
```

---

# Product API

## Create Product

**POST**

```
/products
```

### Request

```json
{
    "sku": "BRG-002",
    "name": "Indomie Goreng",
    "description": "Mie instan goreng",
    "category": "Makanan",
    "unit": "pcs",
    "price": 3500,
    "minimum_stock": 50
}
```

### Response

```json
{
    "data": {
        "id": 2,
        "sku": "BRG-002",
        "name": "Indomie Goreng",
        "description": "Mie instan goreng",
        "category": "Makanan",
        "unit": "pcs",
        "price": 3500,
        "stock": 0,
        "minimum_stock": 50,
        "created_at": "2026-07-02T15:14:22Z",
        "updated_at": "2026-07-02T15:14:22Z"
    },
    "message": "product created successfully"
}
```

---

## Get All Products

**GET**

```
/products
```

---

## Get Product By ID

**GET**

```
/products/:id
```

Example

```
GET /products/2
```

---

## Update Product

**PUT**

```
/products/:id
```

### Request

```json
{
    "sku": "BRG-001",
    "name": "Indomie Goreng Jumbo",
    "description": "Mie instan goreng ukuran besar",
    "category": "Makanan",
    "unit": "pcs",
    "price": 4500,
    "minimum_stock": 30
}
```

---

## Delete Product

**DELETE**

```
/products/:id
```

---

## Low Stock Report

**GET**

```
/reports/low-stock
```

Response

```json
{
    "data": [
        {
            "id": 2,
            "sku": "BRG-002",
            "name": "Indomie Goreng",
            "stock": 15,
            "minimum_stock": 30
        }
    ]
}
```

---

# Stock API

Stock movement automatically updates product stock.

Supported type

```
IN
OUT
```

---

## Create Stock Movement

**POST**

```
/stock
```

### Request (Stock IN)

```json
{
    "product_id": 2,
    "type": "IN",
    "quantity": 100,
    "notes": "Restok dari supplier"
}
```

### Request (Stock OUT)

```json
{
    "product_id": 2,
    "type": "OUT",
    "quantity": 20,
    "notes": "Penjualan"
}
```

### Response

```json
{
    "data": {
        "id": 1,
        "product_id": 2,
        "product": {
            "id": 2,
            "sku": "BRG-002",
            "name": "Indomie Goreng",
            "description": "Mie instan goreng",
            "category": "Makanan",
            "unit": "pcs",
            "price": 3500,
            "stock": 100,
            "minimum_stock": 50,
            "created_at": "2026-07-02T15:14:22Z",
            "updated_at": "2026-07-02T15:17:45Z"
        },
        "type": "IN",
        "quantity": 100,
        "notes": "Restok dari supplier",
        "created_at": "2026-07-02T15:17:45Z",
        "updated_at": "2026-07-02T15:17:45Z"
    },
    "message": "stock movement created successfully"
}
```

---

## Get All Stock Movements

**GET**

```
/stock
```

---

## Get Stock Movement By ID

**GET**

```
/stock/:id
```

Example

```
GET /stock/1
```

---

## Get Stock History By Product

**GET**

```
/stock/product/:productId
```

Example

```
GET /stock/product/2
```

---

# Business Rules

### Product

- SKU must be unique.
- Initial stock is `0`.
- Price cannot be negative.
- Minimum stock cannot be negative.

---

### Stock Movement

- Type must be:
  - `IN`
  - `OUT`
- Quantity must be greater than `0`.
- Stock IN increases product stock.
- Stock OUT decreases product stock.
- Stock cannot become negative.
- All stock transactions use a database transaction to ensure data consistency.

---

# Example Workflow

### 1. Create Product

```
POST /products
```

↓

Stock = `0`

---

### 2. Restock

```
POST /stock
```

```json
{
    "product_id": 2,
    "type": "IN",
    "quantity": 100
}
```

↓

Stock = `100`

---

### 3. Sell Product

```
POST /stock
```

```json
{
    "product_id": 2,
    "type": "OUT",
    "quantity": 25
}
```

↓

Stock = `75`

---

### 4. Check Product

```
GET /products/2
```

↓

Current Stock = `75`

---

### 5. Low Stock Report

```
GET /reports/low-stock
```

Returns products whose current stock is less than or equal to the configured minimum stock.

---

# HTTP Status Code

| Status | Description |
|---------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

---

# Future Improvements

- Authentication (JWT)
- Role Based Access Control (Admin/User)
- Pagination
- Filtering & Search
- Product Image Upload
- Supplier Management
- Purchase Order
- Sales Order
- Dashboard Analytics
- Export Report (Excel/PDF)
- Swagger Documentation
- Docker Support
- Unit Testing
- Integration Testing

---

# Author

**Dmvsnx**

Backend Developer

Built with ❤️ using Golang & Fiber.