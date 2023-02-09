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

type RegisteredServersCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegisteredServersCreate ...
func (c RegisteredServerResourceClient) RegisteredServersCreate(ctx context.Context, id RegisteredServerId, input RegisteredServerCreateParameters) (result RegisteredServersCreateOperationResponse, err error) {
	req, err := c.preparerForRegisteredServersCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegisteredServersCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegisteredServersCreateThenPoll performs RegisteredServersCreate then polls until it's completed
func (c RegisteredServerResourceClient) RegisteredServersCreateThenPoll(ctx context.Context, id RegisteredServerId, input RegisteredServerCreateParameters) error {
	result, err := c.RegisteredServersCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegisteredServersCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegisteredServersCreate: %+v", err)
	}

	return nil
}

// preparerForRegisteredServersCreate prepares the RegisteredServersCreate request.
func (c RegisteredServerResourceClient) preparerForRegisteredServersCreate(ctx context.Context, id RegisteredServerId, input RegisteredServerCreateParameters) (*http.Request, error) {
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

// senderForRegisteredServersCreate sends the RegisteredServersCreate request. The method will close the
// http.Response Body if it receives an error.
func (c RegisteredServerResourceClient) senderForRegisteredServersCreate(ctx context.Context, req *http.Request) (future RegisteredServersCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
