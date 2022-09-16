package sqlpoolssecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSecurityAlertPoliciesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlPoolSecurityAlertPolicy
}

// SqlPoolSecurityAlertPoliciesCreateOrUpdate ...
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input SqlPoolSecurityAlertPolicy) (result SqlPoolSecurityAlertPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSecurityAlertPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSecurityAlertPoliciesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSecurityAlertPoliciesCreateOrUpdate prepares the SqlPoolSecurityAlertPoliciesCreateOrUpdate request.
func (c SqlPoolsSecurityAlertPoliciesClient) preparerForSqlPoolSecurityAlertPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input SqlPoolSecurityAlertPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securityAlertPolicies/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSecurityAlertPoliciesCreateOrUpdate handles the response to the SqlPoolSecurityAlertPoliciesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsSecurityAlertPoliciesClient) responderForSqlPoolSecurityAlertPoliciesCreateOrUpdate(resp *http.Response) (result SqlPoolSecurityAlertPoliciesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
