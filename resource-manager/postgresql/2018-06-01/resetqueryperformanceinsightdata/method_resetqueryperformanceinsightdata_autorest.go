package resetqueryperformanceinsightdata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetQueryPerformanceInsightDataOperationResponse struct {
	HttpResponse *http.Response
	Model        *QueryPerformanceInsightResetDataResult
}

// ResetQueryPerformanceInsightData ...
func (c ResetQueryPerformanceInsightDataClient) ResetQueryPerformanceInsightData(ctx context.Context, id ServerId) (result ResetQueryPerformanceInsightDataOperationResponse, err error) {
	req, err := c.preparerForResetQueryPerformanceInsightData(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resetqueryperformanceinsightdata.ResetQueryPerformanceInsightDataClient", "ResetQueryPerformanceInsightData", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resetqueryperformanceinsightdata.ResetQueryPerformanceInsightDataClient", "ResetQueryPerformanceInsightData", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetQueryPerformanceInsightData(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resetqueryperformanceinsightdata.ResetQueryPerformanceInsightDataClient", "ResetQueryPerformanceInsightData", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetQueryPerformanceInsightData prepares the ResetQueryPerformanceInsightData request.
func (c ResetQueryPerformanceInsightDataClient) preparerForResetQueryPerformanceInsightData(ctx context.Context, id ServerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resetQueryPerformanceInsightData", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetQueryPerformanceInsightData handles the response to the ResetQueryPerformanceInsightData request. The method always
// closes the http.Response Body.
func (c ResetQueryPerformanceInsightDataClient) responderForResetQueryPerformanceInsightData(resp *http.Response) (result ResetQueryPerformanceInsightDataOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
