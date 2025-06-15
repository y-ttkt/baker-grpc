COMPOSE_API = docker compose -f api/compose.yml

.PHONY: api-build api-up api-down api-shell api-logs gen-proto

build:
	@echo ">>> Building API image…"
	$(COMPOSE_API) build

up:
	@echo ">>> Starting API container…"
	$(COMPOSE_API) up -d

down:
	@echo ">>> Stopping API container…"
	$(COMPOSE_API) down

shell:
	@echo ">>> Opening shell into API container…"
	$(COMPOSE_API) exec app bash

logs:
	@echo ">>> Tailing logs for API…"
	$(COMPOSE_API) logs -f

gen-proto:
	@echo ">>> Generating Go code from .proto files in container…"
	$(COMPOSE_API) exec app bash\
        protoc -I proto \
          --go_out=gen/api \
          --go_opt=paths=source_relative \
          --go-grpc_out=gen/api \
          --go-grpc_opt=paths=source_relative \
          proto/*.proto

# ── Rails クライアントの操作例 ─────────────────
# .PHONY: client-up client-down client-shell
# client-up:
#   @echo ">>> Starting Rails container…"
#   $(COMPOSE_CLIENT) up -d
#
# client-down:
#   @echo ">>> Stopping Rails container…"
#   $(COMPOSE_CLIENT) down
#
# client-shell:
#   @echo ">>> Shell into Rails container…"
#   $(COMPOSE_CLIENT) exec web sh