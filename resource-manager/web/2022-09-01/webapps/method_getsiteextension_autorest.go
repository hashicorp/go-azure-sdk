package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteExtensionOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteExtensionInfo
}

// GetSiteExtension ...
func (c WebAppsClient) GetSiteExtension(ctx context.Context, id SiteExtensionId) (result GetSiteExtensionOperationResponse, err error) {
	req, err := c.preparerForGetSiteExtension(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtension", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtension", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSiteExtension(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteExtension", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSiteExtension prepares the GetSiteExtension request.
func (c WebAppsClient) preparerForGetSiteExtension(ctx context.Context, id SiteExtensionId) (*http.Request, error) {
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

// responderForGetSiteExtension handles the response to the GetSiteExtension request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSiteExtension(resp *http.Response) (result GetSiteExtensionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
