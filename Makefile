COMPOSE_API = docker compose -f api/compose.yml

.PHONY: api-build api-up api-down api-shell api-logs

api-build:
	@echo ">>> Building API image…"
	$(COMPOSE_API) build

api-up:
	@echo ">>> Starting API container…"
	$(COMPOSE_API) up -d

api-down:
	@echo ">>> Stopping API container…"
	$(COMPOSE_API) down

api-shell:
	@echo ">>> Opening shell into API container…"
	$(COMPOSE_API) exec app bash

api-logs:
	@echo ">>> Tailing logs for API…"
	$(COMPOSE_API) logs -f

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