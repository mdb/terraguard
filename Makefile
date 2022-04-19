SOURCE=./...
VERSION="0.0.0"

.PHONY: vet \
	fmt \
	test-fmt \
	test \
	goreleaser \
	testdata \
	clean

.DEFAULT_GOAL := build

vet:
	go vet $(SOURCE)

fmt:
	go fmt $(SOURCE)

test-fmt:
	test -z $(shell go fmt $(SOURCE))

test: vet test-fmt
	go test -cover $(SOURCE) -count=1

goreleaser:
	if ! which goreleaser &> /dev/null; then \
		go get github.com/goreleaser/goreleaser; \
	fi

build: goreleaser
	goreleaser release \
		--snapshot \
		--skip-publish \
		--rm-dist

release: goreleaser tag
	goreleaser release \
		--rm-dist

define generate-testdata
	docker run \
		--interactive \
		--tty \
		--volume $(shell pwd):/src \
		--workdir /src/testdata \
		--entrypoint /bin/sh \
		hashicorp/terraform:$(1) \
			-c \
				"terraform init && \
				terraform apply -auto-approve && \
				terraform show -json terraform.tfstate > show-state.json && \
				terraform plan -var 'greeting_one=goodbye' -out 'plan.out' && \
				terraform show -json plan.out > show-plan-unformatted.json"
	docker run \
		--volume $(shell pwd):/src \
		--workdir /src/testdata \
		--entrypoint /bin/sh \
		stedolan/jq \
			-c \
				"cat show-plan-unformatted.json | jq . > show-plan.json"
endef

testdata:
	$(call generate-testdata,1.1.5)

reset-testdata:
	rm testdata/show-plan-unformatted.json || exit 0
	rm testdata/show-plan.json || exit 0
	rm testdata/show-state.json || exit 0
	rm testdata/terraform.tfstate || exit 0
	rm testdata/plan.out || exit 0

tag:
	if git rev-parse $(VERSION) >/dev/null 2>&1; then \
		echo "found existing $(VERSION) git tag"; \
	else \
		echo "creating git tag $(VERSION)"; \
		git tag $(VERSION); \
		git push origin $(VERSION); \
	fi

clean:
	rm -rf dist || exit 0
	rm -rf data || exit 0
	rm -rf testdata/.terraform* || exit 0
	rm -rf testdata/terraform.tfstate.backup || exit 0
	rm -rf testdata/greeting_*.sh || exit 0
	rm -rf testdata/show-plan-unformatted.json || exit 0
