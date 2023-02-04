package hypervsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *HyperVSite
}

// GetSite ...
func (c HyperVSitesClient) GetSite(ctx context.Context, id HyperVSiteId) (result GetSiteOperationResponse, err error) {
	req, err := c.preparerForGetSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "GetSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSite prepares the GetSite request.
func (c HyperVSitesClient) preparerForGetSite(ctx context.Context, id HyperVSiteId) (*http.Request, error) {
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

// responderForGetSite handles the response to the GetSite request. The method always
// closes the http.Response Body.
func (c HyperVSitesClient) responderForGetSite(resp *http.Response) (result GetSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
