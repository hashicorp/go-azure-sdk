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

type SqlPoolGeoBackupPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *GeoBackupPolicyListResult
}

// SqlPoolGeoBackupPoliciesList ...
func (c SqlPoolsGeoBackupPoliciesClient) SqlPoolGeoBackupPoliciesList(ctx context.Context, id SqlPoolId) (result SqlPoolGeoBackupPoliciesListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolGeoBackupPoliciesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolGeoBackupPoliciesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient", "SqlPoolGeoBackupPoliciesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolGeoBackupPoliciesList prepares the SqlPoolGeoBackupPoliciesList request.
func (c SqlPoolsGeoBackupPoliciesClient) preparerForSqlPoolGeoBackupPoliciesList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/geoBackupPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolGeoBackupPoliciesList handles the response to the SqlPoolGeoBackupPoliciesList request. The method always
// closes the http.Response Body.
func (c SqlPoolsGeoBackupPoliciesClient) responderForSqlPoolGeoBackupPoliciesList(resp *http.Response) (result SqlPoolGeoBackupPoliciesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
