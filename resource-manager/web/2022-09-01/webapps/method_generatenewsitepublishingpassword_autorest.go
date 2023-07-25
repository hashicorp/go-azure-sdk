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

type GenerateNewSitePublishingPasswordOperationResponse struct {
	HttpResponse *http.Response
}

// GenerateNewSitePublishingPassword ...
func (c WebAppsClient) GenerateNewSitePublishingPassword(ctx context.Context, id SiteId) (result GenerateNewSitePublishingPasswordOperationResponse, err error) {
	req, err := c.preparerForGenerateNewSitePublishingPassword(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPassword", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPassword", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGenerateNewSitePublishingPassword(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GenerateNewSitePublishingPassword", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGenerateNewSitePublishingPassword prepares the GenerateNewSitePublishingPassword request.
func (c WebAppsClient) preparerForGenerateNewSitePublishingPassword(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/newpassword", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGenerateNewSitePublishingPassword handles the response to the GenerateNewSitePublishingPassword request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGenerateNewSitePublishingPassword(resp *http.Response) (result GenerateNewSitePublishingPasswordOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
