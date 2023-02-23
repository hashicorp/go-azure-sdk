package odata_test

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type bodyOnce struct {
	io.Reader
	closed bool
}

func (b *bodyOnce) Read(p []byte) (n int, err error) {
	if b.closed {
		return 0, fmt.Errorf("closed")
	}
	return b.Reader.Read(p)
}

func (b *bodyOnce) Close() error {
	b.closed = true
	return nil
}

func body(in string) io.ReadCloser {
	return &bodyOnce{strings.NewReader(in), false}
}

func TestFromResponse(t *testing.T) {
	type testCase struct {
		response    *http.Response
		expected    *odata.OData
		shouldError bool
	}

	testCases := []testCase{
		{
			response:    nil,
			expected:    nil,
			shouldError: false,
		},
		{
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
				Body: body(`{
				"@odata.context": "https://graph.microsoft.com/beta/$metadata#servicePrincipals",
				"@odata.nextLink": "https://graph.microsoft.com/beta/1564a4be-0377-4d9b-8aff-5a2b564e177c/servicePrincipals?$skiptoken=X%274453707402000100000035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D34336331643937353963313035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D3433633164393735396331300000000000000000000000%27",
				"value": [
					{
						"id": "00000000-0000-0000-0000-000000000000",
						"deletedDateTime": null,
						"accountEnabled": true,
						"createdDateTime": "2020-07-08T01:22:29Z"
					}
				]
			}`),
			},
			expected: &odata.OData{
				Context:  pointer.To("https://graph.microsoft.com/beta/$metadata#servicePrincipals"),
				NextLink: pointer.To(odata.Link("https://graph.microsoft.com/beta/1564a4be-0377-4d9b-8aff-5a2b564e177c/servicePrincipals?%24skiptoken=X%274453707402000100000035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D34336331643937353963313035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D3433633164393735396331300000000000000000000000%27")),
				Value: []interface{}{map[string]interface{}{
					"id":              "00000000-0000-0000-0000-000000000000",
					"deletedDateTime": nil,
					"accountEnabled":  true,
					"createdDateTime": "2020-07-08T01:22:29Z",
				}},
			},
			shouldError: false,
		},
		{
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
				Body: body(`{
					"@odata.context": "https://graph.microsoft.com/beta/$metadata#users/$entity",
					"@odata.type": "#microsoft.graph.user",
					"@odata.id": "users('8a3b99a7-c82f-44b2-a10d-eb85b6c6b0e4')",
					"@odata.editLink": "users('8a3b99a7-c82f-44b2-a10d-eb85b6c6b0e4')",
					"id": "8a3b99a7-c82f-44b2-a10d-eb85b6c6b0e4",
					"deletedDateTime": null,
					"accountEnabled": true,
					"ageGroup": null,
					"businessPhones@odata.type": "#Collection(String)",
					"businessPhones": [],
					"city": null
				}`),
			},
			expected: &odata.OData{
				Context:  pointer.To("https://graph.microsoft.com/beta/$metadata#users/$entity"),
				Type:     pointer.To(odata.TypeUser),
				Id:       pointer.To(odata.Id("users('8a3b99a7-c82f-44b2-a10d-eb85b6c6b0e4')")),
				EditLink: pointer.To(odata.Link("")),
				Value:    nil,
			},
			shouldError: false,
		},
		{
			response: &http.Response{
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
				Body: body(`{G1bb3r1$h, "Non"5en5e}}}`),
			},
			expected:    nil,
			shouldError: true,
		},
	}

	for n, c := range testCases {
		o, err := odata.FromResponse(c.response)
		if c.shouldError {
			if err == nil {
				t.Errorf("test case %d: expected error but received none", n)
				continue
			}
		} else {
			if err != nil {
				t.Errorf("test case %d: expected no error but received: %+v", n, err)
				continue
			}
		}

		if c.expected == nil && o != nil {
			t.Errorf("test case %d: expected nil odata but received: %#v", n, *o)
		}

		if !reflect.DeepEqual(c.expected, o) {
			if o == nil {
				t.Errorf("test case %d: expected:\n    %s\nreceived:\n    nil", n, spew.Sdump(c.expected))
			} else {
				t.Errorf("test case %d: expected:\n    %s\nreceived:\n    %s", n, spew.Sdump(c.expected), spew.Sdump(o))
			}
		}

		if c.response != nil {
			if _, err := io.ReadAll(c.response.Body); err != nil {
				t.Errorf("test case %d: should have been able to read body but received error: %+v", n, err)
			}
		}
	}
}
