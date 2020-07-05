SwaggerPath := ./api/swagger-spec/swagger.json
Output := prober

.PHONY: prober clean swagger serve-swagger

prober:
	GO111MODULE=on go build -mod vendor -o $(Output) cmd/main.go

clean:
	@rm -f $(Output)

swagger:
	GO111MODULE=on swagger generate spec -o $(SwaggerPath) --scan-models

serve-swagger:
	swagger serve -F=swagger $(SwaggerPath)