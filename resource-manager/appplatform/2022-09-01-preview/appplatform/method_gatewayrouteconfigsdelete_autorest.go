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

type GatewayRouteConfigsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewayRouteConfigsDelete ...
func (c AppPlatformClient) GatewayRouteConfigsDelete(ctx context.Context, id RouteConfigId) (result GatewayRouteConfigsDeleteOperationResponse, err error) {
	req, err := c.preparerForGatewayRouteConfigsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewayRouteConfigsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewayRouteConfigsDeleteThenPoll performs GatewayRouteConfigsDelete then polls until it's completed
func (c AppPlatformClient) GatewayRouteConfigsDeleteThenPoll(ctx context.Context, id RouteConfigId) error {
	result, err := c.GatewayRouteConfigsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing GatewayRouteConfigsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewayRouteConfigsDelete: %+v", err)
	}

	return nil
}

// preparerForGatewayRouteConfigsDelete prepares the GatewayRouteConfigsDelete request.
func (c AppPlatformClient) preparerForGatewayRouteConfigsDelete(ctx context.Context, id RouteConfigId) (*http.Request, error) {
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

// senderForGatewayRouteConfigsDelete sends the GatewayRouteConfigsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewayRouteConfigsDelete(ctx context.Context, req *http.Request) (future GatewayRouteConfigsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
