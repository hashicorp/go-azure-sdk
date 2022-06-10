package key

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type RemoteRenderingAccountsListKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccountKeys
}

// RemoteRenderingAccountsListKeys ...
func (c KeyClient) RemoteRenderingAccountsListKeys(ctx context.Context, id RemoteRenderingAccountId) (result RemoteRenderingAccountsListKeysOperationResponse, err error) {
	req, err := c.preparerForRemoteRenderingAccountsListKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsListKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsListKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRemoteRenderingAccountsListKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "key.KeyClient", "RemoteRenderingAccountsListKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRemoteRenderingAccountsListKeys prepares the RemoteRenderingAccountsListKeys request.
func (c KeyClient) preparerForRemoteRenderingAccountsListKeys(ctx context.Context, id RemoteRenderingAccountId) (*http.Request, error) {
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

// responderForRemoteRenderingAccountsListKeys handles the response to the RemoteRenderingAccountsListKeys request. The method always
// closes the http.Response Body.
func (c KeyClient) responderForRemoteRenderingAccountsListKeys(resp *http.Response) (result RemoteRenderingAccountsListKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
