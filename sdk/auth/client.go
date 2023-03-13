package auth

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

// httpClient returns a shimmed retryablehttp Client, at the moment with default retry,
// backoff and timeout settings, but we may need to change this in the future.
func httpClient() *http.Client {
	retryableClient := retryablehttp.NewClient()
	return retryableClient.StandardClient()
}
