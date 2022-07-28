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

type RevertOperationResponse struct {
	HttpResponse *http.Response
}

// Revert ...
func (c MigrationConfigsClient) Revert(ctx context.Context, id NamespaceId) (result RevertOperationResponse, err error) {
	req, err := c.preparerForRevert(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "Revert", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "Revert", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRevert(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationconfigs.MigrationConfigsClient", "Revert", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRevert prepares the Revert request.
func (c MigrationConfigsClient) preparerForRevert(ctx context.Context, id NamespaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrationConfigurations/$default/revert", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRevert handles the response to the Revert request. The method always
// closes the http.Response Body.
func (c MigrationConfigsClient) responderForRevert(resp *http.Response) (result RevertOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
