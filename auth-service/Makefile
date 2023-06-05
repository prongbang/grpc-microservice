gen:
	protoc --go_out=proto --go_opt=paths=import \
        --go-grpc_out=proto --go-grpc_opt=paths=import \
        proto/$(name).proto

gen_auth:
	make gen name=auth

gen_user:
	make gen name=user

run_server:
	go run .

test_login:
	grpcurl -plaintext -import-path ./proto -proto auth.proto -d '{"username": "admin", "password": "1234"}' '[::1]:50051' auth.Auth/Login
