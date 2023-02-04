package hypervsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSiteOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteSite ...
func (c HyperVSitesClient) DeleteSite(ctx context.Context, id HyperVSiteId) (result DeleteSiteOperationResponse, err error) {
	req, err := c.preparerForDeleteSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "DeleteSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "DeleteSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "DeleteSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSite prepares the DeleteSite request.
func (c HyperVSitesClient) preparerForDeleteSite(ctx context.Context, id HyperVSiteId) (*http.Request, error) {
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

// responderForDeleteSite handles the response to the DeleteSite request. The method always
// closes the http.Response Body.
func (c HyperVSitesClient) responderForDeleteSite(resp *http.Response) (result DeleteSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
