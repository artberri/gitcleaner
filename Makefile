GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test vet

cover-ci:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@sh -c "'$(CURDIR)/scripts/coverage.sh'"
	rm profile_tmp.cov

cover:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@sh -c "'$(CURDIR)/scripts/coverage.sh'"
	go tool cover -html=profile.cov
	rm profile_tmp.cov

test: fmtcheck
	go test -i ./... || exit 1
	go list ./... | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo 'go vet ./...'
	@go vet ./... ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

vendor-sync:
	@govendor sync

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: test vendor-sync
