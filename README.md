A RESTful API that allows users to search the underlying store of cars based on certain search parameters. The API is built using Go Lang. The code is bifurcated into 4 main areas - domain, usecases, interfaces, and infrastructure.

## Domain
Domain encapsulates the core entities of the application. In this case, since we are dealing with Cars, this becomes our core entity. Any business logic pertaining to cars must be defined here. Let's say for example in the future we can to add provisions to the API to add and remove cars. And while doing so, certain business logic needs to be accomplished. For example Rodo doesn't allow cars older than 25 years. Since this is a pure business logic it must reside in the domain layer. A method to AddCars will be added and this logic will be implemented within that method.

## Usecases
A level above the domain is the usecases layer. Usecases encapsulate application spefic operations that we wish to execute against a business entity. In the example problem statement, we want to perform a search. It becomes crucial to see the search as a resource in terms of REST and map it to a usecase. Search becomes the operation we want to carry out.
