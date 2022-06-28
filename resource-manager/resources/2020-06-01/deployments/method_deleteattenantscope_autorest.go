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

type DeleteAtTenantScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteAtTenantScope ...
func (c DeploymentsClient) DeleteAtTenantScope(ctx context.Context, id DeploymentId) (result DeleteAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForDeleteAtTenantScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtTenantScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteAtTenantScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteAtTenantScopeThenPoll performs DeleteAtTenantScope then polls until it's completed
func (c DeploymentsClient) DeleteAtTenantScopeThenPoll(ctx context.Context, id DeploymentId) error {
	result, err := c.DeleteAtTenantScope(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteAtTenantScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteAtTenantScope: %+v", err)
	}

	return nil
}

// preparerForDeleteAtTenantScope prepares the DeleteAtTenantScope request.
func (c DeploymentsClient) preparerForDeleteAtTenantScope(ctx context.Context, id DeploymentId) (*http.Request, error) {
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

// senderForDeleteAtTenantScope sends the DeleteAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForDeleteAtTenantScope(ctx context.Context, req *http.Request) (future DeleteAtTenantScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
