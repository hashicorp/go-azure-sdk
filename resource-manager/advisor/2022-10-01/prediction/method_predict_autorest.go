package prediction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictOperationResponse struct {
	HttpResponse *http.Response
	Model        *PredictionResponse
}

// Predict ...
func (c PredictionClient) Predict(ctx context.Context, id commonids.SubscriptionId, input PredictionRequest) (result PredictOperationResponse, err error) {
	req, err := c.preparerForPredict(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.PredictionClient", "Predict", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.PredictionClient", "Predict", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPredict(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.PredictionClient", "Predict", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPredict prepares the Predict request.
func (c PredictionClient) preparerForPredict(ctx context.Context, id commonids.SubscriptionId, input PredictionRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Advisor/predict", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPredict handles the response to the Predict request. The method always
// closes the http.Response Body.
func (c PredictionClient) responderForPredict(resp *http.Response) (result PredictOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
