package appplatform

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

type DeploymentsRestartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsRestart ...
func (c AppPlatformClient) DeploymentsRestart(ctx context.Context, id DeploymentId) (result DeploymentsRestartOperationResponse, err error) {
	req, err := c.preparerForDeploymentsRestart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsRestart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsRestart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsRestart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsRestartThenPoll performs DeploymentsRestart then polls until it's completed
func (c AppPlatformClient) DeploymentsRestartThenPoll(ctx context.Context, id DeploymentId) error {
	result, err := c.DeploymentsRestart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeploymentsRestart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsRestart: %+v", err)
	}

	return nil
}

// preparerForDeploymentsRestart prepares the DeploymentsRestart request.
func (c AppPlatformClient) preparerForDeploymentsRestart(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsRestart sends the DeploymentsRestart request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsRestart(ctx context.Context, req *http.Request) (future DeploymentsRestartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
