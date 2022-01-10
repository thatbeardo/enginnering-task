A RESTful API that allows users to search the underlying store of cars based on certain search parameters. The API is built using Go Lang. The code is bifurcated into 4 main areas - domain, usecases, interfaces, and infrastructure.

## Domain
Domain encapsulates the core entities of the application. In this case, since we are dealing with Cars, this becomes our core entity. Any business logic pertaining to cars must be defined here. Let's say for example in the future we can to add provisions to the API to add and remove cars. And while doing so, certain business logic needs to be accomplished. For example Rodo doesn't allow cars older than 25 years. Since this is a pure business logic it must reside in the domain layer. A method to AddCars will be added and this logic will be implemented within that method. Domain is highly business specific while application agnostic. It doesn't care what happens outside.

## Usecases
A level above the domain is the usecases layer. Usecases encapsulate application spefic operations that we wish to execute against a business entity. In the example problem statement, we want to perform a search. It becomes crucial to see the search as a resource in terms of REST and map it to a usecase. Search becomes the operation we want to carry out.
Usecases in contrast to domains are both application specific and business specific. Any action that we want to perform on our domain entities passes through usecases. The search operation is a perfect example for a usecase. While it does not modify cars it needs a source of cars to achieve it's goals. In a way it relies on the business logic as well as on what the application does (Deals with cars).

## Interfaces
As the name suggests this layer is responsible to interact with everything outside our application. It can be a database, a call over HTTP/gRPC etc. it can also be a cli application where the call is made using command line arguments. The interface layer manages this boundary and is responsible for translating outside requests into a language that our application can understand. In this case we are making use HTTP calls, decoding the request and serving the end user with search results. It the part of our application that deals with all external agencies. Interfaces are business agnostic - they don't care much about the business logic. At the same time they are very application specific. Depeding on what flavor of external stimulus we are interested in, our interface layer will change.

## Infrastructure
This part of the code is low level implementation details. Although this example doesn't have a database if we were to have a database such as postgres or mySQL, this is where we would place our code. It contians the low level instructions to store and retrieve data. A good test to consider is to make sure that we can swap out databases without changing a single line in usecases or domain. Interfaces by virtue is highly decoupled but if usecases or domains don't finch if databases are swapped out then we have built a very decoupled system.
As an example the infrastructure layer in this project only has a custom logger. The custom loggers implements the interface defined in usecases and hence can be injected from main at runtime. Something similar would happen if we had a database.

## Testing the code
The code is thoroughly tested and there's a script at the root of the project to fire all unit test cases and generate a comprehensive report that can be viewed in the browser. The script is called `cover.sh`. To execute this script you might have to give it permissions to be executed. Add permissions by running

```
sudo chmod 777 cover.sh
```

Once permissions are updated you can execute the script by running

```
./cover.sh
```

## Building the code
Before you begin, make sure you have [Go installed](https://go.dev/doc/install) on your machine. To build the application, navigate to the directory and run
```
go build main.go
```
This will create an executable on your machine in the same directory as you are. The executable will be named 'main'.

## Running the code
Before you can run the code you built, you have to make sure it has permissions to be executed. To grant all permissions run command
```
sudo chmod 777 main
```
To run the executable you just built, execute command.
```
./main
```
You'll get a prompt asking you allow the executable to have networking permissions. Hit allow. The application is now running on your localhost at port 8080.

## Using the application
The default data store is a hardcoded function which returns the [dataset](https://github.com/Honcker/engineering_exercise/blob/main/Exercise_Dataset.json) present in the Github Repository. You can execute the following cURL commands in your SHELL to test the application out.

```
curl --location --request POST 'localhost:8080/api/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "make":"Volvo",
    "model":"XC90",
    "year":2019,
    "budget": 56000
}'
```

## Simulating Errors
The application only supports a `POST` endpoint on path `/api/search`. Making a `GET` request for example will return a `405` status code signyfying an illegal method used to complete the request. Here's a command to get the error:
```
curl --location --request GET 'localhost:8080/api/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "make":"Volvo",
    "model":"XC90",
    "year":2019,
    "budget": 56000
}'
```

Since the application does not support an actual database, it is not possible to inspect results when a genuine database error occurs.