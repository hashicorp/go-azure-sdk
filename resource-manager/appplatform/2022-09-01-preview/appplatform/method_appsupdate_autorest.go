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

type AppsUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AppsUpdate ...
func (c AppPlatformClient) AppsUpdate(ctx context.Context, id AppId, input AppResource) (result AppsUpdateOperationResponse, err error) {
	req, err := c.preparerForAppsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAppsUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AppsUpdateThenPoll performs AppsUpdate then polls until it's completed
func (c AppPlatformClient) AppsUpdateThenPoll(ctx context.Context, id AppId, input AppResource) error {
	result, err := c.AppsUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AppsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AppsUpdate: %+v", err)
	}

	return nil
}

// preparerForAppsUpdate prepares the AppsUpdate request.
func (c AppPlatformClient) preparerForAppsUpdate(ctx context.Context, id AppId, input AppResource) (*http.Request, error) {
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

// senderForAppsUpdate sends the AppsUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForAppsUpdate(ctx context.Context, req *http.Request) (future AppsUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
