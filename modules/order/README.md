# Order Module

## Overview
The Order module manages food orders in the system, connecting users with food items. It follows the Clean Architecture pattern with CQRS and includes RPC client/server infrastructure.

## Database Schema

```sql
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(36) NOT NULL,
  `food_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT '1',
  `total_price` float NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `food_id` (`food_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

## Module Structure

```
modules/order/
├── domain/
│   ├── model.go          # Order, OrderUser, OrderFood models
│   └── error.go          # Domain error constants
├── dtos/
│   ├── create_order.dto.go
│   ├── update_order.dto.go
│   ├── get_order.dto.go
│   ├── list_order.dto.go
│   └── delete_order.dto.go
├── infras/
│   ├── repository/
│   │   ├── repo.go
│   │   ├── insert.go
│   │   ├── find.go
│   │   ├── update.go
│   │   ├── delete.go
│   │   └── orderrpcclient/  # RPC clients for Food & User
│   │       ├── rpc_client.go
│   │       ├── get_food.rpc.go
│   │       ├── get_foods.rpc.go
│   │       ├── get_user.rpc.go
│   │       └── get_users.rpc.go
│   └── controller/
│       ├── controller.go
│       ├── create_order_api.go
│       ├── get_order_api.go
│       ├── list_order_api.go
│       ├── update_order_api.go
│       ├── delete_order_api.go
│       └── orderrpcserver/   # RPC server for Order
│           ├── controller.go
│           ├── get_order.rpc.go
│           └── get_orders.rpc.go
├── service/
│   ├── create_order.svc.go
│   ├── get_order.svc.go
│   ├── list_order.svc.go
│   ├── update_order.svc.go
│   └── delete_order.svc.go
└── module.go              # Module setup and DI
```

## API Endpoints

### HTTP Endpoints
- `POST /v1/orders` - Create a new order
- `GET /v1/orders/:id` - Get order by ID (with user and food data)
- `GET /v1/orders` - List orders (supports filtering and pagination)
- `PATCH /v1/orders/:id` - Update order
- `DELETE /v1/orders/:id` - Delete order

### RPC Endpoints
- `POST /v1/rpc/orders/get-order` - Get single order by ID
- `POST /v1/rpc/orders/get-orders` - Get multiple orders by IDs

## Features

### 1. CQRS Pattern
- **Commands**: CreateOrder, UpdateOrder, DeleteOrder
- **Queries**: GetOrder, ListOrder

### 2. RPC Integration
The order module uses RPC clients to fetch related data:
- **Food RPC Client**: Fetches food details from the food module
- **User RPC Client**: Fetches user details from the user module

### 3. Data Population
When fetching orders, the module automatically populates:
- `order.User` - User information (id, email, first_name, last_name, phone)
- `order.Food` - Food information (id, name, price)

### 4. Filtering Support
List orders supports filtering by:
- `user_id` - Filter by user
- `food_id` - Filter by food
- `status` - Filter by status
- Pagination with `page` and `limit`

## Example Usage

### Create Order
```bash
POST /v1/orders
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "food_id": 123,
  "quantity": 2,
  "total_price": 50.00
}

Response: {"data": 1}  # Order ID
```

### Get Order
```bash
GET /v1/orders/1

Response:
{
  "data": {
    "id": 1,
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "food_id": 123,
    "quantity": 2,
    "total_price": 50.00,
    "status": 1,
    "created_at": "2025-11-24T10:00:00Z",
    "updated_at": "2025-11-24T10:00:00Z",
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe",
      "phone": "1234567890"
    },
    "food": {
      "id": 123,
      "name": "Pizza",
      "price": 25.00
    }
  }
}
```

### List Orders
```bash
GET /v1/orders?user_id=550e8400-e29b-41d4-a716-446655440000&page=1&limit=10

Response:
{
  "data": [...],
  "paging": {
    "page": 1,
    "limit": 10,
    "total": 5
  }
}
```

## Related Infrastructure

### User RPC Server (Created)
The user module now has an RPC server at:
- `POST /v1/rpc/users/get-user` - Get single user
- `POST /v1/rpc/users/get-users` - Get multiple users

### Configuration
Add these service URIs to your environment/flags:
```bash
--food-service-uri=http://localhost:3600/v1/rpc/foods
--user-service-uri=http://localhost:3600/v1/rpc/users
```

## Architecture Benefits

1. **Decoupling**: Orders don't directly access user/food databases
2. **Scalability**: Each service can scale independently
3. **Performance**: Bulk RPC calls minimize network overhead
4. **Maintainability**: Clear separation of concerns
5. **Reusability**: RPC servers can be called by any module

