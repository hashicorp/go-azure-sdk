package batchendpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ListKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *EndpointAuthKeys
}

// ListKeys ...
func (c BatchEndpointClient) ListKeys(ctx context.Context, id BatchEndpointId) (result ListKeysOperationResponse, err error) {
	req, err := c.preparerForListKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchendpoint.BatchEndpointClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchendpoint.BatchEndpointClient", "ListKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchendpoint.BatchEndpointClient", "ListKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListKeys prepares the ListKeys request.
func (c BatchEndpointClient) preparerForListKeys(ctx context.Context, id BatchEndpointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListKeys handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (c BatchEndpointClient) responderForListKeys(resp *http.Response) (result ListKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
