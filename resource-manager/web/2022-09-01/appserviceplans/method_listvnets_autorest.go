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

type ListVnetsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VnetInfoResource
}

// ListVnets ...
func (c AppServicePlansClient) ListVnets(ctx context.Context, id ServerFarmId) (result ListVnetsOperationResponse, err error) {
	req, err := c.preparerForListVnets(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListVnets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListVnets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListVnets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListVnets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListVnets prepares the ListVnets request.
func (c AppServicePlansClient) preparerForListVnets(ctx context.Context, id ServerFarmId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/virtualNetworkConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListVnets handles the response to the ListVnets request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListVnets(resp *http.Response) (result ListVnetsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
