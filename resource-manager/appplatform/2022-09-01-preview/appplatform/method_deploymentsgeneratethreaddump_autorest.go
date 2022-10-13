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

type DeploymentsGenerateThreadDumpOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeploymentsGenerateThreadDump ...
func (c AppPlatformClient) DeploymentsGenerateThreadDump(ctx context.Context, id DeploymentId, input DiagnosticParameters) (result DeploymentsGenerateThreadDumpOperationResponse, err error) {
	req, err := c.preparerForDeploymentsGenerateThreadDump(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGenerateThreadDump", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeploymentsGenerateThreadDump(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGenerateThreadDump", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeploymentsGenerateThreadDumpThenPoll performs DeploymentsGenerateThreadDump then polls until it's completed
func (c AppPlatformClient) DeploymentsGenerateThreadDumpThenPoll(ctx context.Context, id DeploymentId, input DiagnosticParameters) error {
	result, err := c.DeploymentsGenerateThreadDump(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DeploymentsGenerateThreadDump: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeploymentsGenerateThreadDump: %+v", err)
	}

	return nil
}

// preparerForDeploymentsGenerateThreadDump prepares the DeploymentsGenerateThreadDump request.
func (c AppPlatformClient) preparerForDeploymentsGenerateThreadDump(ctx context.Context, id DeploymentId, input DiagnosticParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/generateThreadDump", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeploymentsGenerateThreadDump sends the DeploymentsGenerateThreadDump request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForDeploymentsGenerateThreadDump(ctx context.Context, req *http.Request) (future DeploymentsGenerateThreadDumpOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
