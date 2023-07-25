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

type MigrateMySqlOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// MigrateMySql ...
func (c WebAppsClient) MigrateMySql(ctx context.Context, id SiteId, input MigrateMySqlRequest) (result MigrateMySqlOperationResponse, err error) {
	req, err := c.preparerForMigrateMySql(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "MigrateMySql", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForMigrateMySql(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "MigrateMySql", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// MigrateMySqlThenPoll performs MigrateMySql then polls until it's completed
func (c WebAppsClient) MigrateMySqlThenPoll(ctx context.Context, id SiteId, input MigrateMySqlRequest) error {
	result, err := c.MigrateMySql(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MigrateMySql: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after MigrateMySql: %+v", err)
	}

	return nil
}

// preparerForMigrateMySql prepares the MigrateMySql request.
func (c WebAppsClient) preparerForMigrateMySql(ctx context.Context, id SiteId, input MigrateMySqlRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrateMySql", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForMigrateMySql sends the MigrateMySql request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForMigrateMySql(ctx context.Context, req *http.Request) (future MigrateMySqlOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
