package entity

// OrderStatus represents the state of an order
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusConfirmed  OrderStatus = "CONFIRMED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusShipped    OrderStatus = "SHIPPED"
	OrderStatusDelivered  OrderStatus = "DELIVERED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
	OrderStatusRefunded   OrderStatus = "REFUNDED"
)

// PaymentMethod represents payment options
type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard    PaymentMethod = "DEBIT_CARD"
	PaymentMethodBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentMethodEWallet      PaymentMethod = "E_WALLET"
	PaymentMethodCOD          PaymentMethod = "COD"
)

// PaymentStatus represents the state of a payment
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
	PaymentStatusRefunded  PaymentStatus = "REFUNDED"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeOrderConfirmed  NotificationType = "ORDER_CONFIRMED"
	NotificationTypeOrderShipped    NotificationType = "ORDER_SHIPPED"
	NotificationTypeOrderDelivered  NotificationType = "ORDER_DELIVERED"
	NotificationTypePaymentSuccess  NotificationType = "PAYMENT_SUCCESS"
	NotificationTypePaymentFailed   NotificationType = "PAYMENT_FAILED"
	NotificationTypeFlashSaleAlert  NotificationType = "FLASH_SALE_ALERT"
)

// ProductStatus represents the availability of a product
type ProductStatus string

const (
	ProductStatusActive     ProductStatus = "ACTIVE"
	ProductStatusInactive   ProductStatus = "INACTIVE"
	ProductStatusOutOfStock ProductStatus = "OUT_OF_STOCK"
)
