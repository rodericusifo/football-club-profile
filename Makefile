start-dev:
	@echo "  >  Starting Program..."
	wire wire/wire.go
	go run cmd/main.go -env=dev

start-stag:
	@echo "  >  Starting Program..."
	wire wire/wire.go
	go run cmd/main.go -env=stag

start-prod:
	@echo "  >  Starting Program..."
	wire wire/wire.go
	go run cmd/main.go -env=prod

start-test:
	@echo "  >  Starting Program..."
	wire wire/wire.go
	go run cmd/main.go -env=test