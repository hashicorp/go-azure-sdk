package sqlpoolssensitivitylabels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsUpdateOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolSensitivityLabelsUpdate ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsUpdate(ctx context.Context, id SqlPoolId, input SensitivityLabelUpdateList) (result SqlPoolSensitivityLabelsUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsUpdate prepares the SqlPoolSensitivityLabelsUpdate request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsUpdate(ctx context.Context, id SqlPoolId, input SensitivityLabelUpdateList) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/currentSensitivityLabels", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSensitivityLabelsUpdate handles the response to the SqlPoolSensitivityLabelsUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsUpdate(resp *http.Response) (result SqlPoolSensitivityLabelsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
