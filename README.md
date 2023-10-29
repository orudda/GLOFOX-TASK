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
- **GET /classes/{classId}**: Get a class of Id classId.
- **UPDATE /classes/{classId}**: Update class of Id classId.
- **DELETE /classes/{classId}**: Delete class of Id classId.

Post example:
```json
{
    "Name": "Test",
    "StartDate": "2023-12-01",
    "EndDate": "2023-12-20",
    "Capacity": 10
}
```

### Bookings

- **POST /bookings**: Create a booking for a class.
- **GET /bookings**: Get a list of all bookings.
- **UPDATE /bookings/{bookingId}**: Update booking of Id bookingId.
- **DELETE /bookings/{bookingId}**: Delete booking of Id bookingId.

Post example:
```json
{
    "MemberName": "test",
    "Date": "2023-12-21"
}
```

## Running Tests

To run unit tests, use the following command:

```shell
go test -v ./tests/
```

## Tecnical Informations

 I will explain the decision-making process for each story and the acceptance criteria.
First of all, it was decided to use a simple architecture with the following format:

project/

&nbsp;    |── main.go

&nbsp;    |── controllers/

&nbsp;    |── models/

&nbsp;    |── services/

&nbsp;    |── utils/

&nbsp;    |── tests/

#### Story - Create Classes

- **Implement an API to create classes(\`/classes\`). Assume this api doesn't need to have any authentication to start with.**
Concluded. The first route was created using localhost. It is necessary to pass a JSON as defined previously in API Endpoints.
<br>

- **Few bare minimum details we need to create classes are - class name, start_date, end_date, capacity. For now, assume that there will be only one class per given day. Ex: If a class by name pilates starts on 1st Dec and ends on 20th Dec, with capacity 10, that means Pilates has 20 classes and for each class the maximum capacity of attendance is 10.**
The code was mostly developed by saving just one example with a start and end date. Then it was divided into days.
<br>

- **No need to save the details in any database. Maintain an in-memory array or a file to save the info. (If you want to use the database, that's fine as well).**
To keep this simple but usable it was decided to use the memory array. To simulate a database, each class received an ID.
<br>

- **Use Restful standards and create the api endpoint with proper success and error responses.**
As described in the API Endpoints, the complete CRUD was done using Restful standards. Most cases were covered. Ex: when you try to get, delete or update a class that does not exist, the error description is returned.

#### Story - Book for a class

- **Implement an API endpoint (\`/bookings\`). Assume this api doesn't need to have any authentication to start with.**
For each booking, its ID was created to be stored simulating a database and classId to know in which class it would be allocated.
<br>

- **Few bare minimum details we need for reserving a class are - name(name of the member who is booking the class), date(date for which the member want to book a class)**
As shown in the last topic, in addition to the minimum requirements, Id and IdClass was added.
<br>

- **No need to save the details in DB. If you can maintain the state in an in memory array or a file is good to start with. But no constraints if you want to use a database to save the state.**
The same thing as described in the classes DB
<br>

- **Use REST standards and think through basic api responses for success and failure.**
Most cases were covered. Ex: when you try to add booking to a class that doesn't exist.
<br>

- **No need to consider the scenario of overbooking for a given date. Ex: 14th Dec having acapacity of 20 , but the number of bookings can be greater than 20.**
When a booking id is allocated to a class, its capacity is reduced. The same occurs when the booking is deleted and then the capacity is increased again. An error is thrown when trying to allocate booking when a class no longer has capacity.
