# MoneyTracker

Welcome to MoneyTracker API, developed using Go! This is a simple project developed for me to learn more about the language.

With this API, users can track their expenses and incomes, gaining better control over their financial activities.

## Getting Started

1. **Clone the Repository:**
```
git clone https://github.com/yourusername/finance-api.git
```
2. **Install Dependencies:**
```
go mod tidy
```
3. **Run the Application:**
```
go run main.go
```
4. **Explore the API:**

- Get all transactions:
```
curl http://localhost:9000/
```
- Get a transaction by ID:
```
curl http://localhost:9000/{id}
```
- Insert a transaction:
```
curl -X POST -H "Content-Type: application/json" -d '{
  "type": "Expense/Income",
  "name": "Name to identify the transaction",
  "category": "Category of the Expense/Income",
  "description": "Description of the transaction",
  "amount": 100.00,
  "date": "2024-01-01T00:00:00Z"
}' http://localhost:9000/{id}
```
- Update a transaction:
```
curl -X PUT -H "Content-Type: application/json" -d '{
  "type": "Expense/Income",
  "name": "Name to identify the transaction",
  "category": "Category of the Expense/Income",
  "description": "Description of the transaction",
  "amount": 100.00,
  "date": "YYYY-MM-DDT00:00:00Z"
}' http://localhost:9000/{id}
```
- Delete a transaction:
```
curl -X DELETE http://localhost:9000/transactions/{id}
```
