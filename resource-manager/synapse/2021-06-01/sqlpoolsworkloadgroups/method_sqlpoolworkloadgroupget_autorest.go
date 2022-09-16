package sqlpoolsworkloadgroups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolWorkloadGroupGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadGroup
}

// SqlPoolWorkloadGroupGet ...
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupGet(ctx context.Context, id WorkloadGroupId) (result SqlPoolWorkloadGroupGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadGroupGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolWorkloadGroupGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolWorkloadGroupGet prepares the SqlPoolWorkloadGroupGet request.
func (c SqlPoolsWorkloadGroupsClient) preparerForSqlPoolWorkloadGroupGet(ctx context.Context, id WorkloadGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolWorkloadGroupGet handles the response to the SqlPoolWorkloadGroupGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsWorkloadGroupsClient) responderForSqlPoolWorkloadGroupGet(resp *http.Response) (result SqlPoolWorkloadGroupGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
