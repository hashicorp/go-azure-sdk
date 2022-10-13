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

type DeploymentsStopOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsStop ...
func (c AppPlatformClient) DeploymentsStop(ctx context.Context, id DeploymentId) (result DeploymentsStopOperationResponse, err error) {
	req, err := c.preparerForDeploymentsStop(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStop", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsStop(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStop", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsStopThenPoll performs DeploymentsStop then polls until it's completed
func (c AppPlatformClient) DeploymentsStopThenPoll(ctx context.Context, id DeploymentId) error {
	result, err := c.DeploymentsStop(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeploymentsStop: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsStop: %+v", err)
	}

	return nil
}

// preparerForDeploymentsStop prepares the DeploymentsStop request.
func (c AppPlatformClient) preparerForDeploymentsStop(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsStop sends the DeploymentsStop request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsStop(ctx context.Context, req *http.Request) (future DeploymentsStopOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
