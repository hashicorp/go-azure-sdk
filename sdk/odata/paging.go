package odata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Pager handles custom paging for paginated API responses that do not follow the OData 4.0 standard for JSON services.
// The underlying type should support unmarshalling a JSON response
type Pager interface {
	// NextPageLink returns a *Link describing the URI for the next page of results, it should also clear any state
	// before returning, so that subsequent pages do not inherit the URI from the previous page.
	NextPageLink() *Link
}

// NextLinkFromPager unmarshalls a *http.Response into the provided Pager and invokes its NextPageLink method
func NextLinkFromPager(resp *http.Response, pager Pager) (*Link, error) {
	if pager == nil {
		return nil, fmt.Errorf("internal-error: pager was nil, should be a pointer")
	}

	// Read the response body and close it
	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	// Always reassign the response body
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}

	// Unmarshal
	if err := json.Unmarshal(respBody, pager); err != nil {
		return nil, err
	}

	return pager.NextPageLink(), nil
}
