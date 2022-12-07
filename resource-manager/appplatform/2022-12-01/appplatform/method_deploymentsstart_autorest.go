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

type DeploymentsStartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsStart ...
func (c AppPlatformClient) DeploymentsStart(ctx context.Context, id DeploymentId) (result DeploymentsStartOperationResponse, err error) {
	req, err := c.preparerForDeploymentsStart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsStart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsStartThenPoll performs DeploymentsStart then polls until it's completed
func (c AppPlatformClient) DeploymentsStartThenPoll(ctx context.Context, id DeploymentId) error {
	result, err := c.DeploymentsStart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeploymentsStart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsStart: %+v", err)
	}

	return nil
}

// preparerForDeploymentsStart prepares the DeploymentsStart request.
func (c AppPlatformClient) preparerForDeploymentsStart(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsStart sends the DeploymentsStart request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsStart(ctx context.Context, req *http.Request) (future DeploymentsStartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
