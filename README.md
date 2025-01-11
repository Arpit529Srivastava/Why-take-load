# Simple Load Balancer with Round Robin Algorithm

### This project is a basic implementation of a load balancer in Go using the Round Robin algorithm. The load balancer distributes incoming HTTP requests to a set of backend servers in a circular order. The project also has Three RESTApi.

## A load balancer performs the following functions:

- Distributes client requests/network load efficiently across multiple servers
- Ensures high availability and reliability by sending requests only to servers that are online
- Provides the flexibility to add or subtract servers as demand dictates

## Therefore our goals for this project are to:

- Build a load balancer that can send traffic to two or more servers.
- Health check the servers.
- Handle a server going offline (failing a health check).
- Handle a server coming back online (passing a health check).

## Tech Stack Used are:

### Web Framework

- Net/HTTP
- Gin

### Algorithm Implementation

- Round Robin Logic: Custom implementation of the Round Robin algorithm to distribute requests evenly across multiple backend servers.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/arpit529srivastava/Why-Take-Load.git
   cd your-repo-name
   ```

2. Install dependencies:
   intitalize the go mod
   ```sh
   go mod init ......
   go mod tidy
   ```

### Running the REST APIs

1. first run the other apis go files

   ```sh
   go run api1.go
   go run api2.go
   go run api3.go
   ```

### Running the Load Balancer

1. Start the load balancer:
   ```sh
   go run main.go
   ```

### Access the APIs

1. Access the APIs using a tool like `curl` or Postman:
   ```sh
   curl http://localhost:8080
   ```

### Usage

Once everything is up and running, you can send HTTP requests to the load balancer, and it will distribute them to the backend servers using the Round Robin algorithm.

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request

### License

Distributed under the MIT License. See `LICENSE` for more information.
