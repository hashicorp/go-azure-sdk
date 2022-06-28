package deployments

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

type DeleteAtScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteAtScope ...
func (c DeploymentsClient) DeleteAtScope(ctx context.Context, id ScopedDeploymentId) (result DeleteAtScopeOperationResponse, err error) {
	req, err := c.preparerForDeleteAtScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteAtScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteAtScopeThenPoll performs DeleteAtScope then polls until it's completed
func (c DeploymentsClient) DeleteAtScopeThenPoll(ctx context.Context, id ScopedDeploymentId) error {
	result, err := c.DeleteAtScope(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteAtScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteAtScope: %+v", err)
	}

	return nil
}

// preparerForDeleteAtScope prepares the DeleteAtScope request.
func (c DeploymentsClient) preparerForDeleteAtScope(ctx context.Context, id ScopedDeploymentId) (*http.Request, error) {
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

// senderForDeleteAtScope sends the DeleteAtScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForDeleteAtScope(ctx context.Context, req *http.Request) (future DeleteAtScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
