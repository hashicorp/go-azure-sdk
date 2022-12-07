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

type GatewaysDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewaysDelete ...
func (c AppPlatformClient) GatewaysDelete(ctx context.Context, id GatewayId) (result GatewaysDeleteOperationResponse, err error) {
	req, err := c.preparerForGatewaysDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewaysDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewaysDeleteThenPoll performs GatewaysDelete then polls until it's completed
func (c AppPlatformClient) GatewaysDeleteThenPoll(ctx context.Context, id GatewayId) error {
	result, err := c.GatewaysDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing GatewaysDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewaysDelete: %+v", err)
	}

	return nil
}

// preparerForGatewaysDelete prepares the GatewaysDelete request.
func (c AppPlatformClient) preparerForGatewaysDelete(ctx context.Context, id GatewayId) (*http.Request, error) {
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

// senderForGatewaysDelete sends the GatewaysDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewaysDelete(ctx context.Context, req *http.Request) (future GatewaysDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
