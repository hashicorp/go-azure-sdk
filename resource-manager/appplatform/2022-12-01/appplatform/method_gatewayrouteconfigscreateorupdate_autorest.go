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

type GatewayRouteConfigsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewayRouteConfigsCreateOrUpdate ...
func (c AppPlatformClient) GatewayRouteConfigsCreateOrUpdate(ctx context.Context, id RouteConfigId, input GatewayRouteConfigResource) (result GatewayRouteConfigsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForGatewayRouteConfigsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewayRouteConfigsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewayRouteConfigsCreateOrUpdateThenPoll performs GatewayRouteConfigsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) GatewayRouteConfigsCreateOrUpdateThenPoll(ctx context.Context, id RouteConfigId, input GatewayRouteConfigResource) error {
	result, err := c.GatewayRouteConfigsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GatewayRouteConfigsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewayRouteConfigsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForGatewayRouteConfigsCreateOrUpdate prepares the GatewayRouteConfigsCreateOrUpdate request.
func (c AppPlatformClient) preparerForGatewayRouteConfigsCreateOrUpdate(ctx context.Context, id RouteConfigId, input GatewayRouteConfigResource) (*http.Request, error) {
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

// senderForGatewayRouteConfigsCreateOrUpdate sends the GatewayRouteConfigsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewayRouteConfigsCreateOrUpdate(ctx context.Context, req *http.Request) (future GatewayRouteConfigsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
