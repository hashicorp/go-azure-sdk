package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePremierAddOnOperationResponse struct {
	HttpResponse *http.Response
}

// DeletePremierAddOn ...
func (c WebAppsClient) DeletePremierAddOn(ctx context.Context, id PremierAddonId) (result DeletePremierAddOnOperationResponse, err error) {
	req, err := c.preparerForDeletePremierAddOn(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOn", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOn", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeletePremierAddOn(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePremierAddOn", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeletePremierAddOn prepares the DeletePremierAddOn request.
func (c WebAppsClient) preparerForDeletePremierAddOn(ctx context.Context, id PremierAddonId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeletePremierAddOn handles the response to the DeletePremierAddOn request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeletePremierAddOn(resp *http.Response) (result DeletePremierAddOnOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
