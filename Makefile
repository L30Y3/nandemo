# Makefile for Nandemo integration testing and service orchestration

.PHONY: up down test logs clean

up:
	docker-compose up --build -d


down:
	docker-compose down


test:
	cd integration-test && go test -v

logs:
	docker-compose logs -f gateway-service

clean:
	docker system prune -f

wait:
	@echo "Waiting for gateway-service to become healthy..."
	@until curl -s http://localhost:8080/health | grep -q "Gateway Service OK"; do \
	  echo "Waiting for gateway..."; \
	  sleep 1; \
	done
	@echo "Gateway is up!"