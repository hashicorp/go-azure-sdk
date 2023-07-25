package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHybridConnectionPlanLimitOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnectionLimits
}

// GetHybridConnectionPlanLimit ...
func (c AppServicePlansClient) GetHybridConnectionPlanLimit(ctx context.Context, id ServerFarmId) (result GetHybridConnectionPlanLimitOperationResponse, err error) {
	req, err := c.preparerForGetHybridConnectionPlanLimit(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnectionPlanLimit", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnectionPlanLimit", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHybridConnectionPlanLimit(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnectionPlanLimit", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHybridConnectionPlanLimit prepares the GetHybridConnectionPlanLimit request.
func (c AppServicePlansClient) preparerForGetHybridConnectionPlanLimit(ctx context.Context, id ServerFarmId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridConnectionPlanLimits/limit", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetHybridConnectionPlanLimit handles the response to the GetHybridConnectionPlanLimit request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetHybridConnectionPlanLimit(resp *http.Response) (result GetHybridConnectionPlanLimitOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
