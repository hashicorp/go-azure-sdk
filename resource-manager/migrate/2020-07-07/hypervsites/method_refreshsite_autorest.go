package hypervsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefreshSiteOperationResponse struct {
	HttpResponse *http.Response
}

// RefreshSite ...
func (c HyperVSitesClient) RefreshSite(ctx context.Context, id HyperVSiteId) (result RefreshSiteOperationResponse, err error) {
	req, err := c.preparerForRefreshSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "RefreshSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "RefreshSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRefreshSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "RefreshSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRefreshSite prepares the RefreshSite request.
func (c HyperVSitesClient) preparerForRefreshSite(ctx context.Context, id HyperVSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/refresh", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRefreshSite handles the response to the RefreshSite request. The method always
// closes the http.Response Body.
func (c HyperVSitesClient) responderForRefreshSite(resp *http.Response) (result RefreshSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
