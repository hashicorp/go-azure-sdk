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

type AppsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AppsDelete ...
func (c AppPlatformClient) AppsDelete(ctx context.Context, id AppId) (result AppsDeleteOperationResponse, err error) {
	req, err := c.preparerForAppsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAppsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AppsDeleteThenPoll performs AppsDelete then polls until it's completed
func (c AppPlatformClient) AppsDeleteThenPoll(ctx context.Context, id AppId) error {
	result, err := c.AppsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing AppsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AppsDelete: %+v", err)
	}

	return nil
}

// preparerForAppsDelete prepares the AppsDelete request.
func (c AppPlatformClient) preparerForAppsDelete(ctx context.Context, id AppId) (*http.Request, error) {
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

// senderForAppsDelete sends the AppsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForAppsDelete(ctx context.Context, req *http.Request) (future AppsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
