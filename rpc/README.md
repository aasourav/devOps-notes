#### What is RPC
RPC is an older protocol that extends conventional local procedure calling(normal functionc call), allowing the called procedure to exist in a different address space than the calling procedure. These two processes can reside on the same system or on different systems connected by a network.

RPC uses (int go (net/rpc)) `binary serialization`(REST used `text serialization`) by the help of go's `gob` encoding.
establishing RPC using `net/rpc` is only for  `go-to-go` communication. for varius language we have to use `protobuf` 

###### Text serialization
```md
    POST /api/user HTTP/1.1
    Host: example.com
    Content-Type: application/json
    Content-Length: 41

    {"id":123,"name":"Alice","age":30}
```

###### Binary serialization
```md
    12 07 73 75 63 63 65 73 73 18 01
```

#### why RPC is faster than REST??


##### REST journey:

###### Layer 7 (Applicatoin layer):  The client generates a JSON object to send to the server
```md

{
  "id": 123,
  "name": "Alice",
  "age": 30
}
```

###### Layer 6 (Presentation Layer):  the client will encode the JSON data into text (which is already in the JSON format, so no further transformation is needed). 

```md
{"id": 123, "name": "Alice", "age": 30}
```

###### Layer 5 (Session Layer): The session layer manages the communication session between the client and server. This layer doesn’t change the data itself but ensures the session is kept open and maintained

###### Layer 4 (Transport Layer):   The transport layer uses TCP to break down the data into smaller segments and ensure reliable delivery. 

```md
Segment 1: {"id": 123, "name":
Segment 2: "Alice", "age": 30}
```

###### Layer 3 (Network Layer):   The data is encapsulated in IP packets. Each packet has a header with information like the source IP address, destination IP address, etc. 

```md
IP Header: [Source IP: 192.168.1.1] [Destination IP: 203.0.113.2]
Payload: {"id": 123, "name": "Alice", "age": 30}
```

###### Layer 2 (Data Link Layer): Ethernet Frame: The data is encapsulated into Ethernet frames for transmission across physical networks (Wi-Fi, Ethernet, etc.).

```md
IP Header: [Source IP: 192.168.1.1] [Destination IP: 203.0.113.2]
Payload: {"id": 123, "name": "Alice", "age": 30}
```

###### Layer 1 (Physical Layer): The data is transmitted over the physical medium, such as Ethernet cables, Wi-Fi, or fiber optic cables. The data is converted into electrical signals or light pulses for transmission through the hardware


##### RPC journey:<u></u>

###### Layer 7 (Applicatoin layer):  The client generates a JSON object to send to the server
```md
{"id": 123, "name": "Alice", "age": 30}
```

###### Layer 6 (Presentation Layer): Data Serialization (Binary Encoding): In the presentation layer, the data is serialized into binary format for transmission. 

```md
08 7b 12 05 41 6c 69 63 65 18 1e
```

###### Layer 5 (Session Layer): The session layer ensures the communication session is maintained over a long period, especially with gRPC (using HTTP/2). It doesn’t modify the data directly but keeps track of the connection.

###### Layer 4 (Transport Layer): TCP Protocol (via HTTP/2): In the transport layer, the Protobuf binary data is sent using HTTP/2. HTTP/2 uses TCP for reliable delivery

```md
Segment 1: 08 7b 12 05
Segment 2: 41 6c 69 63 65
Segment 3: 18 1e
```

###### Layer 3 (Network Layer):   The data is encapsulated in IP packets. Each packet has a header with information like the source IP address, destination IP address, etc. 

```md
IP Header: [Source IP: 192.168.1.1] [Destination IP: 203.0.113.2]c
Payload: 08 7b 12 05 41 6c 69 63 65 18 1e
```

###### Layer 2 (Data Link Layer): Ethernet Frame: The data is encapsulated into Ethernet frames for transmission across physical networks (Wi-Fi, Ethernet, etc.).

```md
Ethernet Header: [Source MAC: 00:1A:2B:3C:4D:5E] [Dest MAC: 00:1A:2B:3C:4D:6F]
Payload: 08 7b 12 05 41 6c 69 63 65 18 1e
```

###### Layer 1 (Physical Layer): The data is transmitted over the physical medium, such as Ethernet cables, Wi-Fi, or fiber optic cables. The data is converted into electrical signals or light pulses for transmission through the hardware


#### Characteristics of RPC

### Why RPC is faster than REST

