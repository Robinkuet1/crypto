# Peer Discovery

**How does a new node discover its peers?**

In an decentralized peer to peer network, it is very important to discover new peers to communicate with.

To join the network a new node has to connect to an already running node. The List of all running nodes will then be provided by the already running node. If the new node does not know an address of an already running node, some central nodes are provided by the project.

After joining every node has a list of some other nodes. To keep the distribution high when retrieving nodes the already connected node doesn't just return it's nodes but requests the list of nodes from its nodes recursively. From them a random selection is choosen. This ensures an random distribution of connections.