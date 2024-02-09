# Go Load Balancer

This is a lightweight load balancer implemented in Go, leveraging various features of the language and standard library to efficiently distribute incoming requests across backend servers. 

## Features

### Round Robin Selection
The load balancer utilizes a round-robin selection algorithm to evenly distribute incoming requests among available backend servers.

### Reverse Proxy
It utilizes the `ReverseProxy` functionality provided by the standard library to proxy requests to the selected backend servers.

### Mutexes and Atomic Operations
Mutexes and atomic operations are used for concurrency control, ensuring thread safety during operations such as updating internal states and counting statistics.

### Closures and Callbacks
Closures and callbacks are employed for flexibility and modularity, allowing customization of the load balancer's behavior and enabling users to define custom logic for request handling.

### Select Operation
The load balancer utilizes the `select` operation to handle concurrent requests efficiently, allowing it to manage multiple connections concurrently without blocking.

## Potential Improvements

### Heap for Sorting Alive Backends
Implementing a heap-based data structure for sorting alive backend servers can optimize backend selection, reducing the search surface and improving overall performance.

### Statistics Collection
Enhance the load balancer with features to collect and analyze statistics, providing insights into request distribution, backend server health, and overall system performance.

### Weighted Round-Robin/Least Connections
Implement advanced load balancing algorithms such as weighted round-robin or least connections to improve load distribution based on backend server capacities and current loads.

## How to Use

1. Clone the repository:
   ```
   git clone https://github.com/your_username/load-balancer.git
   ```

2. Build the load balancer:
   ```
   go build -o load-balancer main.go
   ```

3. Run the load balancer:
   ```
   ./load-balancer
   ```

4. Access the load balancer via the specified port (default port is 3030).


## Contributing

Contributions to this project are welcome! If you have any ideas for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.