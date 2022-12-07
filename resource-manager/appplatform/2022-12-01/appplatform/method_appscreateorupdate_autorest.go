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

type AppsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AppsCreateOrUpdate ...
func (c AppPlatformClient) AppsCreateOrUpdate(ctx context.Context, id AppId, input AppResource) (result AppsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForAppsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAppsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AppsCreateOrUpdateThenPoll performs AppsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) AppsCreateOrUpdateThenPoll(ctx context.Context, id AppId, input AppResource) error {
	result, err := c.AppsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AppsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AppsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForAppsCreateOrUpdate prepares the AppsCreateOrUpdate request.
func (c AppPlatformClient) preparerForAppsCreateOrUpdate(ctx context.Context, id AppId, input AppResource) (*http.Request, error) {
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

// senderForAppsCreateOrUpdate sends the AppsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForAppsCreateOrUpdate(ctx context.Context, req *http.Request) (future AppsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
