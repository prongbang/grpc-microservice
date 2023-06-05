gen:
	protoc --go_out=proto --go_opt=paths=import \
        --go-grpc_out=proto --go-grpc_opt=paths=import \
        proto/$(name).proto

gen_auth:
	make gen name=auth

run_server:
	go run .

test_login:
	curl -X POST http://localhost:8000/v1/login -d '{"username": "admin", "password": "1234"}'