# Receipt Processor

A RESTful API service that processes receipts and calculates reward points based on receipt data.

## Project Structure

```
receipt-processor-web-service/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   └── router.go           # API routes and handlers
│   ├── models/
│   │   └── types.go            # Data structures
│   └── service/
│       └── points_calculator.go # Points calculation logic
├── test/
│   └── main_test.go            # Integration tests
└── README.md
```
## API Endpoints

### 1. Process Receipt
- **POST** `/receipts/process`
- Processes a receipt and returns a unique ID
- Example request:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
  ],
  "total": "6.49"
}
```
- Example response:
```json
{ 
    "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" 
}
```

### 2. Get Points
- **GET** `/receipts/{id}/points`
- Returns points calculated for a given receipt ID using the ID received from the previous request
- Example request:
```bash
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```
- Example response:
```json
{
  "points": 32
}
```

## Running the Application

```bash
# To from workspace folder, start the server
go run receipt-processor/cmd/main.go 

# The server will start on port 8080
```

## Testing

```bash
# To from workspace folder, run tests with verbose output
go test -v receipt-processor/test/main_test.go
```

## Points Calculation Rules

Points are awarded based on the following rules:
- One point for every alphanumeric character in the retailer name
- 50 points if the total is a round dollar amount
- 25 points if the total is a multiple of 0.25
- 5 points for every two items
- Points based on item description length and price
- 6 points if the day in the purchase date is odd
- 10 points if the time of purchase is between 2:00pm and 4:00pm