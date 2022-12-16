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

type SampleInputOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SampleInput ...
func (c SubscriptionsClient) SampleInput(ctx context.Context, id LocationId, input SampleInput) (result SampleInputOperationResponse, err error) {
	req, err := c.preparerForSampleInput(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "SampleInput", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSampleInput(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "SampleInput", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SampleInputThenPoll performs SampleInput then polls until it's completed
func (c SubscriptionsClient) SampleInputThenPoll(ctx context.Context, id LocationId, input SampleInput) error {
	result, err := c.SampleInput(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SampleInput: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SampleInput: %+v", err)
	}

	return nil
}

// preparerForSampleInput prepares the SampleInput request.
func (c SubscriptionsClient) preparerForSampleInput(ctx context.Context, id LocationId, input SampleInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sampleInput", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSampleInput sends the SampleInput request. The method will close the
// http.Response Body if it receives an error.
func (c SubscriptionsClient) senderForSampleInput(ctx context.Context, req *http.Request) (future SampleInputOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
