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

### Logging and Monitoring
- Prometheus and Grafana: For monitoring and visualizing metrics like request count, latency, etc.

### Containerization and Deployment
- Docker: To containerize the load balancer for easy deployment.
- Kubernetes: If you plan to deploy the load balancer in a Kubernetes cluster.

### CI/CD
- GitHub Actions or CircleCI: To automate the build, test, and deployment pipeline.