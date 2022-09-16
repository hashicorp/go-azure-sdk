package sensitivitylabels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolSensitivityLabelsEnableRecommendationOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolSensitivityLabelsEnableRecommendation ...
func (c SensitivityLabelsClient) SqlPoolSensitivityLabelsEnableRecommendation(ctx context.Context, id ColumnId) (result SqlPoolSensitivityLabelsEnableRecommendationOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsEnableRecommendation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsEnableRecommendation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsEnableRecommendation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsEnableRecommendation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsEnableRecommendation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsEnableRecommendation prepares the SqlPoolSensitivityLabelsEnableRecommendation request.
func (c SensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsEnableRecommendation(ctx context.Context, id ColumnId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sensitivityLabels/recommended/enable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSensitivityLabelsEnableRecommendation handles the response to the SqlPoolSensitivityLabelsEnableRecommendation request. The method always
// closes the http.Response Body.
func (c SensitivityLabelsClient) responderForSqlPoolSensitivityLabelsEnableRecommendation(resp *http.Response) (result SqlPoolSensitivityLabelsEnableRecommendationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
