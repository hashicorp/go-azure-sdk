package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePrivateEndpointConnectionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeletePrivateEndpointConnection ...
func (c StaticSitesClient) DeletePrivateEndpointConnection(ctx context.Context, id StaticSitePrivateEndpointConnectionId) (result DeletePrivateEndpointConnectionOperationResponse, err error) {
	req, err := c.preparerForDeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeletePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeletePrivateEndpointConnection(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeletePrivateEndpointConnection", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeletePrivateEndpointConnectionThenPoll performs DeletePrivateEndpointConnection then polls until it's completed
func (c StaticSitesClient) DeletePrivateEndpointConnectionThenPoll(ctx context.Context, id StaticSitePrivateEndpointConnectionId) error {
	result, err := c.DeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeletePrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeletePrivateEndpointConnection: %+v", err)
	}

	return nil
}

// preparerForDeletePrivateEndpointConnection prepares the DeletePrivateEndpointConnection request.
func (c StaticSitesClient) preparerForDeletePrivateEndpointConnection(ctx context.Context, id StaticSitePrivateEndpointConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeletePrivateEndpointConnection sends the DeletePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForDeletePrivateEndpointConnection(ctx context.Context, req *http.Request) (future DeletePrivateEndpointConnectionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
