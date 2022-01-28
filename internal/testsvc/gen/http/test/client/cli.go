// Code generated by goa v3.5.4, DO NOT EDIT.
//
// test HTTP client CLI support package
//
// Command:
// $ goa gen github.com/goadesign/clue/internal/testsvc/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	test "github.com/goadesign/clue/internal/testsvc/gen/test"
)

// BuildHTTPMethodPayload builds the payload for the test http_method endpoint
// from CLI flags.
func BuildHTTPMethodPayload(testHTTPMethodBody string, testHTTPMethodI string) (*test.Fields, error) {
	var err error
	var body HTTPMethodRequestBody
	{
		err = json.Unmarshal([]byte(testHTTPMethodBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"s\": \"Voluptas exercitationem vitae.\"\n   }'")
		}
	}
	var i int
	{
		var v int64
		v, err = strconv.ParseInt(testHTTPMethodI, 10, 64)
		i = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for i, must be INT")
		}
	}
	v := &test.Fields{
		S: body.S,
	}
	v.I = &i

	return v, nil
}
