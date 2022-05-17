package generate

import (
	"fmt"

	normalizer "github.com/golang/protobuf/protoc-gen-go/generator"
)

// normalize normalizes operation and var names from gql schema
func normalize(s string) string {
	fmt.Println(s, normalizer.CamelCase(s))
	return normalizer.CamelCase(s)
}
