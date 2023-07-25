package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPremierAddOnsOperationResponse struct {
	HttpResponse *http.Response
	Model        *PremierAddOn
}

// ListPremierAddOns ...
func (c WebAppsClient) ListPremierAddOns(ctx context.Context, id SiteId) (result ListPremierAddOnsOperationResponse, err error) {
	req, err := c.preparerForListPremierAddOns(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOns", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOns", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListPremierAddOns(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPremierAddOns", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListPremierAddOns prepares the ListPremierAddOns request.
func (c WebAppsClient) preparerForListPremierAddOns(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/premierAddons", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListPremierAddOns handles the response to the ListPremierAddOns request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPremierAddOns(resp *http.Response) (result ListPremierAddOnsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
