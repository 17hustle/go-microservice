# simple-microservice
A simple microservice in GO
Monolith
A monolithic architecture is a traditional software design approach where an application is built as a single, unified unit. All components of the application, such as the user interface, business logic, and database management, are tightly integrated and interdependent. In a monolithic application:

All functionalities are managed within a single codebase.
The application is typically deployed as one executable or package.
Scaling often requires scaling the entire application, even if only one part needs more resources.
Advantages of Monolithic Architecture:
Easier to develop initially, especially for small applications.
Simpler to deploy since there’s only one package to manage.
Easier to test as everything is in one place.
Disadvantages of Monolithic Architecture:
As the application grows, it becomes harder to maintain, modify, and deploy.
Any changes or updates require redeploying the entire application.
Scaling can be inefficient, as you must scale the entire application even if only a specific function needs more resources.
Microservices
A microservices architecture breaks down a large application into smaller, loosely coupled services, each responsible for a specific functionality. These services communicate with each other over APIs, often using protocols like HTTP or messaging queues. Each microservice is independently deployable, scalable, and can be developed using different programming languages and technologies.

Advantages of Microservices Architecture:
Easier to maintain and update because each service is independent.
Enables more efficient scaling, as only the services that need more resources are scaled.
Encourages the use of the best technology for each service, leading to better overall performance.
Disadvantages of Microservices Architecture:
More complex to develop and manage due to the distributed nature of services.
Requires more careful design of inter-service communication and data consistency.
Deployment and testing can be more challenging.
Docker's Role in Monoliths and Microservices
Docker is a platform that allows developers to package applications and their dependencies into containers. Containers are lightweight, portable, and consistent environments that ensure applications run the same way regardless of where they are deployed.

How Docker Relates to Monoliths:
Containerization: Docker can be used to package a monolithic application into a single container, making it easier to deploy across different environments (e.g., development, testing, production).
Isolation: Docker provides an isolated environment, reducing conflicts between dependencies on the host machine and the application.
How Docker Relates to Microservices:
Service Isolation: Each microservice can be containerized and run in its own Docker container, allowing for independent deployment and scaling.
Consistency and Portability: Docker ensures that microservices work consistently across various environments.
Orchestration: Docker can be combined with orchestration tools like Kubernetes to manage the deployment, scaling, and operation of microservices in a distributed system.
In summary, Docker plays a crucial role in both monolithic and microservices architectures by providing a consistent, portable, and isolated environment for deploying and running applications.

Packages used explained:
1. fmt
Purpose: The fmt package implements formatted I/O with functions similar to C's printf and scanf.

Common Uses:

Printing text to the console (standard output) using functions like fmt.Println, fmt.Printf, etc.

2. log
Purpose: The log package provides a simple logging mechanism. It allows you to output messages for debugging or logging errors.

Common Uses:

Logging messages with timestamps.
Logging errors and other information to the console or to a file.
Terminating the program with log.Fatal if a critical error occurs.

3. net/http
Purpose: The net/http package provides HTTP client and server implementations. It's the standard package for building web servers and making HTTP requests in Go.

Common Uses:

Creating a web server using http.ListenAndServe.
Handling HTTP requests using http.HandleFunc or http.Handle.
Making HTTP requests to external services using http.Get, http.Post, etc

4. os
Purpose: The os package provides a platform-independent interface to operating system functionality, such as dealing with environment variables, file operations, and program exit.

Common Uses:

Reading environment variables using os.Getenv.
Exiting the program with a specific status code using os.Exit.
Interacting with the file system to create, read, write, and delete files.

5. github.com/golang-jwt/jwt:
For creating, parsing, and validating JSON Web Tokens (JWTs)

What does the program do??
JWT Authentication in Go
Overview
This Go application demonstrates how to use JSON Web Tokens (JWTs) for authentication. It includes functions to generate a JWT and protect an endpoint by verifying the token's validity.

Key Features
Token Generation:

The GetJWT function creates a JWT signed with a secret key.
The token includes claims such as authorization status, client identification, audience (aud), issuer (iss), and expiration time (exp).
Token Verification:

The isAuthorized function is used as middleware to protect routes.
It verifies the JWT provided in the request header, checking:
The signing method to ensure it is HMAC.
The token's audience (aud) and issuer (iss) against expected values.
If the token is valid and meets the criteria, the request is passed to the protected endpoint.
Protected Endpoint:

The / route is protected by isAuthorized middleware.
The Index function generates a JWT and returns it to the client if the token generation is successful.
Setup
Environment Variable:

Ensure the environment variable SECRET_KEY is set with a secret key for signing the JWT.
Running the Application:

Start the server by running the Go application. The server listens on port 8080.
Access the protected route by sending a request with the JWT token.
Example
To get a JWT, send a request to the / endpoint. If the token is successfully generated, it will be returned in the response.

Error Handling
If token generation fails or the token is invalid, appropriate HTTP error responses are returned.


