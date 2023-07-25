package webapps

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

type MigrateStorageOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type MigrateStorageOperationOptions struct {
	SubscriptionName *string
}

func DefaultMigrateStorageOperationOptions() MigrateStorageOperationOptions {
	return MigrateStorageOperationOptions{}
}

func (o MigrateStorageOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o MigrateStorageOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.SubscriptionName != nil {
		out["subscriptionName"] = *o.SubscriptionName
	}

	return out
}

// MigrateStorage ...
func (c WebAppsClient) MigrateStorage(ctx context.Context, id SiteId, input StorageMigrationOptions, options MigrateStorageOperationOptions) (result MigrateStorageOperationResponse, err error) {
	req, err := c.preparerForMigrateStorage(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "MigrateStorage", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMigrateStorage(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "MigrateStorage", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MigrateStorageThenPoll performs MigrateStorage then polls until it's completed
func (c WebAppsClient) MigrateStorageThenPoll(ctx context.Context, id SiteId, input StorageMigrationOptions, options MigrateStorageOperationOptions) error {
	result, err := c.MigrateStorage(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing MigrateStorage: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MigrateStorage: %+v", err)
	}

	return nil
}

// preparerForMigrateStorage prepares the MigrateStorage request.
func (c WebAppsClient) preparerForMigrateStorage(ctx context.Context, id SiteId, input StorageMigrationOptions, options MigrateStorageOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/migrate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMigrateStorage sends the MigrateStorage request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForMigrateStorage(ctx context.Context, req *http.Request) (future MigrateStorageOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
