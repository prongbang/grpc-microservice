gen:
	protoc --go_out=proto --go_opt=paths=import \
        --go-grpc_out=proto --go-grpc_opt=paths=import \
        proto/$(name).proto

gen_user:
	make gen name=user

run_server:
	go run .

test_login:
	grpcurl -plaintext -import-path ./proto -proto user.proto -d '{"username": "admin"}' '[::1]:50052' user.User/GetUser
