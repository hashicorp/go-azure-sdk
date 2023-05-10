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

type DeleteAtManagementGroupScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteAtManagementGroupScope ...
func (c DeploymentsClient) DeleteAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (result DeleteAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForDeleteAtManagementGroupScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteAtManagementGroupScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteAtManagementGroupScopeThenPoll performs DeleteAtManagementGroupScope then polls until it's completed
func (c DeploymentsClient) DeleteAtManagementGroupScopeThenPoll(ctx context.Context, id Providers2DeploymentId) error {
	result, err := c.DeleteAtManagementGroupScope(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteAtManagementGroupScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteAtManagementGroupScope: %+v", err)
	}

	return nil
}

// preparerForDeleteAtManagementGroupScope prepares the DeleteAtManagementGroupScope request.
func (c DeploymentsClient) preparerForDeleteAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId) (*http.Request, error) {
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

// senderForDeleteAtManagementGroupScope sends the DeleteAtManagementGroupScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForDeleteAtManagementGroupScope(ctx context.Context, req *http.Request) (future DeleteAtManagementGroupScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
