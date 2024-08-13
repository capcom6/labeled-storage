# Labeled Key-Value Storage

## Overview

This open-source project provides a storage solution that allows attaching key-value labels to each item. This means items can be retrieved not only by key but also by any combination of labels' values. The project is designed to be extensible with different storage backends and utilizes gRPC for communication.

## Features

- **Label-Based Retrieval**: Retrieve items not only by key but also by any combination of labels' values.
- **Pluggable Storage Backends**: Default in-memory storage with future support for Redis, MongoDB, and S3 for large objects.
- **gRPC Communication**: Utilizes gRPC for efficient and robust communication.

## Use Cases

- **Service registry**: Store and manage service metadata, such as instance IDs, hostnames, and ports, with labels for querying and filtering.
- **Remote config storage**: Store and manage configuration data for applications, with labels for environment, region, or other relevant factors.
- **Content management**: Store and manage content, such as images or documents, with labels for metadata like author, date, or category.

## Installation

To get started, clone the repository and navigate to the project directory:

```sh
git clone https://github.com/capcom6/labeled-storage.git
cd labeled-storage
```

Build the project:

```sh
make
```

## Usage

### Starting the Server

To start the server, run the following command:

```sh
go run main.go
```

### gRPC API

The storage service exposes several gRPC endpoints:

- **Get**: Retrieves the item with the specified key.
- **Find**: Retrieves items with the specified labels.
- **Replace**: Replaces the item with the specified key.
- **DeleteOne**: Deletes the item with the specified key.
- **DeleteMany**: Deletes items with the specified labels.

### Example gRPC Client (Go)

Here is a simple example of how to interact with the storage service using a gRPC client in Go:

```go
package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "github.com/capcom6/labeled-storage/pkg/api"
)

func main() {
    conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewStorageServiceClient(conn)

    // Example: Get an item by key
    resp, err := client.Get(context.Background(), &pb.GetRequest{Key: "exampleKey"})
    if err != nil {
        log.Fatalf("could not get item: %v", err)
    }
    log.Printf("Item: %v", resp.Item)
}
```

## Contributing

We welcome contributions! Please fork the repository and submit pull requests.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
