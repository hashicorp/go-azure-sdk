package blobservice

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetServicePropertiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *BlobServiceProperties
}

// GetServiceProperties ...
func (c BlobServiceClient) GetServiceProperties(ctx context.Context, id BlobServiceId) (result GetServicePropertiesOperationResponse, err error) {
	req, err := c.preparerForGetServiceProperties(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blobservice.BlobServiceClient", "GetServiceProperties", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "blobservice.BlobServiceClient", "GetServiceProperties", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetServiceProperties(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blobservice.BlobServiceClient", "GetServiceProperties", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetServiceProperties prepares the GetServiceProperties request.
func (c BlobServiceClient) preparerForGetServiceProperties(ctx context.Context, id BlobServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetServiceProperties handles the response to the GetServiceProperties request. The method always
// closes the http.Response Body.
func (c BlobServiceClient) responderForGetServiceProperties(resp *http.Response) (result GetServicePropertiesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
