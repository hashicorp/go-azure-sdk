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

type SqlPoolGeoBackupPoliciesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *GeoBackupPolicy
}

// SqlPoolGeoBackupPoliciesCreateOrUpdate ...
func (c SqlPoolsGeoBackupPoliciesClient) SqlPoolGeoBackupPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input GeoBackupPolicy) (result SqlPoolGeoBackupPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolGeoBackupPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolGeoBackupPoliciesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolGeoBackupPoliciesCreateOrUpdate prepares the SqlPoolGeoBackupPoliciesCreateOrUpdate request.
func (c SqlPoolsGeoBackupPoliciesClient) preparerForSqlPoolGeoBackupPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input GeoBackupPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/geoBackupPolicies/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolGeoBackupPoliciesCreateOrUpdate handles the response to the SqlPoolGeoBackupPoliciesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsGeoBackupPoliciesClient) responderForSqlPoolGeoBackupPoliciesCreateOrUpdate(resp *http.Response) (result SqlPoolGeoBackupPoliciesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
