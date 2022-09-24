# RPC - GRPC

### RPC

>Remote Procedure Call  -> 过程远程调用

**RPC process: client - stub - server**

### gRPC one of RPC from Google - http2.0的protobuf ([IDL](https://en.wikipedia.org/wiki/IDL_(programming_language)))


4 service types of gRPC:
- Unary
- Client-side streaming
- Server-side streaming
- Bidirectional streaming (two-way)

---

#### All contents on below...
>describe a route service，and the service name is RouteGuide.
four serviceTypes described in service: Point, Rectangle, Feature, RouteSummary, Chat.
four serviceMethods provided in service:
>>1. GetFuture(input: Point) --> output: Feature
>>2. ListFuture(input: Rectangle) --> output stream: Features
>>3. RecordRoute(input stream: real-time continuous points) --> output: RouteSummary
>>4. Recommend(input stream: recommendationRequest) --> output stream: Features



