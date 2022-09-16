package sqlpoolsconnectionpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolConnectionPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlPoolConnectionPolicy
}

// SqlPoolConnectionPoliciesGet ...
func (c SqlPoolsConnectionPoliciesClient) SqlPoolConnectionPoliciesGet(ctx context.Context, id SqlPoolId) (result SqlPoolConnectionPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolConnectionPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsconnectionpolicies.SqlPoolsConnectionPoliciesClient", "SqlPoolConnectionPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsconnectionpolicies.SqlPoolsConnectionPoliciesClient", "SqlPoolConnectionPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolConnectionPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsconnectionpolicies.SqlPoolsConnectionPoliciesClient", "SqlPoolConnectionPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolConnectionPoliciesGet prepares the SqlPoolConnectionPoliciesGet request.
func (c SqlPoolsConnectionPoliciesClient) preparerForSqlPoolConnectionPoliciesGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/connectionPolicies/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolConnectionPoliciesGet handles the response to the SqlPoolConnectionPoliciesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsConnectionPoliciesClient) responderForSqlPoolConnectionPoliciesGet(resp *http.Response) (result SqlPoolConnectionPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
