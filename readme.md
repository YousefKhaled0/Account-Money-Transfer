# Account Transfer Application

This application allows you to transfer money between accounts and retrieve account information.

## Getting Started

Follow these instructions to get your application up and running.

### Installation

1. Clone the repository:

  ```bash
  git clone https://github.com/YousefKhaled0/account-funds-transfer.git
  ```

2. Run the application:

  ```bash
  go run main.go
  ```

  Your app should now be running at http://localhost:8080.

### Running Tests

To run tests for your application, follow these steps:

1. Open a terminal and navigate to your project directory.

2. Run the tests using the following command:

  ```bash
  go test tests
  ```

  This will execute your test suite and provide feedback on the test results.

## APIs Documentation

### Transfer Funds

###### Transfer money from one account to another.

**URI: /transfer**

**Method: POST**

**Request Body: JSON**

**Example request body:**

```json
{
    "from":"3d253e29-8785-464f-8fa0-9e4b57699db9",
    "to":"17f904c1-806f-4252-9103-74e7a5d3e340",
    "amount": 10
}
```

**Response: JSON**

**Example response:**

```json
{
    "fromAccount": "3d253e29-8785-464f-8fa0-9e4b57699db9",
    "toAccount": "17f904c1-806f-4252-9103-74e7a5d3e340",
    "transferAmount": 10,
    "balanceAfterTransfer": "77.11"
}
```



### Get All Accounts
###### Retrieve a list of accounts.

**URI /accounts**

**Method: GET**

**Query Parameters:**

**page (optional): Page number (default is 1).**

**pageSize (optional): Number of accounts per page (default is 10).**

**Response: JSON**

**Example response:**

```json

[
    {
        "id": "03eb9399-e526-431f-812f-2fda01659022",
        "name": "Browsecat",
        "balance": "3172.14"
    },
    {
        "id": "04943793-8f35-4d73-aa93-0ef2da57d22e",
        "name": "Flashset",
        "balance": "3724.11"
    },
    {
        "id": "054f9801-f070-4f16-bde7-155430417d43",
        "name": "Quimba",
        "balance": "1011.56"
    }
]
```
