package registeredserverresource

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

type RegisteredServerstriggerRolloverOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegisteredServerstriggerRollover ...
func (c RegisteredServerResourceClient) RegisteredServerstriggerRollover(ctx context.Context, id RegisteredServerId, input TriggerRolloverRequest) (result RegisteredServerstriggerRolloverOperationResponse, err error) {
	req, err := c.preparerForRegisteredServerstriggerRollover(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServerstriggerRollover", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegisteredServerstriggerRollover(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServerstriggerRollover", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegisteredServerstriggerRolloverThenPoll performs RegisteredServerstriggerRollover then polls until it's completed
func (c RegisteredServerResourceClient) RegisteredServerstriggerRolloverThenPoll(ctx context.Context, id RegisteredServerId, input TriggerRolloverRequest) error {
	result, err := c.RegisteredServerstriggerRollover(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegisteredServerstriggerRollover: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegisteredServerstriggerRollover: %+v", err)
	}

	return nil
}

// preparerForRegisteredServerstriggerRollover prepares the RegisteredServerstriggerRollover request.
func (c RegisteredServerResourceClient) preparerForRegisteredServerstriggerRollover(ctx context.Context, id RegisteredServerId, input TriggerRolloverRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/triggerRollover", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRegisteredServerstriggerRollover sends the RegisteredServerstriggerRollover request. The method will close the
// http.Response Body if it receives an error.
func (c RegisteredServerResourceClient) senderForRegisteredServerstriggerRollover(ctx context.Context, req *http.Request) (future RegisteredServerstriggerRolloverOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
