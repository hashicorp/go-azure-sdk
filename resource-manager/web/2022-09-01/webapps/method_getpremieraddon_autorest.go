package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPremierAddOnOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// GetPremierAddOn ...
func (c WebAppsClient) GetPremierAddOn(ctx context.Context, id PremierAddonId) (result GetPremierAddOnOperationResponse, err error) {
	req, err := c.preparerForGetPremierAddOn(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOn", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOn", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPremierAddOn(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPremierAddOn", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPremierAddOn prepares the GetPremierAddOn request.
func (c WebAppsClient) preparerForGetPremierAddOn(ctx context.Context, id PremierAddonId) (*http.Request, error) {
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

// responderForGetPremierAddOn handles the response to the GetPremierAddOn request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPremierAddOn(resp *http.Response) (result GetPremierAddOnOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
