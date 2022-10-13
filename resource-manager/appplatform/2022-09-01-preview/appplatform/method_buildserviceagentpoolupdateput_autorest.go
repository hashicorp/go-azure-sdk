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

type BuildServiceAgentPoolUpdatePutOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BuildServiceAgentPoolUpdatePut ...
func (c AppPlatformClient) BuildServiceAgentPoolUpdatePut(ctx context.Context, id AgentPoolId, input BuildServiceAgentPoolResource) (result BuildServiceAgentPoolUpdatePutOperationResponse, err error) {
	req, err := c.preparerForBuildServiceAgentPoolUpdatePut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolUpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBuildServiceAgentPoolUpdatePut(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolUpdatePut", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BuildServiceAgentPoolUpdatePutThenPoll performs BuildServiceAgentPoolUpdatePut then polls until it's completed
func (c AppPlatformClient) BuildServiceAgentPoolUpdatePutThenPoll(ctx context.Context, id AgentPoolId, input BuildServiceAgentPoolResource) error {
	result, err := c.BuildServiceAgentPoolUpdatePut(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BuildServiceAgentPoolUpdatePut: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BuildServiceAgentPoolUpdatePut: %+v", err)
	}

	return nil
}

// preparerForBuildServiceAgentPoolUpdatePut prepares the BuildServiceAgentPoolUpdatePut request.
func (c AppPlatformClient) preparerForBuildServiceAgentPoolUpdatePut(ctx context.Context, id AgentPoolId, input BuildServiceAgentPoolResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBuildServiceAgentPoolUpdatePut sends the BuildServiceAgentPoolUpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBuildServiceAgentPoolUpdatePut(ctx context.Context, req *http.Request) (future BuildServiceAgentPoolUpdatePutOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
