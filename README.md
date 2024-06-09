# CrudGo

This repository provides an example of a CRUD API in Go, implemented using the gin-gonic framework and following the hexagonal architecture. 

To get started, please create a `.env` file in the root directory of the project and populate it with the following variables:

```
MONGO_URI=
MONGO_DBNAME=
MONGO_COLLECTION=
API_TOKEN=
PORT=
ELEMENTS_TO_PAGINATE=
```

The data structure in MongoDB is as follows:

```json
{
  "id": 22,
  "name": "Hong Xiang",
  "email": "hongxiang17@gmail.com",
  "age": 27,
  "entryDate": {"$date": "2018-06-25T21:46:44.000Z"},
  "country": "Chile"
}
```

## Available Endpoints

- `GET /user?page=x` (default page is 0): Retrieves a paginated list of users.
- `GET /user/id/:id`: Retrieves a user by their ID.
- `GET /search/:field/:value` (default page is 0): Searches for users based on a specific field and value.
- `DELETE /user/:id`: Deletes a user by their ID.
- `POST /user/`: Creates a new user.

## Makefile Commands

- `make run`: Runs the project locally.
- `make install`: Installs project dependencies.
- `make build`: Builds the project.
- `make docker-build`: Builds a Docker image of the project.
- `make docker-run`: Runs the project using Docker.
- `make unit-testing`: Runs unit tests.

To run the project locally, execute the following command:

```
make run
```

Please ensure that you load the example data from the `db` folder.


