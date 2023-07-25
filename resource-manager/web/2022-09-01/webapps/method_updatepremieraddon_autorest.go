package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePremierAddOnOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// UpdatePremierAddOn ...
func (c WebAppsClient) UpdatePremierAddOn(ctx context.Context, id PremierAddonId, input PremierAddOnPatchResource) (result UpdatePremierAddOnOperationResponse, err error) {
	req, err := c.preparerForUpdatePremierAddOn(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdatePremierAddOn", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdatePremierAddOn", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdatePremierAddOn(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdatePremierAddOn", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdatePremierAddOn prepares the UpdatePremierAddOn request.
func (c WebAppsClient) preparerForUpdatePremierAddOn(ctx context.Context, id PremierAddonId, input PremierAddOnPatchResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdatePremierAddOn handles the response to the UpdatePremierAddOn request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdatePremierAddOn(resp *http.Response) (result UpdatePremierAddOnOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
