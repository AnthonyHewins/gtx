gtx: ## Build main binary
	go build -o bin/gtx cmd/gtx/*.go

deploy: gtx ## Move to ~/.local/bin
	mv ./bin/gtx $(HOME)/.local/bin
