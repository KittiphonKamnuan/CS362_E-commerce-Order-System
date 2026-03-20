# CS362 Final Project Design Assignment
## E-commerce Order System — Core Feature: Flash Sale with Rate Limiting

---

## Step 1: Data Design — Analysis Class Diagram

**File:** `classdiagram.mmd`

Analysis Class Diagram ครอบคลุม Entities ทั้งหมดที่เกี่ยวข้องกับระบบ Flash Sale:

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
- `Order` 1 --> 0..* `Payment` (1-to-many, supports refund flow)
- `Inventory` 1 --> * `InventoryLock` (protected by locks during flash sale)
- `Order` ..> `RateLimiter` (validates flash sale request)

---

## Step 2: Architectural Mapping — Layered Architecture

**Git Action:** โครงสร้างโฟลเดอร์นี้สร้างบน Git Repository จริงของกลุ่ม

```
CS362_E-commerce/
├── go.mod
├── main.go
├── entity/
│   ├── enums.go           # OrderStatus, PaymentMethod, PaymentStatus, NotificationType, ProductStatus
│   ├── customer.go        # Customer, Address
│   ├── product.go         # Product, Category
│   ├── inventory.go       # Inventory, InventoryLock
│   ├── cart.go            # Cart, CartItem
│   ├── order.go           # Order, OrderItem
│   ├── payment.go         # Payment
│   └── notification.go    # Notification
├── repository/
│   ├── order_repository.go
│   ├── inventory_repository.go
│   ├── payment_repository.go
│   ├── product_repository.go
│   ├── cart_repository.go
│   └── notification_repository.go
├── service/
│   ├── order_service.go
│   ├── inventory_service.go
│   ├── payment_service.go
│   ├── cart_service.go
│   ├── product_service.go
│   ├── notification_service.go
│   └── rate_limiter_service.go
├── controller/
│   ├── order_handler.go
│   ├── product_handler.go
│   ├── cart_handler.go
│   ├── payment_handler.go
│   └── notification_handler.go
└── dto/
    ├── request/
    │   ├── order_request.go
    │   ├── cart_request.go
    │   └── payment_request.go
    └── response/
        ├── common_response.go
        ├── order_response.go
        └── product_response.go
```

**Layer Mapping จาก Class Diagram:**

| Layer | หน้าที่ | ไฟล์ตัวอย่าง |
|-------|---------|--------------|
| Entity | Domain objects + business logic | `entity/order.go` |
| Repository | Data access interface (persistence) | `repository/order_repository.go` |
| Service | Business logic + orchestration | `service/order_service.go` |
| Controller | HTTP handler + routing | `controller/order_handler.go` |
| DTO | Request/Response data transfer | `dto/request/order_request.go` |

---

## Step 3: Interface & Controller Contract

### Task 1 — API Specification (Controller Layer)

#### Order Endpoints

