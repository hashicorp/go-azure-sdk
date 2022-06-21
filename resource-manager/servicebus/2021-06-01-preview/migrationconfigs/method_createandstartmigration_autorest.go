package migrationconfigs

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

type CreateAndStartMigrationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateAndStartMigration ...
func (c MigrationConfigsClient) CreateAndStartMigration(ctx context.Context, id ConfigId, input MigrationConfigProperties) (result CreateAndStartMigrationOperationResponse, err error) {
	req, err := c.preparerForCreateAndStartMigration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "CreateAndStartMigration", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateAndStartMigration(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "CreateAndStartMigration", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateAndStartMigrationThenPoll performs CreateAndStartMigration then polls until it's completed
func (c MigrationConfigsClient) CreateAndStartMigrationThenPoll(ctx context.Context, id ConfigId, input MigrationConfigProperties) error {
	result, err := c.CreateAndStartMigration(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateAndStartMigration: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateAndStartMigration: %+v", err)
	}

	return nil
}

// preparerForCreateAndStartMigration prepares the CreateAndStartMigration request.
func (c MigrationConfigsClient) preparerForCreateAndStartMigration(ctx context.Context, id ConfigId, input MigrationConfigProperties) (*http.Request, error) {
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

// senderForCreateAndStartMigration sends the CreateAndStartMigration request. The method will close the
// http.Response Body if it receives an error.
func (c MigrationConfigsClient) senderForCreateAndStartMigration(ctx context.Context, req *http.Request) (future CreateAndStartMigrationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
