# my first golang project, as a product management api!

a simple, headless (without a ui, to be exact) product management api which lets you add, retrieve & update products.

## features

- create, read, update & delete products with http requests.

## technologies & libraries used

- echo for handling web requests
- pgx for postgres database operations
- gommon for logging functionality
- docker for database containerization
- integration testing implementation

## prerequisites

before running this application, make sure you have the following installed:

- go
- docker

## docker setup

docker is used to containerize the database setup. to setup docker, you need to run `test_db.sh` file located in `test/scripts`.

## how to run

1. clone the repository:
```bash
git clone https://github.com/nurmorca/letsGO.git
cd letsGO
```

2. set your docker container up as des

3. run the application using
```bash
go run main.go
```

4. open your postman and send your requests. endpoints can be seen on `product_controller.go` file located in `controller`.
