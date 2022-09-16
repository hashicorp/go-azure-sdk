package sqlpoolssensitivitylabels

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SensitivityLabel
}

// SqlPoolSensitivityLabelsGet ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsGet(ctx context.Context, id SensitivityLabelSourceId) (result SqlPoolSensitivityLabelsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsGet prepares the SqlPoolSensitivityLabelsGet request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsGet(ctx context.Context, id SensitivityLabelSourceId) (*http.Request, error) {
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

// responderForSqlPoolSensitivityLabelsGet handles the response to the SqlPoolSensitivityLabelsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsGet(resp *http.Response) (result SqlPoolSensitivityLabelsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
