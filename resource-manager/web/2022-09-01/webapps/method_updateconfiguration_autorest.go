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

type UpdateConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// UpdateConfiguration ...
func (c WebAppsClient) UpdateConfiguration(ctx context.Context, id SiteId, input SiteConfigResource) (result UpdateConfigurationOperationResponse, err error) {
	req, err := c.preparerForUpdateConfiguration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateConfiguration prepares the UpdateConfiguration request.
func (c WebAppsClient) preparerForUpdateConfiguration(ctx context.Context, id SiteId, input SiteConfigResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateConfiguration handles the response to the UpdateConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateConfiguration(resp *http.Response) (result UpdateConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
