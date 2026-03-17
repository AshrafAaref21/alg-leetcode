.PHONY: test testv fmt new

test:
	go test ./...

testv:
	go test -v ./...

fmt:
	go fmt ./...

new:
	@if [ -z "$(ID)" ] || [ -z "$(SLUG)" ]; then \
		echo "Usage: make new ID=<number> SLUG=<problem-slug>"; \
		exit 1; \
	fi
	@./scripts/new_problem.sh $(ID) "$(SLUG)"
