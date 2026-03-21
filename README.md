# CS362 Final Project Design Assignment
## E-commerce Order System — Core Feature: Flash Sale with Rate Limiting

---

## Project Structure

```
CS362_E-commerce/
├── cmd/
│   └── server/
│       └── main.go          # Server entry point
├── handler/                 # HTTP request handlers (Controller layer)
├── model/                   # Data structures (Entity layer)
├── service/                 # Business logic layer
├── repository/              # Database access layer
├── go.mod
├── classdiagram.mmd
├── usecase.mmd
└── README.md
```

---

## Git Branch Strategy

```
main
 └── uat
      └── sit
           └── develop
                ├── feat-A/handler
                ├── feat-A/service
                ├── feat-A/repository
                ├── feat-B/handler
                ├── feat-B/service
                └── feat-B/repository
```

| Branch | หน้าที่ |
|--------|---------|
| `main` | Production-ready code |
| `uat` | User Acceptance Testing |
| `sit` | System Integration Testing |
| `develop` | Integration branch สำหรับ merge feature ทั้งหมด |
| `feat-A/*` | Feature A — PlaceOrder Flow (Order + Inventory) |
| `feat-B/*` | Feature B — Payment & Notification Flow |

---

## Step 1: Data Design — Analysis Class Diagram

**File:** `classdiagram.mmd`

| Namespace | Entities |
|-----------|----------|
| CoreDomain | `Customer`, `Address`, `Product`, `Category` |
| InventoryManagement | `Inventory`, `InventoryLock` |
| ShoppingCart | `Cart`, `CartItem` |
| OrderManagement | `Order`, `OrderItem`, `OrderStatus` |
| PaymentProcessing | `Payment`, `PaymentMethod`, `PaymentStatus` |
| NotificationSystem | `Notification`, `NotificationType` |
| FlashSale | `RateLimiter` |

**Key Relationships:**
- `Customer` 1 *-- * `Address` (Composition)
- `Cart` 1 *-- * `CartItem` (Composition)
- `Order` 1 *-- * `OrderItem` (Composition)
- `Customer` 1 o-- * `Order` (Aggregation)
- `Order` 1 --> 0..* `Payment` (1-to-many)
- `Inventory` 1 --> * `InventoryLock` (protected by locks during flash sale)
- `Order` ..> `RateLimiter` (validates flash sale request)

---

## Step 2: Architectural Mapping — Layered Architecture

| Layer | Folder | หน้าที่ |
|-------|--------|---------|
| Handler | `handler/` | รับ HTTP request, validate input, return response |
| Service | `service/` | Business logic, orchestration |
| Repository | `repository/` | Database access, queries |
| Model | `model/` | Data structures, domain objects |
| Entry Point | `cmd/server/` | Start HTTP server |

---

## Step 3: Interface & Controller Contract

### Task 1 — API Specification

#### Order Endpoints
| Method | Path | Description | Status |
|--------|------|-------------|--------|
| POST | `/api/v1/orders` | Place Order (Flash Sale) | 201 / 429 / 409 |
| GET | `/api/v1/orders/{orderId}` | Get Order by ID | 200 / 404 |
| GET | `/api/v1/orders?page=1&pageSize=10` | List Orders | 200 |
| PATCH | `/api/v1/orders/{orderId}/cancel` | Cancel Order | 200 / 409 |

#### Product Endpoints
| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/products` | List Products |
| GET | `/api/v1/products/{productId}` | Get Product by ID |
| GET | `/api/v1/products/search?q=keyword` | Search Products |

#### Cart Endpoints
| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/cart` | Get Cart |
| POST | `/api/v1/cart/items` | Add Item to Cart |
| PATCH | `/api/v1/cart/items/{productId}` | Update Item Quantity |
| DELETE | `/api/v1/cart/items/{productId}` | Remove Item from Cart |

#### Payment Endpoints
| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v1/payments` | Initiate Payment |
| POST | `/api/v1/payments/{paymentId}/refund` | Refund Payment |

#### Sample Request / Response

**POST /api/v1/orders — Request:**
```json
{
  "cartId": "cart_abc",
  "paymentMethod": "CREDIT_CARD",
  "shippingAddressId": "addr_xyz"
}
```

**201 Created:**
```json
{ "success": true, "data": { "orderId": "order_001", "status": "CONFIRMED", "totalAmount": 2499.00 } }
```

**429 Too Many Requests:**
```json
{ "success": false, "error": { "code": "RATE_LIMIT_EXCEEDED", "retryAfter": 60 } }
```

**409 Conflict:**
```json
{ "success": false, "error": { "code": "INSUFFICIENT_STOCK", "productId": "prod_xyz" } }
```

---

### Task 2 — Go Interface Files

#### Service Interfaces

```go
// service/order_service.go
type OrderService interface {
    PlaceOrder(ctx, customerID, req) (*OrderResponse, error)
    GetOrderByID(ctx, customerID, orderID) (*OrderResponse, error)
    GetOrderHistory(ctx, customerID, page, pageSize) (*OrderListResponse, error)
    CancelOrder(ctx, customerID, orderID) (*OrderResponse, error)
}

