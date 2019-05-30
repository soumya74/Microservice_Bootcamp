COMMIT = ${shell git rev-parse HEAD}
BRANCH = ${shell git rev-parse --abbrev-ref HEAD}
BUILDPATH = $(CURDIR)

build_linux :
	@echo "Started Building for linux environment ..."
	@echo "CommitId : "${COMMIT}
	@echo "Branch : "${BRANCH}
	@echo "BuildPath :"${BUILDPATH}
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	@echo "Build ended successfully"

build :
	@echo "Started Building ..."
	@echo "CommitId : "${COMMIT}
	@echo "Branch : "${BRANCH}
	@echo "BuildPath :"${BUILDPATH}
	go build -o app .
	@echo "Build ended successfully"

test :
	go test
