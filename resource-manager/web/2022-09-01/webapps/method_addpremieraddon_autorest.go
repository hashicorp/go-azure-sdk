package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddPremierAddOnOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// AddPremierAddOn ...
func (c WebAppsClient) AddPremierAddOn(ctx context.Context, id PremierAddonId, input PremierAddOn) (result AddPremierAddOnOperationResponse, err error) {
	req, err := c.preparerForAddPremierAddOn(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOn", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOn", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAddPremierAddOn(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "AddPremierAddOn", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAddPremierAddOn prepares the AddPremierAddOn request.
func (c WebAppsClient) preparerForAddPremierAddOn(ctx context.Context, id PremierAddonId, input PremierAddOn) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAddPremierAddOn handles the response to the AddPremierAddOn request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForAddPremierAddOn(resp *http.Response) (result AddPremierAddOnOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
