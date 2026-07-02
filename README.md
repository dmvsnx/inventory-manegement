# Inventory Management API

Simple Inventory Management REST API built with **Go**, **Fiber v2**, **GORM**, and **PostgreSQL** using Clean Architecture.

## Features

- Product Management (CRUD)
- Stock In / Stock Out
- Automatic Stock Update
- Low Stock Report
- Filter Stock by Type & Date Range
- Transaction Support (atomic stock movements)
- Soft Delete

## Tech Stack

- Go 1.26
- Fiber v2
- GORM
- PostgreSQL
- Repository Pattern + Usecase + Handler

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── app/            # Server setup & graceful shutdown
│   ├── config/         # Environment config loader
│   ├── database/       # PostgreSQL connection + auto-migrate
│   ├── delivery/
│   │   ├── handlers/   # HTTP handlers
│   │   └── routes/     # Route registration
│   ├── dto/            # Request/response DTOs
│   ├── model/          # GORM models
│   ├── repository/     # Data access layer
│   ├── usecase/        # Business logic layer
│   └── utils/          # Validator service
├── .env.sample
├── go.mod
├── go.sum
└── README.md
```

## Installation

```bash
git clone https://github.com/dmvsnx/inventory-manegement.git
cd inventory-manegement
go mod tidy
```

### Setup `.env`

```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=inventory_db
```

### Run

```bash
go run cmd/main.go
```

## API Base URL

```
http://localhost:3000/api
```

## Product API

### Create Product

**POST** `/api/products`

```json
{
  "sku": "BRG-001",
  "name": "Indomie Goreng",
  "description": "Mie instan goreng",
  "category": "Makanan",
  "unit": "pcs",
  "price": 3500,
  "minimum_stock": 50
}
```

**Response** `201`

```json
{
  "message": "product created successfully",
  "data": {
    "id": 1,
    "sku": "BRG-001",
    "name": "Indomie Goreng",
    "description": "Mie instan goreng",
    "category": "Makanan",
    "unit": "pcs",
    "price": 3500,
    "stock": 0,
    "minimum_stock": 50,
    "created_at": "2026-07-02T15:14:22Z",
    "updated_at": "2026-07-02T15:14:22Z"
  }
}
```

### Get All Products

**GET** `/api/products`

**Response** `200`

```json
{
  "data": [
    {
      "id": 1,
      "sku": "BRG-001",
      "name": "Indomie Goreng",
      "description": "Mie instan goreng",
      "category": "Makanan",
      "unit": "pcs",
      "price": 3500,
      "stock": 90,
      "minimum_stock": 50,
      "created_at": "2026-07-02T15:14:22Z",
      "updated_at": "2026-07-02T15:17:45Z"
    }
  ]
}
```

### Get Product By ID

**GET** `/api/products/:id`

**Response** `200`

```json
{
  "data": {
    "id": 1,
    "sku": "BRG-001",
    "name": "Indomie Goreng",
    "description": "Mie instan goreng",
    "category": "Makanan",
    "unit": "pcs",
    "price": 3500,
    "stock": 90,
    "minimum_stock": 50,
    "created_at": "2026-07-02T15:14:22Z",
    "updated_at": "2026-07-02T15:17:45Z"
  }
}
```

### Update Product

**PUT** `/api/products/:id`

**Request**

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

**Response** `200`

```json
{
  "message": "product updated successfully",
  "data": {
    "id": 1,
    "sku": "BRG-001",
    "name": "Indomie Goreng Jumbo",
    "description": "Mie instan goreng ukuran besar",
    "category": "Makanan",
    "unit": "pcs",
    "price": 4500,
    "stock": 90,
    "minimum_stock": 30,
    "created_at": "2026-07-02T15:14:22Z",
    "updated_at": "2026-07-02T15:20:00Z"
  }
}
```

### Delete Product

**DELETE** `/api/products/:id`

**Response** `200`

```json
{
  "message": "product deleted successfully"
}
```

### Low Stock Report

**GET** `/api/reports/low-stock`

Returns products where `stock < minimum_stock`.

**Response** `200`

```json
{
  "data": [
    {
      "id": 2,
      "sku": "BRG-002",
      "name": "Aqua Galon",
      "description": "Air minum galon",
      "category": "Minuman",
      "unit": "pcs",
      "price": 20000,
      "stock": 3,
      "minimum_stock": 10,
      "created_at": "2026-07-02T15:14:22Z",
      "updated_at": "2026-07-02T15:17:45Z"
    }
  ]
}
```

## Stock API

Stock movements automatically update product stock in a single database transaction.

### Create Stock Movement

**POST** `/api/stock`

**Request (Stock IN)**

```json
{
  "product_id": 1,
  "type": "IN",
  "quantity": 100,
  "notes": "Restok dari supplier"
}
```

**Request (Stock OUT)**

```json
{
  "product_id": 1,
  "type": "OUT",
  "quantity": 10,
  "notes": "Penjualan offline"
}
```

**Response** `201`

```json
{
  "message": "stock movement created successfully",
  "data": {
    "id": 1,
    "product_id": 1,
    "product": {
      "id": 1,
      "sku": "BRG-001",
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
  }
}
```

### Get All Stock Movements

**GET** `/api/stock`

Optional query params:
- `?type=IN` — filter by IN/OUT
- `?start=2026-01-01T00:00:00Z&end=2026-07-02T23:59:59Z` — filter by date range

**Response** `200`

```json
{
  "data": [
    {
      "id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "sku": "BRG-001",
        "name": "Indomie Goreng",
        "price": 3500,
        "stock": 100,
        "minimum_stock": 50
      },
      "type": "IN",
      "quantity": 100,
      "notes": "Restok dari supplier",
      "created_at": "2026-07-02T15:17:45Z",
      "updated_at": "2026-07-02T15:17:45Z"
    }
  ]
}
```

### Get Stock Movement By ID

**GET** `/api/stock/:id`

**Response** `200`

```json
{
  "data": {
    "id": 1,
    "product_id": 1,
    "product": {
      "id": 1,
      "sku": "BRG-001",
      "name": "Indomie Goreng",
      "price": 3500,
      "stock": 100,
      "minimum_stock": 50
    },
    "type": "IN",
    "quantity": 100,
    "notes": "Restok dari supplier",
    "created_at": "2026-07-02T15:17:45Z",
    "updated_at": "2026-07-02T15:17:45Z"
  }
}
```

### Get Stock History By Product

**GET** `/api/stock/product/:productId`

**Response** `200`

```json
{
  "data": [
    {
      "id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "sku": "BRG-001",
        "name": "Indomie Goreng",
        "price": 3500,
        "stock": 100,
        "minimum_stock": 50
      },
      "type": "IN",
      "quantity": 100,
      "notes": "Restok dari supplier",
      "created_at": "2026-07-02T15:17:45Z",
      "updated_at": "2026-07-02T15:17:45Z"
    }
  ]
}
```

## Business Rules

### Product
- SKU must be unique
- Initial stock is `0`
- Price must be > 0
- `unit` defaults to `"pcs"`

### Stock Movement
- Type must be `IN` or `OUT`
- Quantity must be > 0
- Stock IN increases product stock
- Stock OUT decreases product stock
- Stock cannot go below `0`
- All transactions use database transaction (atomic)

## HTTP Status Codes

| Status | Description  |
|--------|--------------|
| 200    | Success      |
| 201    | Created      |
| 400    | Bad Request  |
| 404    | Not Found    |
| 409    | Conflict     |
| 500    | Server Error |

## Author

**Dmvsnx** — Backend Developer
