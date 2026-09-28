# Go gRPC 시작하기

Protobuf 정의부터 Go gRPC 서버와 클라이언트까지, gRPC 공식 helloworld 예제를 따라 만든 최소 구성입니다.

블로그 글: [gRPC 서버, 클라이언트 만들기](https://songtomtom.github.io/blog/go-grpc-server-client)

## 구조

```
proto/v1/helloworld.proto        서비스와 메시지 정의
proto/v1/helloworld.pb.go        protoc 가 생성한 메시지 코드
proto/v1/helloworld_grpc.pb.go   protoc 가 생성한 서비스 코드
server/server.go                 Greeter 서비스 구현
client/client.go                 서버를 호출하는 클라이언트
```

## 실행

```bash
# 코드 생성 (proto 를 바꿨을 때만)
make proto

# 서버
go run server/server.go
# server listening at [::]:50051

# 다른 터미널에서 클라이언트
go run client/client.go
# Greeting: Hello world
```

`protoc` 와 Go 플러그인이 필요합니다.

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```
