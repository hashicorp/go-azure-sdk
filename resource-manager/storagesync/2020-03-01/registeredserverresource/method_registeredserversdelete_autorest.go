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

type RegisteredServersDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegisteredServersDelete ...
func (c RegisteredServerResourceClient) RegisteredServersDelete(ctx context.Context, id RegisteredServerId) (result RegisteredServersDeleteOperationResponse, err error) {
	req, err := c.preparerForRegisteredServersDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegisteredServersDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegisteredServersDeleteThenPoll performs RegisteredServersDelete then polls until it's completed
func (c RegisteredServerResourceClient) RegisteredServersDeleteThenPoll(ctx context.Context, id RegisteredServerId) error {
	result, err := c.RegisteredServersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegisteredServersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegisteredServersDelete: %+v", err)
	}

	return nil
}

// preparerForRegisteredServersDelete prepares the RegisteredServersDelete request.
func (c RegisteredServerResourceClient) preparerForRegisteredServersDelete(ctx context.Context, id RegisteredServerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRegisteredServersDelete sends the RegisteredServersDelete request. The method will close the
// http.Response Body if it receives an error.
func (c RegisteredServerResourceClient) senderForRegisteredServersDelete(ctx context.Context, req *http.Request) (future RegisteredServersDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
