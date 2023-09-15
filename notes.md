# How servers handle multiple connections?

- TCP is the most reliable way for two machines to communicate over the network
- Logical entity that wraps communcation specifics -> socket (TCP, UDP)
- TCP server:
  - simple process that runs in a machine that listens to a port
  - any machine that wants to talk to server has to connect over the port and establishes the connection

Step 1: start listening on the port

- reserve a port and start listening to it

Step 2: wait for a client to connect

- invoke the "accept" system call and wait for a client to connect (this is a blocking call and the server woult not proceed until some client connects)

Step 3: read the request and send the response

- once the connection is established

1. invoke the 'read' system call to read the request (blocking)
2. invoke the 'write' system call to send the response (blocking)
3. close the connection

Step 4:

- multiple requests being processed concurrently
- paralelize the processing:
  - once the client connects,
    - fork a thread to process and respond
  - let main thread come back to 'accept' as quickly as possible

Optimizations:

1. limiting the number of threads
2. add thread pool to save on thread creation time
3. connection timeout
4. TCP backlog queue configuration
