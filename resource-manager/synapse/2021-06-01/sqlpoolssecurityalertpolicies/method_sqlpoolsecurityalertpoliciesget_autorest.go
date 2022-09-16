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

type SqlPoolSecurityAlertPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlPoolSecurityAlertPolicy
}

// SqlPoolSecurityAlertPoliciesGet ...
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesGet(ctx context.Context, id SqlPoolId) (result SqlPoolSecurityAlertPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSecurityAlertPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSecurityAlertPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSecurityAlertPoliciesGet prepares the SqlPoolSecurityAlertPoliciesGet request.
func (c SqlPoolsSecurityAlertPoliciesClient) preparerForSqlPoolSecurityAlertPoliciesGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securityAlertPolicies/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSecurityAlertPoliciesGet handles the response to the SqlPoolSecurityAlertPoliciesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSecurityAlertPoliciesClient) responderForSqlPoolSecurityAlertPoliciesGet(resp *http.Response) (result SqlPoolSecurityAlertPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
