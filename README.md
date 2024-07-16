## Welcome to DEUNA Payment System

### Introduction
This development has been created for evaluation purposes. It was designed under **Hexagonal Architecture** (so-called "**Ports and Adapters**").

### Note: Please the document until the end, there are several things WIP.
(Sorry for the above text in bold)


## Technologies used

- Golang 1.22
- Docker 27.0.3
- Docker Compose 2.28.1
- PostgreSQL 16.3

## Bootstrapping

At the moment of writting this document, the complete app hasn't been dockerized yet, so we need to setup the database:
```bash
# In the root of the project, I mean ./deuna-payment-system
$ docker compose up -d # Watch out, we're running this docker compose in background
```

There are two services that need to run in different terminals.

```bash
# Run deuna-payment-gateway
# With your prompt at the root of the project, you have to move to main.go of deuna-payment-gateway
$ cd cmd/deuna-payment-gateway
$ go run main.go # This server will running at port 8081

# The second service will be executed by the same way as first
$ cd cmd/deuna-acquiring-bank-simulator
$ go run main.go # This server will running at port 8082
```

## Way of work

1. First of all, we need to create a merchant, let me tell you how we can do it:
   ```bash
    curl --location 'localhost:8082/api/v1/merchants/new' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "name": "Pepe Grillo",
      "balance": 0.0,
      "notification_email": "pepe.grillo@deuna.com"
    }'
   ```
2. The merchant has been stored disabled status, hence we need update that status to enabled ()
   ```bash
   # As you can see, it should be execute by the service 'deuna-acquiring-bank-simulator'
    curl --location --request PATCH 'localhost:8082/api/v1/merchants/change_status/48616812' \
    --header 'Content-Type: application/json' \
    --data '{
      "status": true
    }'
   ```
3. With the enabled merchant we can start to create transactions and its respective refunds if it needed.
   ```bash
   # The transactions will be executed into the first service 'deuna-payment-gateway'
   curl --location 'localhost:8081/api/v1/payments/new' \
    --header 'Content-Type: application/json' \
    --data '{
      "amount": 500.00,
      "currency": "USD",
      "card_number": "4242424242424242",
      "expiry": "12/24",
      "cvv": "123",
      "transaction_type": "DEPOSIT",
      "merchant_code": "06195832"
    }'
   ```
4. If we want to refund this transaction, we need to indicate the params as follow:
   ```bash
   curl --location 'localhost:8081/api/v1/payments/refund' \
    --header 'Content-Type: application/json' \
    --data '{
        "transaction_id": "01J2XMFS5CM8VYGKAYB38XEQ5D",
        "merchant_code": "06195832",
        "amount": 100.00,
        "transaction_type": "REFUND"
    }'
   ```
5. Do we want to info about a merchant by its merchant_code/merchant_id (it's the same thing => technical debt)?
   ```bash
   curl --location --request GET 'localhost:8082/api/v1/merchants/by_code/48616812' \
    --header 'Content-Type: application/json' \
    --data '{
        "enabled": true
    }'
   ```
6. The same thing with id:
   ```bash
   curl --location --request GET 'localhost:8082/api/v1/merchants/by_id/9' \
    --header 'Content-Type: application/json' \
    --data '{
        "enabled": true
    }'
   ```

## WIP
  1. Write the tests.
  2. Logger service (it will be injected into each layer) for tracking operations.
  3. Reporting of every operations such as payments as well as refunds.
  4. Fix some weird behaviors.
  5. Dockerize completely the application, this includes both servers.

  

