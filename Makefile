SOURCE=./...
VERSION="0.0.2"

.DEFAULT_GOAL := build

vet:
	go vet $(SOURCE)
.PHONY: vet

fmt:
	go fmt $(SOURCE)
.PHONY: fmt

test-fmt:
	test -z $(shell go fmt $(SOURCE))
.PHONY: test-fmt

test: vet test-fmt
	go test -cover $(SOURCE) -count=1
.PHONY: test

tools:
	echo "Installing tools from tools.go"
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
.PHONY: tools

build: tools
	goreleaser release \
		--snapshot \
		--skip-publish \
		--rm-dist
.PHONY: build

release: tools tag
	goreleaser release \
		--rm-dist
.PHONY: release

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
				terraform plan -var 'greeting=goodbye' -out 'plan.out' && \
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
.PHONY: testdata

reset-testdata:
	git rm testdata/show-plan.json || exit 0
	git rm testdata/show-state.json || exit 0
	git rm testdata/terraform.tfstate || exit 0
.PHONY: reset-testdata

check-tag:
	./scripts/ensure_unique_version.sh "$(VERSION)"
.PHONY: check-tag

tag: check-tag
	@echo "creating git tag $(VERSION)"
	@git tag $(VERSION)
	@git push origin $(VERSION)

clean:
	rm -rf dist || exit 0
	rm -rf data || exit 0
	rm -rf testdata/.terraform* || exit 0
	rm -rf testdata/terraform.tfstate.backup || exit 0
	rm -rf testdata/greeting_*.sh || exit 0
	rm -rf testdata/show-plan-unformatted.json || exit 0
	rm -rf testdata/plan.out || exit 0
.PHONY: clean
