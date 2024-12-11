gen_start:
	swag init -g cmd/main.go; go run cmd/main.go

swagger:
	swag init -g cmd/main.go