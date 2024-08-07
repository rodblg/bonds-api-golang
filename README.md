# bonds-api-golang

This repository exposes a REST API that mimics the purchase flow of bonds. Allows users to both publish and buy documents. 
I decided to build the service with Go due to its simplicity for building http services, implementing chi as the Go router.

# Features
- Register users and login with your defined email/password and grant authorization 
- Publish new bonds
- Get Bonds by Id
- Get all Bonds
- Buy available Bonds

# Technologies
- Go 1.20
- Mongo
- Docker
- Chi 

# Installation

Setup and run in a local environment:

1) Select the main branch

2) Clone this repository into your local environment, opening a cmd terminal in the path that you want your files: 
```
git clone https://github.com/rodblg/bonds-api-golang.git
```
3) Open and run your docker daemon 

4) With the docker-compose.yml file pull the mongo image and build the api container. Run this command in your terminal inside the root of your cloned repository. The -d tag is optional if you don't want to see the logs of the containers building.
```
docker-compose up -d
```
5) You are all set to make http requests into the service

# Path operations

This project defines 6 path operations

- Login: /auth/login 

This POST endpoint allows you to log into an existing account with a username and password. The endpoint for login into an existing account requires a basic authentication header:
Username: your user email address
Password: password associated with your user

Once your credentials are validated you will receive your token authorization as a bearer token in your authorization header in the response

- CreateUser: /user/
This Post endpoint allows you to signup and create a new user account. Requires a JSON object in the request Body. The format of the object is as follows:
```
{
  "name": "John",
  "last_name": "Smith",
  "email": "john.doe@example.com",
  "password":"testing"
}
```
For the following path operations you need the token authorization as a beare token in your requests.

- NewBond: /user/bond
This Post endpoint allows you to define as an authorized user new bonds. This Requires a JSON object in the request Body. The format of the object is as follows: 
```
{
"name": "12200-Year Treasury Note",
"face_value": 10000.0,
"current_value": 9850.25,
"isin": "912828VM6",
"issuer": "Department of the Treasury",
"interest_rate": 0.025,
"interest_payment_frequency": "Semi-annual",
"maturity_date": "2023-11-15T00:00:00Z",
"description": "This is a sample 10-year Treasury Note issued by the government."
}
```
- GetBond: /user/bond/{id}
This Get endpoint allows you to get a defined bond with its unique ID

- GetAllBonds: /user/bond 
This Get endpoint allows you to get all bonds and visualize the ID from all saved bonds

- BuyBond: /user/bond/buy/{id}
This Get endpoint allows you to buy an available bond

# Postman configuration
You can use the postman collection to make new requests into the API following this steps:

1) Download the file ./postman/bondApi.postman_collection.json 
2) Import into your postman workspace
