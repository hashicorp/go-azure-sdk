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

type BuildpackBindingDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BuildpackBindingDelete ...
func (c AppPlatformClient) BuildpackBindingDelete(ctx context.Context, id BuildpackBindingId) (result BuildpackBindingDeleteOperationResponse, err error) {
	req, err := c.preparerForBuildpackBindingDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBuildpackBindingDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BuildpackBindingDeleteThenPoll performs BuildpackBindingDelete then polls until it's completed
func (c AppPlatformClient) BuildpackBindingDeleteThenPoll(ctx context.Context, id BuildpackBindingId) error {
	result, err := c.BuildpackBindingDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing BuildpackBindingDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BuildpackBindingDelete: %+v", err)
	}

	return nil
}

// preparerForBuildpackBindingDelete prepares the BuildpackBindingDelete request.
func (c AppPlatformClient) preparerForBuildpackBindingDelete(ctx context.Context, id BuildpackBindingId) (*http.Request, error) {
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

// senderForBuildpackBindingDelete sends the BuildpackBindingDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBuildpackBindingDelete(ctx context.Context, req *http.Request) (future BuildpackBindingDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
