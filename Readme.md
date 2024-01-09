## About

Basic coding, REST API test using Golang make a design system and transaction flow.

## Feature list
 - User Side
    - User register
    - User input the address
    - User choose the products to be purchased with subscription and/or one-time purchase scheme
    - User pay the bill
    - User skip the delivery due to certain reasons (ex: They have other agenda that prevent them to receive the delivered goods)
    - User cancel the order
 - Supplier Side
    - Supplier register as seller
    - Supplier create the store and complete the address
    - Supplier create products that can be purchased either daily or one-time purchase
    - Supplier determine the price of each product
    - Supplier determine the selling area
 - Additional
    - If a product can be sold by more than one seller, define the correct algorithm to determine which order to be sent from which seller (assuming the closest mileage and route)!
    - There is a cut-off time everyday, which is the latest time an order can be placed for the next day delivery. All orders placed beyond cut-off time will automatically delivered on the day after tomorrow.

## Built with

- Golang
- MySQL
- Visual Studio Code

## Prerequisites

-   Golang >=1.18
-   Postgres >= 8.0.33

## Local Installation

1. Clone this project
    ```
    git clone https://github.com/berrylradianh/basic-coding-kulina.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
   APP_PORT="${APP_PORT}"
   DB_CONNECTION="${DB_CONNECTION}"
   DB_HOST="${DB_HOST}"
   DB_PORT="${DB_PORT}"
   DB_NAME="${DB_NAME}"
   DB_USERNAME="${DB_USERNAME}"
   DB_PASSWORD="${DB_PASSWORD}"
   SECRET_KEY="${SECRET_KEY}"
    ```

4.  Run the application
    ```
    go run main.go
    ```