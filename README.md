# Clean Arch Echo Service Dating

Clean arch echo, with structure repository, service(use case), model (entity), api with routes-middleware-controller.

-   Internal 
    This folder store main code application : controller, helper, command, middleware, repository, routes dan service 

    Controller : 
    This layer receives all requests from the client that fire into our application

    Service :
    This layer is going to handle with real core business logic of our application

    Repository :
    This layer is to connect with a database like Mysql

    Helper : 
    This layer is help business logic on layer service. It will split logic when it more line code will write.

    Command : 
    This layer for migrate, module used application, and command line used application

    Routes :
    This layer for endpoint application

    Middleware : 
    This layer for middleware application for handling db transaction, jwt auth, logger, etc.


-   Package
    This folder for library external used, config application

    External : 
    This layer is to connect with an external service like a redis, mongo, rabbitmq, etc

    Library :
    This Layer is configuration application mysq/redis/mongo, http request used application, etc

-   Resource 
    This folder for store constant, model structure application, response application


## List Library Application

| Library                            | 
| ---------------------------------- | 
| Mysql Driver Gorm                  |
| Logger with Zap Logger             |
| Echo Framework                     |
| Middleware JWT                     |
| Middleware Request ID Echo         |
| Middleware Database Transaction    |
| Depency Injection Go Uber Fx       |
| Viper Config yaml                  |
| Redis                              |

## How To Install & Run

-   rename config.yml.example -> config.yml
-   go mod tidy
-   migrate data : go run . m
-   running : go run . s
-   build : go build .

## How To Build Docker Application 
-   docker build -t dating-api .
-   docker run -d -p 8080:8080 --name dating-api dating-api:latest

## How To Run Unit Testing
-   go mod tidy
-   go test -v ./internal/controller

