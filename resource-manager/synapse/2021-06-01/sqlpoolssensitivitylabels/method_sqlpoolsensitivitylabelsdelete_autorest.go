package sqlpoolssensitivitylabels

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolSensitivityLabelsDelete ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsDelete(ctx context.Context, id SensitivityLabelSourceId) (result SqlPoolSensitivityLabelsDeleteOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsDelete prepares the SqlPoolSensitivityLabelsDelete request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsDelete(ctx context.Context, id SensitivityLabelSourceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSensitivityLabelsDelete handles the response to the SqlPoolSensitivityLabelsDelete request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsDelete(resp *http.Response) (result SqlPoolSensitivityLabelsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
