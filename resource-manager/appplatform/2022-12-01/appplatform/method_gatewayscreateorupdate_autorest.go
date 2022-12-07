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

type GatewaysCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewaysCreateOrUpdate ...
func (c AppPlatformClient) GatewaysCreateOrUpdate(ctx context.Context, id GatewayId, input GatewayResource) (result GatewaysCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForGatewaysCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewaysCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewaysCreateOrUpdateThenPoll performs GatewaysCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) GatewaysCreateOrUpdateThenPoll(ctx context.Context, id GatewayId, input GatewayResource) error {
	result, err := c.GatewaysCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GatewaysCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewaysCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForGatewaysCreateOrUpdate prepares the GatewaysCreateOrUpdate request.
func (c AppPlatformClient) preparerForGatewaysCreateOrUpdate(ctx context.Context, id GatewayId, input GatewayResource) (*http.Request, error) {
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

// senderForGatewaysCreateOrUpdate sends the GatewaysCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewaysCreateOrUpdate(ctx context.Context, req *http.Request) (future GatewaysCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
