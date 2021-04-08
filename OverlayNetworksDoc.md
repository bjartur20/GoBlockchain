# Overlay Networkd

## How to run
### Compile
```bash
go build cmd/daemon/daemon.go
```
### Run a node 
```bash
daemon
```

## How it works
The daemon starts by intializing the network's configuration from the flags. Then, it initializes the naming service, starts the DHT node and the router.

### DHT Node/Kademlia
The infohash is stored as a constant so it's the same for each node. This is decoded at the beginning of starting the node. After decoding the infohash, we create the DHT node with an empty configuaration file, so that our node finds the currently running nodes that we have on the public network.
After creating the node, we start it and a goroutine to drain the nodes that our new node can now connect to.

### Routes


## Tests


## Authors
Bjartur Þórhallsson

Laurynas Cetyrkinas

Ýmir Þórleifsson