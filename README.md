# grpc-microservice
```shell
+-------------+       +--------------+       +--------------+
| Web Service | ----> | Auth Service | ----> | User Service |
+-------------+       +--------------+       +--------------+
```

## Install

- Protobuf

```shell
brew install protobuf
```

- protoc-gen-go

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

- protoc-gen-grpc

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- grpcurl

```shell
brew install grpcurl
```


## Run Service

- [user-service](user-service)

```shell
cd user-service
make run_server
```

Test

```shell
cd user-service
grpcurl -plaintext -import-path ./proto -proto user.proto -d '{"username": "admin"}' '[::1]:50052' user.User/GetUser
```

Output

```json
{
  "id": "1",
  "name": "Administrator",
  "username": "admin",
  "password": "1234"
}
```

- [auth-service](auth-service)

```shell
cd auth-service
make run_server
```

Test

```shell
cd auth-service
grpcurl -plaintext -import-path ./proto -proto auth.proto -d '{"username": "admin", "password": "1234"}' '[::1]:50051' auth.Auth/Login
```

Output

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlbSIsIm5hbWUiOiJkZXYgZGF5IiwiaWF0IjoxNTE2MjM5MDIyfQ.yNC-7RUVZCveMOANZcT7KWMczVkb_T7KnHv3fmMLiCI"
}
```

- [web-service](web-service)

```shell
cd web-service
make run_server
```

Test

```shell
curl -X POST http://localhost:8000/v1/login -d '{"username": "admin", "password": "1234"}'
```

Output

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlbSIsIm5hbWUiOiJkZXYgZGF5IiwiaWF0IjoxNTE2MjM5MDIyfQ.yNC-7RUVZCveMOANZcT7KWMczVkb_T7KnHv3fmMLiCI"
}
```