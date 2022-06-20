package v2workspaceconnectionresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type WorkspaceConnectionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceConnectionPropertiesV2BasicResource
}

// WorkspaceConnectionsGet ...
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsGet(ctx context.Context, id ConnectionId) (result WorkspaceConnectionsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceConnectionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceConnectionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceConnectionsGet prepares the WorkspaceConnectionsGet request.
func (c V2WorkspaceConnectionResourceClient) preparerForWorkspaceConnectionsGet(ctx context.Context, id ConnectionId) (*http.Request, error) {
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

// responderForWorkspaceConnectionsGet handles the response to the WorkspaceConnectionsGet request. The method always
// closes the http.Response Body.
func (c V2WorkspaceConnectionResourceClient) responderForWorkspaceConnectionsGet(resp *http.Response) (result WorkspaceConnectionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
