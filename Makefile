BUILD_NAME = url-shorting-service 

run:
	@echo 'Binary build started'
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BUILD_NAME) cmd/main.go
	@echo '[+] Build done'
	@echo 'Start docker-compose ...'
	@docker-compose up -d --build || (echo "[!] Docker-compose failed $$?"; exit 1)
	@echo '[+] Service started'

stop:
	@docker-compose stop

clean:
	@rm -rf $(BUILD_NAME)

.PHONY: run clean