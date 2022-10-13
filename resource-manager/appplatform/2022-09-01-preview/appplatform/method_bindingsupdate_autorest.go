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

type BindingsUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BindingsUpdate ...
func (c AppPlatformClient) BindingsUpdate(ctx context.Context, id BindingId, input BindingResource) (result BindingsUpdateOperationResponse, err error) {
	req, err := c.preparerForBindingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBindingsUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BindingsUpdateThenPoll performs BindingsUpdate then polls until it's completed
func (c AppPlatformClient) BindingsUpdateThenPoll(ctx context.Context, id BindingId, input BindingResource) error {
	result, err := c.BindingsUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BindingsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BindingsUpdate: %+v", err)
	}

	return nil
}

// preparerForBindingsUpdate prepares the BindingsUpdate request.
func (c AppPlatformClient) preparerForBindingsUpdate(ctx context.Context, id BindingId, input BindingResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBindingsUpdate sends the BindingsUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBindingsUpdate(ctx context.Context, req *http.Request) (future BindingsUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
