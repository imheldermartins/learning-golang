# RQLITE
## Introduction

### Get Image of rqlite using Docker
```bash
docker pull rqlite/rqlite:9.4.5 # stable version
```

### Node 1 - Create and run container as [*Leader*](#leader)
```bash
docker run -d --name rqlite1 \
  -p 4001:4001 -p 4002:4002 \
  rqlite/rqlite:9.4.5 \
  -http-addr 0.0.0.0:4001 \
  -raft-addr 0.0.0.0:4002 \
  -bootstrap-expect 3
```
**Explanation of each line above terminal script:**

1. ```sh
    # run image rqlite with`rqlite1` as name
    docker run -d --name rqlite1 \ ...
    ```
2. ```sh
    # 4001: http port rqlite | 4002: Internal port Raft communication
    -p 4001:4001 -p 4002:4002 \
    ```
3. ```sh
    # Stablish  a default version to preventing broken
    rqlite/rqlite:9.4.5
    ```
4. ```sh
    # Define a global listening port for all interfaces
    -http-addr 0.0.0.0:4001
    # For example '/status', '/db/execute' and '/db/query' external calls
    ```
5. ```sh
    # Expects 3 nodes availables
    -bootstrap-expect 3
    ```

## Node 2
Adding some parameters
```sh
docker run -d --name rqlite2 \
  -p 4003:4001 -p 4004:4002 \
  rqlite/rqlite:9.4.5 \
  -http-addr 0.0.0.0:4001 \
  -raft-addr 0.0.0.0:4002 \
  -join http://host.docker.internal:4001
```

## Leader