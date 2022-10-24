package tables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateOperationResponse struct {
	HttpResponse *http.Response
}

// Migrate ...
func (c TablesClient) Migrate(ctx context.Context, id TableId) (result MigrateOperationResponse, err error) {
	req, err := c.preparerForMigrate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "Migrate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "Migrate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMigrate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tables.TablesClient", "Migrate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMigrate prepares the Migrate request.
func (c TablesClient) preparerForMigrate(ctx context.Context, id TableId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForMigrate handles the response to the Migrate request. The method always
// closes the http.Response Body.
func (c TablesClient) responderForMigrate(resp *http.Response) (result MigrateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
