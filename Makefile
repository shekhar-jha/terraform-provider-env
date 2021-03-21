TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=github.com
NAMESPACE=shekhar-jha
NAME=env
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

git-status:
	@status=$$(git status --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "Not all the changes have been committed. Please ensure that output of 'git status --porcelain' is empty"; \
		exit 1; \
	fi

releasePrereq:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is not specified with personal access token. Please specify the same to automatically push the release to github )
endif
ifndef GPG_FINGERPRINT
	$(error GPG_FINGERPRINT is not specified. Please run 'gpg --list-secret-keys --keyid-format LONG' to identify the fingerprint that should be used to sign the release)
endif

tag:
	@GIT_TAG_VALUE=$$(git describe --abbrev=0 --tags);\
	GIT_EXPECTED_TAG_VALUE="v${VERSION}" \
	echo "Current tag $${GIT_TAG_VALUE} Expected tag: $${GIT_EXPECTED_TAG_VALUE}" \
	if [ ! -z "$${GIT_TAG_VALUE}" && "$${GIT_TAG_VALUE}" != "$${GIT_EXPECTED_TAG_VALUE}"]; \
	then \
		git tag ${EXPECTED_TAG_VALUE}\
	fi

release: releasePrereq git-status tag
	goreleaser --rm-dist

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	rm ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
