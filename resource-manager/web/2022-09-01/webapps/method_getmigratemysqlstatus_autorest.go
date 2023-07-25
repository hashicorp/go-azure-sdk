package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMigrateMySqlStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *MigrateMySqlStatus
}

// GetMigrateMySqlStatus ...
func (c WebAppsClient) GetMigrateMySqlStatus(ctx context.Context, id SiteId) (result GetMigrateMySqlStatusOperationResponse, err error) {
	req, err := c.preparerForGetMigrateMySqlStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMigrateMySqlStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMigrateMySqlStatus prepares the GetMigrateMySqlStatus request.
func (c WebAppsClient) preparerForGetMigrateMySqlStatus(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrateMySql/status", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetMigrateMySqlStatus handles the response to the GetMigrateMySqlStatus request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetMigrateMySqlStatus(resp *http.Response) (result GetMigrateMySqlStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
