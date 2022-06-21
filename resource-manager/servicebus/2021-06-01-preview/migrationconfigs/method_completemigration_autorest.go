package migrationconfigs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompleteMigrationOperationResponse struct {
	HttpResponse *http.Response
}

// CompleteMigration ...
func (c MigrationConfigsClient) CompleteMigration(ctx context.Context, id ConfigId) (result CompleteMigrationOperationResponse, err error) {
	req, err := c.preparerForCompleteMigration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "CompleteMigration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "CompleteMigration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCompleteMigration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "CompleteMigration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCompleteMigration prepares the CompleteMigration request.
func (c MigrationConfigsClient) preparerForCompleteMigration(ctx context.Context, id ConfigId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/upgrade", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCompleteMigration handles the response to the CompleteMigration request. The method always
// closes the http.Response Body.
func (c MigrationConfigsClient) responderForCompleteMigration(resp *http.Response) (result CompleteMigrationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
