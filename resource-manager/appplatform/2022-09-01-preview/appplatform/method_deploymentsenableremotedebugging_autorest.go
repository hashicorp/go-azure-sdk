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

type DeploymentsEnableRemoteDebuggingOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsEnableRemoteDebugging ...
func (c AppPlatformClient) DeploymentsEnableRemoteDebugging(ctx context.Context, id DeploymentId, input RemoteDebuggingPayload) (result DeploymentsEnableRemoteDebuggingOperationResponse, err error) {
	req, err := c.preparerForDeploymentsEnableRemoteDebugging(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsEnableRemoteDebugging", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsEnableRemoteDebugging(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsEnableRemoteDebugging", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsEnableRemoteDebuggingThenPoll performs DeploymentsEnableRemoteDebugging then polls until it's completed
func (c AppPlatformClient) DeploymentsEnableRemoteDebuggingThenPoll(ctx context.Context, id DeploymentId, input RemoteDebuggingPayload) error {
	result, err := c.DeploymentsEnableRemoteDebugging(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DeploymentsEnableRemoteDebugging: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsEnableRemoteDebugging: %+v", err)
	}

	return nil
}

// preparerForDeploymentsEnableRemoteDebugging prepares the DeploymentsEnableRemoteDebugging request.
func (c AppPlatformClient) preparerForDeploymentsEnableRemoteDebugging(ctx context.Context, id DeploymentId, input RemoteDebuggingPayload) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/enableRemoteDebugging", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsEnableRemoteDebugging sends the DeploymentsEnableRemoteDebugging request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsEnableRemoteDebugging(ctx context.Context, req *http.Request) (future DeploymentsEnableRemoteDebuggingOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
