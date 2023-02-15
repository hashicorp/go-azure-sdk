package apimanagementservice

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

type MigrateToStv2OperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MigrateToStv2 ...
func (c ApiManagementServiceClient) MigrateToStv2(ctx context.Context, id ServiceId) (result MigrateToStv2OperationResponse, err error) {
	req, err := c.preparerForMigrateToStv2(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "MigrateToStv2", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMigrateToStv2(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "MigrateToStv2", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MigrateToStv2ThenPoll performs MigrateToStv2 then polls until it's completed
func (c ApiManagementServiceClient) MigrateToStv2ThenPoll(ctx context.Context, id ServiceId) error {
	result, err := c.MigrateToStv2(ctx, id)
	if err != nil {
		return fmt.Errorf("performing MigrateToStv2: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MigrateToStv2: %+v", err)
	}

	return nil
}

// preparerForMigrateToStv2 prepares the MigrateToStv2 request.
func (c ApiManagementServiceClient) preparerForMigrateToStv2(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrateToStv2", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMigrateToStv2 sends the MigrateToStv2 request. The method will close the
// http.Response Body if it receives an error.
func (c ApiManagementServiceClient) senderForMigrateToStv2(ctx context.Context, req *http.Request) (future MigrateToStv2OperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
