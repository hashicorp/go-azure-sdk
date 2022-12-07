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

type BindingsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BindingsDelete ...
func (c AppPlatformClient) BindingsDelete(ctx context.Context, id BindingId) (result BindingsDeleteOperationResponse, err error) {
	req, err := c.preparerForBindingsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBindingsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BindingsDeleteThenPoll performs BindingsDelete then polls until it's completed
func (c AppPlatformClient) BindingsDeleteThenPoll(ctx context.Context, id BindingId) error {
	result, err := c.BindingsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing BindingsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BindingsDelete: %+v", err)
	}

	return nil
}

// preparerForBindingsDelete prepares the BindingsDelete request.
func (c AppPlatformClient) preparerForBindingsDelete(ctx context.Context, id BindingId) (*http.Request, error) {
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

// senderForBindingsDelete sends the BindingsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBindingsDelete(ctx context.Context, req *http.Request) (future BindingsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
