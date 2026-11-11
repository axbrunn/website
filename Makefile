# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/web: run the cmd/web application
.PHONY: run/web
run/web:
	@go run ./cmd/web --http-port=8080 --log-level="debug"