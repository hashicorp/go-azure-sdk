package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHybridConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// GetHybridConnection ...
func (c AppServicePlansClient) GetHybridConnection(ctx context.Context, id HybridConnectionNamespaceRelayId) (result GetHybridConnectionOperationResponse, err error) {
	req, err := c.preparerForGetHybridConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHybridConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetHybridConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHybridConnection prepares the GetHybridConnection request.
func (c AppServicePlansClient) preparerForGetHybridConnection(ctx context.Context, id HybridConnectionNamespaceRelayId) (*http.Request, error) {
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

// responderForGetHybridConnection handles the response to the GetHybridConnection request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetHybridConnection(resp *http.Response) (result GetHybridConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
