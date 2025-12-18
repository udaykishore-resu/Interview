
# what is the difference between REST and gRPC and GraphQL? When to use what ?
## REST (Representational State Transfer)
- **Architecture:** Resource-based, uses HTTP methods (GET, POST, PUT, DELETE)
- **Data Format:** Typically JSON/XML
- **Communication:** Synchronous HTTP/1.1, stateless
- **Strengths:** Simple, cacheable, human-readable, wide browser support
- **Weaknesses:** Over/under-fetching, multiple round trips for complex data

## gRPC (Google Remote Procedure Call)
- **Architecture:** RPC-based, client-server model
- **Data Format:** Protocol Buffers (binary)
- **Communication:** HTTP/2, supports streaming (unary, server, client, bidirectional)
- **Strengths:** High performance, strongly typed, auto code generation, bidirectional streaming
- **Weaknesses:** Requires HTTP/2 proxy for browsers, less human-readable

## GraphQL
- **Architecture:** Query language for APIs
- **Data Format:** JSON (requests in GraphQL syntax)
- **Communication:** Typically HTTP POST for all operations
- **Strengths:** Flexible queries, single endpoint, no over/under-fetching, self-documenting
- **Weaknesses:** Complex caching, N+1 query problem, performance issues with deep nesting

## When to Use What?
|Scenario|	Recommended|
|--------|-------------|
|Public API, web/mobile clients|	REST (mature, browser-friendly)|
|Internal microservices, high performance|	gRPC (low latency, efficient)|
|IoT, real-time streaming|	gRPC (bidirectional streaming)|
|Complex data requirements, client flexibility|	GraphQL (avoid over-fetching)|
|Simple CRUD operations|	REST|
|Polyglot environments|	gRPC (multi-language support)|

# How microservices are communicate with each other using REST and gRPC? 

