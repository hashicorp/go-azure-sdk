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

type SqlPoolRecommendedSensitivityLabelsUpdateOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolRecommendedSensitivityLabelsUpdate ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolRecommendedSensitivityLabelsUpdate(ctx context.Context, id SqlPoolId, input RecommendedSensitivityLabelUpdateList) (result SqlPoolRecommendedSensitivityLabelsUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolRecommendedSensitivityLabelsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolRecommendedSensitivityLabelsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolRecommendedSensitivityLabelsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolRecommendedSensitivityLabelsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolRecommendedSensitivityLabelsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolRecommendedSensitivityLabelsUpdate prepares the SqlPoolRecommendedSensitivityLabelsUpdate request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolRecommendedSensitivityLabelsUpdate(ctx context.Context, id SqlPoolId, input RecommendedSensitivityLabelUpdateList) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recommendedSensitivityLabels", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolRecommendedSensitivityLabelsUpdate handles the response to the SqlPoolRecommendedSensitivityLabelsUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolRecommendedSensitivityLabelsUpdate(resp *http.Response) (result SqlPoolRecommendedSensitivityLabelsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
