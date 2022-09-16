package sqlpoolssqlpooluseractivities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolDataWarehouseUserActivitiesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataWarehouseUserActivities
}

// SqlPoolDataWarehouseUserActivitiesGet ...
func (c SqlPoolsSqlPoolUserActivitiesClient) SqlPoolDataWarehouseUserActivitiesGet(ctx context.Context, id SqlPoolId) (result SqlPoolDataWarehouseUserActivitiesGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolDataWarehouseUserActivitiesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooluseractivities.SqlPoolsSqlPoolUserActivitiesClient", "SqlPoolDataWarehouseUserActivitiesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooluseractivities.SqlPoolsSqlPoolUserActivitiesClient", "SqlPoolDataWarehouseUserActivitiesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolDataWarehouseUserActivitiesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssqlpooluseractivities.SqlPoolsSqlPoolUserActivitiesClient", "SqlPoolDataWarehouseUserActivitiesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolDataWarehouseUserActivitiesGet prepares the SqlPoolDataWarehouseUserActivitiesGet request.
func (c SqlPoolsSqlPoolUserActivitiesClient) preparerForSqlPoolDataWarehouseUserActivitiesGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dataWarehouseUserActivities/current", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolDataWarehouseUserActivitiesGet handles the response to the SqlPoolDataWarehouseUserActivitiesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSqlPoolUserActivitiesClient) responderForSqlPoolDataWarehouseUserActivitiesGet(resp *http.Response) (result SqlPoolDataWarehouseUserActivitiesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
