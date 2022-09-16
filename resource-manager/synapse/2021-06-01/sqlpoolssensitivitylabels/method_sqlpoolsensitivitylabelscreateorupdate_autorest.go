package sqlpoolssensitivitylabels

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SensitivityLabel
}

// SqlPoolSensitivityLabelsCreateOrUpdate ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsCreateOrUpdate(ctx context.Context, id SensitivityLabelSourceId, input SensitivityLabel) (result SqlPoolSensitivityLabelsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsCreateOrUpdate prepares the SqlPoolSensitivityLabelsCreateOrUpdate request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsCreateOrUpdate(ctx context.Context, id SensitivityLabelSourceId, input SensitivityLabel) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSensitivityLabelsCreateOrUpdate handles the response to the SqlPoolSensitivityLabelsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsCreateOrUpdate(resp *http.Response) (result SqlPoolSensitivityLabelsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
