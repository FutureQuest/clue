// Code generated by goa v3.5.4, DO NOT EDIT.
//
// HTTP request path constructors for the test service.
//
// Command:
// $ goa gen github.com/goadesign/clue/internal/testsvc/design

package server

import (
	"fmt"
)

// HTTPMethodTestPath returns the URL path to the test service http_method HTTP endpoint.
func HTTPMethodTestPath(i int) string {
	return fmt.Sprintf("/%v", i)
}
