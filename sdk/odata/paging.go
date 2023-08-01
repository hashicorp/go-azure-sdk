package odata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pager interface {
	NextPageLink() *Link
}

func NextLinkFromPager(resp *http.Response, pager Pager) (*Link, error) {
	// Read the response body and close it
	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	// Always reassign the response body
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}

	// Unmarshal
	if err := json.Unmarshal(respBody, &pager); err != nil {
		return nil, err
	}

	return pager.NextPageLink(), nil
}
