package key

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type SpatialAnchorsAccountsListKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccountKeys
}

// SpatialAnchorsAccountsListKeys ...
func (c KeyClient) SpatialAnchorsAccountsListKeys(ctx context.Context, id SpatialAnchorsAccountId) (result SpatialAnchorsAccountsListKeysOperationResponse, err error) {
	req, err := c.preparerForSpatialAnchorsAccountsListKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsListKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsListKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSpatialAnchorsAccountsListKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "SpatialAnchorsAccountsListKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSpatialAnchorsAccountsListKeys prepares the SpatialAnchorsAccountsListKeys request.
func (c KeyClient) preparerForSpatialAnchorsAccountsListKeys(ctx context.Context, id SpatialAnchorsAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSpatialAnchorsAccountsListKeys handles the response to the SpatialAnchorsAccountsListKeys request. The method always
// closes the http.Response Body.
func (c KeyClient) responderForSpatialAnchorsAccountsListKeys(resp *http.Response) (result SpatialAnchorsAccountsListKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
