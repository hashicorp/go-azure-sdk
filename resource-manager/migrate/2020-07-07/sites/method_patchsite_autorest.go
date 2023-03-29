package sites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMwareSite
}

// PatchSite ...
func (c SitesClient) PatchSite(ctx context.Context, id VMwareSiteId, input VMwareSite) (result PatchSiteOperationResponse, err error) {
	req, err := c.preparerForPatchSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sites.SitesClient", "PatchSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sites.SitesClient", "PatchSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPatchSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sites.SitesClient", "PatchSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPatchSite prepares the PatchSite request.
func (c SitesClient) preparerForPatchSite(ctx context.Context, id VMwareSiteId, input VMwareSite) (*http.Request, error) {
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

// responderForPatchSite handles the response to the PatchSite request. The method always
// closes the http.Response Body.
func (c SitesClient) responderForPatchSite(resp *http.Response) (result PatchSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
