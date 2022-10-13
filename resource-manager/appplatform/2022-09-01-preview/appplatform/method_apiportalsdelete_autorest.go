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

type ApiPortalsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApiPortalsDelete ...
func (c AppPlatformClient) ApiPortalsDelete(ctx context.Context, id ApiPortalId) (result ApiPortalsDeleteOperationResponse, err error) {
	req, err := c.preparerForApiPortalsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApiPortalsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApiPortalsDeleteThenPoll performs ApiPortalsDelete then polls until it's completed
func (c AppPlatformClient) ApiPortalsDeleteThenPoll(ctx context.Context, id ApiPortalId) error {
	result, err := c.ApiPortalsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ApiPortalsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApiPortalsDelete: %+v", err)
	}

	return nil
}

// preparerForApiPortalsDelete prepares the ApiPortalsDelete request.
func (c AppPlatformClient) preparerForApiPortalsDelete(ctx context.Context, id ApiPortalId) (*http.Request, error) {
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

// senderForApiPortalsDelete sends the ApiPortalsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForApiPortalsDelete(ctx context.Context, req *http.Request) (future ApiPortalsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
