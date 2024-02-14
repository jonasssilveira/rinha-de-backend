format:
	goimports -w .

lint:
	docker run -t --rm -v "$(shell pwd)":/app -w /app golangci/golangci-lint golangci-lint run -v

test:
	go test -v ./...

test-coverage:
	go test -coverprofile coverage.out ./...
	@go tool cover -func=coverage.out | awk '/^total:/{ total += $$3; count++ } END { avg = total / count; if (count > 0) { printf "\033[1;33mAverage coverage: \033[0m"; if (avg > 95) printf "\033[1;32m%.2f%%\033[0m\n", avg; else printf "\033[1;31m%.2f%%\033[0m\n", avg; } else printf "\033[1;31mNo coverage information found\033[0m\n" }'
	go tool cover -html=coverage.out

vet:
	go vet ./...

staticcheck:
	 docker run -v "$(shell pwd)":/go/src/app -w /go/src/app -ti devdrops/staticcheck:latest staticcheck ./...

revive:
	@docker run -v "$(shell pwd)":/var/revive ghcr.io/mgechev/revive -config /var/revive/revive.toml -formatter friendly ./...

remove-coverage-files:
	rm cover_report.out coverage.out

sqlc-generate:
	docker run --rm -v "$(shell pwd)/db:/src" -w /src sqlc/sqlc generate

all-checks:
	-$(MAKE) format
	-$(MAKE) lint
	-$(MAKE) test-coverage
	-$(MAKE) vet
	-$(MAKE) staticcheck
	-$(MAKE) revive
	-$(MAKE) remove-coverage-files