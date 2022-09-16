package sqlpoolsgeobackuppolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolGeoBackupPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *GeoBackupPolicy
}

// SqlPoolGeoBackupPoliciesGet ...
func (c SqlPoolsGeoBackupPoliciesClient) SqlPoolGeoBackupPoliciesGet(ctx context.Context, id SqlPoolId) (result SqlPoolGeoBackupPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolGeoBackupPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolGeoBackupPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolGeoBackupPoliciesGet prepares the SqlPoolGeoBackupPoliciesGet request.
func (c SqlPoolsGeoBackupPoliciesClient) preparerForSqlPoolGeoBackupPoliciesGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/geoBackupPolicies/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolGeoBackupPoliciesGet handles the response to the SqlPoolGeoBackupPoliciesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsGeoBackupPoliciesClient) responderForSqlPoolGeoBackupPoliciesGet(resp *http.Response) (result SqlPoolGeoBackupPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
