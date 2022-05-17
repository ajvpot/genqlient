package generate

import normalizer "github.com/golang/protobuf/protoc-gen-go/generator"

// normalize normalizes operation and var names from gql schema
func normalize(s string) string {
	return normalizer.CamelCase(s)
}
