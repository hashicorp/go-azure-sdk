package sqlpoolsworkloadclassifiers

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolWorkloadClassifierGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadClassifier
}

// SqlPoolWorkloadClassifierGet ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierGet(ctx context.Context, id WorkloadClassifierId) (result SqlPoolWorkloadClassifierGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadClassifierGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolWorkloadClassifierGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolWorkloadClassifierGet prepares the SqlPoolWorkloadClassifierGet request.
func (c SqlPoolsWorkloadClassifiersClient) preparerForSqlPoolWorkloadClassifierGet(ctx context.Context, id WorkloadClassifierId) (*http.Request, error) {
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

// responderForSqlPoolWorkloadClassifierGet handles the response to the SqlPoolWorkloadClassifierGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsWorkloadClassifiersClient) responderForSqlPoolWorkloadClassifierGet(resp *http.Response) (result SqlPoolWorkloadClassifierGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
