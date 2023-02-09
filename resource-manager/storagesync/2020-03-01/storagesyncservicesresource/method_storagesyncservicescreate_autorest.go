package storagesyncservicesresource

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

type StorageSyncServicesCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// StorageSyncServicesCreate ...
func (c StorageSyncServicesResourceClient) StorageSyncServicesCreate(ctx context.Context, id StorageSyncServiceId, input StorageSyncServiceCreateParameters) (result StorageSyncServicesCreateOperationResponse, err error) {
	req, err := c.preparerForStorageSyncServicesCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStorageSyncServicesCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StorageSyncServicesCreateThenPoll performs StorageSyncServicesCreate then polls until it's completed
func (c StorageSyncServicesResourceClient) StorageSyncServicesCreateThenPoll(ctx context.Context, id StorageSyncServiceId, input StorageSyncServiceCreateParameters) error {
	result, err := c.StorageSyncServicesCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing StorageSyncServicesCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StorageSyncServicesCreate: %+v", err)
	}

	return nil
}

// preparerForStorageSyncServicesCreate prepares the StorageSyncServicesCreate request.
func (c StorageSyncServicesResourceClient) preparerForStorageSyncServicesCreate(ctx context.Context, id StorageSyncServiceId, input StorageSyncServiceCreateParameters) (*http.Request, error) {
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

// senderForStorageSyncServicesCreate sends the StorageSyncServicesCreate request. The method will close the
// http.Response Body if it receives an error.
func (c StorageSyncServicesResourceClient) senderForStorageSyncServicesCreate(ctx context.Context, req *http.Request) (future StorageSyncServicesCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
