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

type BindingsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BindingsCreateOrUpdate ...
func (c AppPlatformClient) BindingsCreateOrUpdate(ctx context.Context, id BindingId, input BindingResource) (result BindingsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForBindingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBindingsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BindingsCreateOrUpdateThenPoll performs BindingsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) BindingsCreateOrUpdateThenPoll(ctx context.Context, id BindingId, input BindingResource) error {
	result, err := c.BindingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BindingsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BindingsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForBindingsCreateOrUpdate prepares the BindingsCreateOrUpdate request.
func (c AppPlatformClient) preparerForBindingsCreateOrUpdate(ctx context.Context, id BindingId, input BindingResource) (*http.Request, error) {
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

// senderForBindingsCreateOrUpdate sends the BindingsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBindingsCreateOrUpdate(ctx context.Context, req *http.Request) (future BindingsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
