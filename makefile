DEV_TAG=dev
STG_TAG=stage
PROD_TAG=v0.0.1
IMAGE_NAME=gl-shopping


build:
	go build ./...

test:
	go test ./...  -v -count=1

