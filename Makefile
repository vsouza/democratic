PROJECT=Democratic

run:
	@go run cmd/$(PROJECT)/main.go

test:
	@go test