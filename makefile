# install on macos
# brew install openapi-generator

gen-go:
	openapi-generator generate -i api/openapi.yaml -g go -o .
