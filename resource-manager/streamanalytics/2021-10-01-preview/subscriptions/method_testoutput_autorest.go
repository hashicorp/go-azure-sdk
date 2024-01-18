package subscriptions

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

type TestOutputOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *TestDatasourceResult
}

// TestOutput ...
func (c SubscriptionsClient) TestOutput(ctx context.Context, id LocationId, input TestOutput) (result TestOutputOperationResponse, err error) {
	req, err := c.preparerForTestOutput(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "TestOutput", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTestOutput(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "TestOutput", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TestOutputThenPoll performs TestOutput then polls until it's completed
func (c SubscriptionsClient) TestOutputThenPoll(ctx context.Context, id LocationId, input TestOutput) error {
	result, err := c.TestOutput(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing TestOutput: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after TestOutput: %+v", err)
	}

	return nil
}

// preparerForTestOutput prepares the TestOutput request.
func (c SubscriptionsClient) preparerForTestOutput(ctx context.Context, id LocationId, input TestOutput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/testOutput", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTestOutput sends the TestOutput request. The method will close the
// http.Response Body if it receives an error.
func (c SubscriptionsClient) senderForTestOutput(ctx context.Context, req *http.Request) (future TestOutputOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
