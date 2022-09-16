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

type SqlPoolSensitivityLabelsDisableRecommendationOperationResponse struct {
	HttpResponse *http.Response
}

// SqlPoolSensitivityLabelsDisableRecommendation ...
func (c SensitivityLabelsClient) SqlPoolSensitivityLabelsDisableRecommendation(ctx context.Context, id ColumnId) (result SqlPoolSensitivityLabelsDisableRecommendationOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsDisableRecommendation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsDisableRecommendation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsDisableRecommendation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolSensitivityLabelsDisableRecommendation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sensitivitylabels.SensitivityLabelsClient", "SqlPoolSensitivityLabelsDisableRecommendation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolSensitivityLabelsDisableRecommendation prepares the SqlPoolSensitivityLabelsDisableRecommendation request.
func (c SensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsDisableRecommendation(ctx context.Context, id ColumnId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sensitivityLabels/recommended/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolSensitivityLabelsDisableRecommendation handles the response to the SqlPoolSensitivityLabelsDisableRecommendation request. The method always
// closes the http.Response Body.
func (c SensitivityLabelsClient) responderForSqlPoolSensitivityLabelsDisableRecommendation(resp *http.Response) (result SqlPoolSensitivityLabelsDisableRecommendationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
