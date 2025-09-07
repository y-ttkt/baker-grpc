COMPOSE_API = docker compose -f api/compose.yml
COMPOSE_CLIENT = docker compose -f client/compose.yml

.PHONY: api-build api-up api-down api-shell api-logs gen-proto \
        client-build client-up client-down client-shell client-logs client-rail

# ── API ───────────────────────────────────────────────
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
	$(COMPOSE_API) exec app bash -lc '\
	  cd /app/api && \
	  protoc \
	    -I ../proto \
	    -I /usr/include \
	    --go_out=gen/api --go_opt=paths=source_relative \
	    --go-grpc_out=gen/api --go-grpc_opt=paths=source_relative \
	    ../proto/*.proto \
	'

# ── Rails Client ──────────────────────────────────────
client-build:
	@echo ">>> Building Rails image…"
	$(COMPOSE_CLIENT) build

client-up:
	@echo ">>> Starting Rails container…"
	$(COMPOSE_CLIENT) up -d

client-down:
	@echo ">>> Stopping Rails container…"
	$(COMPOSE_CLIENT) down

client-shell:
	@echo ">>> Opening shell into Rails container…"
	$(COMPOSE_CLIENT) exec web bash

client-logs:
	@echo ">>> Tailing logs for Rails…"
	$(COMPOSE_CLIENT) logs -f

client-rails:
	@echo ">>> Running Rails server (foreground)…"
	$(COMPOSE_CLIENT) run --rm --service-ports web bash -lc "bundle exec rails server -b 0.0.0.0 -p 3000"