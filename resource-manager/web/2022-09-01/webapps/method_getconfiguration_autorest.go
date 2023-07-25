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

type GetConfigurationOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// GetConfiguration ...
func (c WebAppsClient) GetConfiguration(ctx context.Context, id SiteId) (result GetConfigurationOperationResponse, err error) {
	req, err := c.preparerForGetConfiguration(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfiguration", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfiguration", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetConfiguration(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfiguration", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetConfiguration prepares the GetConfiguration request.
func (c WebAppsClient) preparerForGetConfiguration(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetConfiguration handles the response to the GetConfiguration request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetConfiguration(resp *http.Response) (result GetConfigurationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