| Method | Path | Description | Status Code |
|--------|------|-------------|-------------|
| POST | `/api/v1/orders` | Place Order (Flash Sale) | 201 / 429 / 409 |
| GET | `/api/v1/orders/{orderId}` | Get Order by ID | 200 / 404 |
| GET | `/api/v1/orders?page=1&pageSize=10` | List Orders (paginated) | 200 |
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
{
  "success": true,
  "data": {
    "orderId": "order_001",
    "status": "CONFIRMED",
    "totalAmount": 2499.00
  }
}
```

**429 Too Many Requests (Rate Limited):**
```json
{
  "success": false,
  "error": { "code": "RATE_LIMIT_EXCEEDED", "retryAfter": 60 }
}
```

**409 Conflict (Out of Stock):**
```json
{
  "success": false,
  "error": { "code": "INSUFFICIENT_STOCK", "productId": "prod_xyz" }
}
```

---

### Task 2 — Go Interface Files (Service & Repository Layers)

#### `service/order_service.go`
```go
type OrderService interface {
    PlaceOrder(ctx, customerID, req) (*OrderResponse, error)
    GetOrderByID(ctx, customerID, orderID) (*OrderResponse, error)
    GetOrderHistory(ctx, customerID, page, pageSize) (*OrderListResponse, error)
    CancelOrder(ctx, customerID, orderID) (*OrderResponse, error)
}
```

#### `service/rate_limiter_service.go`
```go
type RateLimiterService interface {
    IsLimitExceeded(ctx, customerID) (bool, error)
    ValidateOrderRequest(ctx, customerID) error
    RecordRequest(ctx, customerID) error
    ResetCounter(ctx, customerID) error
}
```

#### `service/inventory_service.go`
```go
type InventoryService interface {
    ValidateAndReserve(ctx, productID, qty) error
    AcquireLock(ctx, orderID, productID, customerID, qty) (*InventoryLock, error)
    ReleaseLock(ctx, lockID) error
    FulfillOrder(ctx, orderID) error
}
```

#### `repository/order_repository.go`
```go
type OrderRepository interface {
    Create(ctx, order) (*Order, error)
    FindByID(ctx, orderID) (*Order, error)
    FindByCustomerID(ctx, customerID, offset, limit) ([]*Order, int, error)
    UpdateStatus(ctx, orderID, status) error
}
```

#### `repository/inventory_repository.go`
```go
type InventoryRepository interface {
    FindByProductID(ctx, productID) (*Inventory, error)
    ReserveStock(ctx, productID, qty) error   // SELECT FOR UPDATE
    ReleaseStock(ctx, productID, qty) error
    DeductStock(ctx, productID, qty) error
    AcquireLock(ctx, lock) (*InventoryLock, error)
}
```

---

### Task 3 — Entity Logic Refinement

#### `entity/order.go`

| Method | Logic |
|--------|-------|
| `CalculateTotal()` | Sum ของ `OrderItem.GetSubtotal()` ทุกรายการ |
| `UpdateStatus(status)` | Validate state machine transitions ก่อน update |
| `CanBeCancelled()` | `true` เฉพาะ PENDING หรือ CONFIRMED |
| `IsValid()` | ตรวจ CustomerID, Items, TotalAmount > 0 |

**State Machine:**
```
PENDING → CONFIRMED → PROCESSING → SHIPPED → DELIVERED → REFUNDED
PENDING → CANCELLED
CONFIRMED → CANCELLED
```

#### `entity/inventory.go`

| Method | Logic |
|--------|-------|
| `GetAvailableStock()` | `Quantity - ReservedQuantity` |
| `ReserveStock(qty)` | ตรวจ available stock ก่อน เพิ่ม ReservedQuantity |
| `ReleaseStock(qty)` | ตรวจ over-release ก่อน ลด ReservedQuantity |
| `DeductStock(qty)` | ลด Quantity จริง (permanent fulfillment) |
| `InventoryLock.IsExpired()` | `time.Now().After(ExpiresAt)` |

#### `entity/cart.go`

| Method | Logic |
|--------|-------|
| `CalculateTotal()` | Sum ของ `CartItem.GetSubtotal()` |
| `IsValid()` | ตรวจว่า items ไม่ว่างเปล่า |
| `Clear()` | Reset items slice และ TotalPrice = 0 |

#### `entity/payment.go`

| Method | Logic |
|--------|-------|
| `MarkCompleted(txID, paidAt)` | Set TransactionID, PaidAt, Status = COMPLETED |
| `MarkFailed(reason)` | Set FailureReason, Status = FAILED |
| `Verify()` | `Status == COMPLETED && TransactionID != ""` |
| `IsValid()` | ตรวจ OrderID และ Amount > 0 |

---

## Flash Sale PlaceOrder Flow

```
PlaceOrder(ctx, customerID, req):
  1. RateLimiterService.ValidateOrderRequest()    → 429 if exceeded
  2. CartRepository.FindByCustomerID()
     cart.IsValid()                               → 400 if empty
  3. for each CartItem:
       InventoryService.ValidateAndReserve()      → 409 if no stock
  4. for each CartItem:
       InventoryService.AcquireLock() [TTL=15min]
  5. Build Order from Cart
     order.CalculateTotal()
     order.IsValid()
  6. OrderRepository.Create()
  7. PaymentService.InitiatePayment()
  8. order.UpdateStatus(CONFIRMED)
  9. go CartService.ClearCart()                   [async goroutine]
 10. go NotificationService.Send(ORDER_CONFIRMED) [async goroutine]
 11. return OrderResponse
```

---

## Running the Project

```bash
# Build
go build ./...

# Run server (port 8080)
go run main.go

# Health check
curl http://localhost:8080/health
```
