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

type GetServerFarmSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// GetServerFarmSkus ...
func (c AppServicePlansClient) GetServerFarmSkus(ctx context.Context, id ServerFarmId) (result GetServerFarmSkusOperationResponse, err error) {
	req, err := c.preparerForGetServerFarmSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetServerFarmSkus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetServerFarmSkus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetServerFarmSkus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetServerFarmSkus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetServerFarmSkus prepares the GetServerFarmSkus request.
func (c AppServicePlansClient) preparerForGetServerFarmSkus(ctx context.Context, id ServerFarmId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetServerFarmSkus handles the response to the GetServerFarmSkus request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetServerFarmSkus(resp *http.Response) (result GetServerFarmSkusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
