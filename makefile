server:
	go run main.go

docs:
	swag init --parseDependency --parseInternal --parseDepth 1 -g routes/routes.go

.PHONY: server 	docs