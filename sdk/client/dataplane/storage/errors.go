// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

var _ error = &storageError{}

type storageError struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`

	// Details is whatever else came with the error, such as AuthenticationErrorDetail, which is often the only part that says what was wrong.
	Details []storageErrorDetail `xml:",any"`
}

type storageErrorDetail struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
	Nested  string `xml:",innerxml"`
}

func (e *storageError) Error() string {
	out := e.Message
	if e.Code != "" {
		out = fmt.Sprintf("%s: %s", e.Code, out)
	}

	for _, d := range e.Details {
		v := strings.TrimSpace(d.Value)
		if v == "" {
			v = strings.TrimSpace(d.Nested)
		}
		if v == "" {
			continue
		}

		out = fmt.Sprintf("%s\n  %s: %s", out, d.XMLName.Local, v)
	}

	return out
}

var _ client.ResponseErrorParser = &ErrorParser{}

// ErrorParser is a custom error parse for Azure Storage errors
type ErrorParser struct{}

// FromResponse returns the Storage error in the response, or nil when there is none to parse, so that the caller reports the response itself. Not
// every failure has one: a HEAD request gets no body, and Data Lake and Tables answer in JSON.
func (e *ErrorParser) FromResponse(resp *http.Response) error {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not parse response body: %+v", err)
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))
	respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

	res := storageError{}
	if err = xml.Unmarshal(respBody, &res); err != nil || (res.Code == "" && res.Message == "") {
		return nil
	}

	return &res
}
