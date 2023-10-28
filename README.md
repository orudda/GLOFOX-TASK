# GLOFOX TASK

This is a simple Go application that allows studio owners to manage classes and members to book classes.

## Dependencies

Before you start, make sure you have the following dependencies installed:

- Go (1.14 or higher): [Download Go](https://golang.org/dl/)
- Chi Router: I use Chi as the router for our RESTful API. You can install it using:

    ```shell
    go get github.com/go-chi/chi
    ```

- Testify: I use the `testify` library for writing unit tests. You can install it using:

    ```shell
    go get github.com/stretchr/testify
    ```

## Compile and Run

1. Clone this repository:

    ```shell
    git clone https://github.com/orudda/GLOFOX-TASK
    cd GLOFOX-TASK
    ```

2. Compile and run the application:

    ```shell
    go run main.go
    ```

   The server will start on port 8080. You can access the API at `http://localhost:8080` or `127.0.0.1:8080`.

## API Endpoints

### Classes

- **GET /classes**: Get a list of all classes.
- **POST /classes**: Create a new class.

### Bookings

- **POST /bookings**: Create a booking for a class.

## Running Tests

To run unit tests, use the following command:

```shell
go test -v ./tests/