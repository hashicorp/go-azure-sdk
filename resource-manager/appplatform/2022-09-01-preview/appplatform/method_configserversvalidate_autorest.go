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

type ConfigServersValidateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ConfigServersValidate ...
func (c AppPlatformClient) ConfigServersValidate(ctx context.Context, id SpringId, input ConfigServerSettings) (result ConfigServersValidateOperationResponse, err error) {
	req, err := c.preparerForConfigServersValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersValidate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForConfigServersValidate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersValidate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ConfigServersValidateThenPoll performs ConfigServersValidate then polls until it's completed
func (c AppPlatformClient) ConfigServersValidateThenPoll(ctx context.Context, id SpringId, input ConfigServerSettings) error {
	result, err := c.ConfigServersValidate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ConfigServersValidate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ConfigServersValidate: %+v", err)
	}

	return nil
}

// preparerForConfigServersValidate prepares the ConfigServersValidate request.
func (c AppPlatformClient) preparerForConfigServersValidate(ctx context.Context, id SpringId, input ConfigServerSettings) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configServers/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForConfigServersValidate sends the ConfigServersValidate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForConfigServersValidate(ctx context.Context, req *http.Request) (future ConfigServersValidateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
