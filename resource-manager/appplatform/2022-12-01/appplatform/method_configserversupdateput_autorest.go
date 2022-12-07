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

type ConfigServersUpdatePutOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigServersUpdatePut ...
func (c AppPlatformClient) ConfigServersUpdatePut(ctx context.Context, id SpringId, input ConfigServerResource) (result ConfigServersUpdatePutOperationResponse, err error) {
	req, err := c.preparerForConfigServersUpdatePut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersUpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigServersUpdatePut(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersUpdatePut", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigServersUpdatePutThenPoll performs ConfigServersUpdatePut then polls until it's completed
func (c AppPlatformClient) ConfigServersUpdatePutThenPoll(ctx context.Context, id SpringId, input ConfigServerResource) error {
	result, err := c.ConfigServersUpdatePut(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConfigServersUpdatePut: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigServersUpdatePut: %+v", err)
	}

	return nil
}

// preparerForConfigServersUpdatePut prepares the ConfigServersUpdatePut request.
func (c AppPlatformClient) preparerForConfigServersUpdatePut(ctx context.Context, id SpringId, input ConfigServerResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configServers/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForConfigServersUpdatePut sends the ConfigServersUpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigServersUpdatePut(ctx context.Context, req *http.Request) (future ConfigServersUpdatePutOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
