package ports

//go:generate oapi-codegen -generate types -o "openapi_types.gen.go" -package "ports" "../../../api/openapi/shorturl.yml"
//go:generate oapi-codegen -generate gin -o "openapi_server.gen.go" -package "ports" "../../../api/openapi/shorturl.yml"
