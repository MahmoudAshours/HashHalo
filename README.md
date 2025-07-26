# HashHalo
![](HashHalo.png)
HashHalo is a lightweight simulation of consistent hashing to demonstrate how distributed load balancing works across server nodes. The project simulates key distribution across a hash ring and visualizes the minimal reallocation that occurs when servers join or leave the system.

## Project Objectives

- Implement consistent hashing from scratch.

- Simulate server nodes and dynamic joining/leaving.

- Efficiently assign and reassign "keys" (requests/users/files) to nodes.

- Measure impact on reallocation and performance.

Optional: Add UI or metrics to visualize hash ring and load distribution.

## Key Concepts

- Consistent Hashing: A technique to distribute data across a cluster in a way that minimizes rebalancing when nodes join or leave.

- Virtual Nodes: Improve uniformity by assigning multiple positions in the hash ring to a single server.

- Hash Ring: A circular space (0 to 2³² or based on hash function output) where keys and nodes are mapped using a hash function.

## Tech Stack

Backend (Core Logic):
- Golang (ideal for low-latency logic, concurrency)

Optional Add-ons:
Redis or MongoDB – simulate distributed storage

- Vue.js/React – to visualize hash ring and server loads

- gRPC/REST – simulate service calls

## Project Structure
```
HashHalo/
│
├── src/
│   ├── hash_ring.go     # Implements the hash ring logic
│   ├── server_node.go   # Represents a server
│   ├── main.go          # Simulation / entry point
│
├── test/
│   └── hash_test.go     # Unit tests for consistent hashing
│
├── data/
│   └── keys.json        # Sample input keys to distribute
│
├── README.md
└── visualization/       # (Optional) UI to visualize hash ring
```
## Feature Description
- AddNode(nodeID)	Add a server to the ring
- RemoveNode(nodeID)	Remove a server from the ring
- GetNode(key)	Get the server that owns the key
- Rebalance(keys)	Show how few keys are remapped on node changes
- Virtual Nodes Support	Add multiple points per server
- Metrics (Keys/server, reallocation)	Show how well load is balanced

## Example Flow

- Start with 3 servers: A, B, C
- Insert 1000 keys (simulate user sessions or file storage)
- Observe which server handles which key
- Add a new server D
- Show how only ~1/N of the keys are remapped (not all)
- Visualize changes in server load