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

type DeploymentsStartJFROperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsStartJFR ...
func (c AppPlatformClient) DeploymentsStartJFR(ctx context.Context, id DeploymentId, input DiagnosticParameters) (result DeploymentsStartJFROperationResponse, err error) {
	req, err := c.preparerForDeploymentsStartJFR(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStartJFR", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsStartJFR(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsStartJFR", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsStartJFRThenPoll performs DeploymentsStartJFR then polls until it's completed
func (c AppPlatformClient) DeploymentsStartJFRThenPoll(ctx context.Context, id DeploymentId, input DiagnosticParameters) error {
	result, err := c.DeploymentsStartJFR(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DeploymentsStartJFR: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsStartJFR: %+v", err)
	}

	return nil
}

// preparerForDeploymentsStartJFR prepares the DeploymentsStartJFR request.
func (c AppPlatformClient) preparerForDeploymentsStartJFR(ctx context.Context, id DeploymentId, input DiagnosticParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/startJFR", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsStartJFR sends the DeploymentsStartJFR request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsStartJFR(ctx context.Context, req *http.Request) (future DeploymentsStartJFROperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
