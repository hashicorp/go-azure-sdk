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

type ConfigServersUpdatePatchOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigServersUpdatePatch ...
func (c AppPlatformClient) ConfigServersUpdatePatch(ctx context.Context, id SpringId, input ConfigServerResource) (result ConfigServersUpdatePatchOperationResponse, err error) {
	req, err := c.preparerForConfigServersUpdatePatch(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersUpdatePatch", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigServersUpdatePatch(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersUpdatePatch", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigServersUpdatePatchThenPoll performs ConfigServersUpdatePatch then polls until it's completed
func (c AppPlatformClient) ConfigServersUpdatePatchThenPoll(ctx context.Context, id SpringId, input ConfigServerResource) error {
	result, err := c.ConfigServersUpdatePatch(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConfigServersUpdatePatch: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigServersUpdatePatch: %+v", err)
	}

	return nil
}

// preparerForConfigServersUpdatePatch prepares the ConfigServersUpdatePatch request.
func (c AppPlatformClient) preparerForConfigServersUpdatePatch(ctx context.Context, id SpringId, input ConfigServerResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configServers/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForConfigServersUpdatePatch sends the ConfigServersUpdatePatch request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigServersUpdatePatch(ctx context.Context, req *http.Request) (future ConfigServersUpdatePatchOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