// service/rate_limiter_service.go
type RateLimiterService interface {
    IsLimitExceeded(ctx, customerID) (bool, error)
    ValidateOrderRequest(ctx, customerID) error
    RecordRequest(ctx, customerID) error
    ResetCounter(ctx, customerID) error
}

// service/inventory_service.go
type InventoryService interface {
    ValidateAndReserve(ctx, productID, qty) error
    AcquireLock(ctx, orderID, productID, customerID, qty) (*InventoryLock, error)
    ReleaseLock(ctx, lockID) error
    FulfillOrder(ctx, orderID) error
}
```

#### Repository Interfaces

```go
// repository/order_repository.go
type OrderRepository interface {
    Create(ctx, order) (*Order, error)
    FindByID(ctx, orderID) (*Order, error)
    FindByCustomerID(ctx, customerID, offset, limit) ([]*Order, int, error)
    UpdateStatus(ctx, orderID, status) error
}

// repository/inventory_repository.go
type InventoryRepository interface {
    FindByProductID(ctx, productID) (*Inventory, error)
    ReserveStock(ctx, productID, qty) error   // SELECT FOR UPDATE
    ReleaseStock(ctx, productID, qty) error
    DeductStock(ctx, productID, qty) error
    AcquireLock(ctx, lock) (*InventoryLock, error)
}
```

---

### Task 3 — Entity Logic

| Entity | Method | Logic |
|--------|--------|-------|
| `Order` | `CalculateTotal()` | Sum ของ OrderItem subtotals |
| `Order` | `UpdateStatus(status)` | State machine validation |
| `Order` | `CanBeCancelled()` | PENDING หรือ CONFIRMED เท่านั้น |
| `Order` | `IsValid()` | ตรวจ CustomerID, Items, TotalAmount |
| `Inventory` | `GetAvailableStock()` | Quantity - ReservedQuantity |
| `Inventory` | `ReserveStock(qty)` | ตรวจ available ก่อน reserve |
| `InventoryLock` | `IsExpired()` | time.Now().After(ExpiresAt) |
| `Cart` | `CalculateTotal()` | Sum ของ CartItem subtotals |
| `Cart` | `IsValid()` | ตรวจ items ไม่ว่าง |
| `Payment` | `MarkCompleted()` | Set status = COMPLETED |
| `Payment` | `Verify()` | Status == COMPLETED && TransactionID != "" |

**Order State Machine:**
```
PENDING → CONFIRMED → PROCESSING → SHIPPED → DELIVERED → REFUNDED
PENDING → CANCELLED
CONFIRMED → CANCELLED
```

---

## Flash Sale PlaceOrder Flow

```
1.  RateLimiterService.ValidateOrderRequest()     → 429 if exceeded
2.  CartRepository.FindByCustomerID()
    cart.IsValid()                                → 400 if empty
3.  for each CartItem:
      InventoryService.ValidateAndReserve()       → 409 if no stock
4.  for each CartItem:
      InventoryService.AcquireLock() [TTL=15min]
5.  Build Order → order.CalculateTotal() + order.IsValid()
6.  OrderRepository.Create()
7.  PaymentService.InitiatePayment()
8.  order.UpdateStatus(CONFIRMED)
9.  go CartService.ClearCart()                    [async goroutine]
10. go NotificationService.Send(ORDER_CONFIRMED)  [async goroutine]
11. return OrderResponse
```

---

## Team Division

| คน | Branch | หน้าที่ |
|----|--------|---------|
| ไผ่ | `feat-A/handler` | API Spec + รับ Request (PlaceOrder Flow) |
| มิว | `feat-A/service` | Business Logic (Order + Inventory + RateLimiter) |
| หนุ่ม | `feat-A/repository` | Database Mapping (Order + Inventory) |
| แฮม | `feat-B/handler` | API Spec + รับ Request (Payment & Notification) |
| ครีม | `feat-B/service` | Business Logic (Payment + Notification) |
| มาร์ค | `feat-B/repository` | Database Mapping (Payment + Notification) |

---

## Running the Project

```bash
go run cmd/server/main.go
```
