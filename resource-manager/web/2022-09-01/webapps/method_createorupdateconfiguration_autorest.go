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

type CreateOrUpdateConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// CreateOrUpdateConfiguration ...
func (c WebAppsClient) CreateOrUpdateConfiguration(ctx context.Context, id SiteId, input SiteConfigResource) (result CreateOrUpdateConfigurationOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateConfiguration(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateConfiguration prepares the CreateOrUpdateConfiguration request.
func (c WebAppsClient) preparerForCreateOrUpdateConfiguration(ctx context.Context, id SiteId, input SiteConfigResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateConfiguration handles the response to the CreateOrUpdateConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateConfiguration(resp *http.Response) (result CreateOrUpdateConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
