package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetProcessModuleOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessModuleInfo
}

// GetProcessModule ...
func (c WebAppsClient) GetProcessModule(ctx context.Context, id ModuleId) (result GetProcessModuleOperationResponse, err error) {
	req, err := c.preparerForGetProcessModule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcessModule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcessModule prepares the GetProcessModule request.
func (c WebAppsClient) preparerForGetProcessModule(ctx context.Context, id ModuleId) (*http.Request, error) {
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

// responderForGetProcessModule handles the response to the GetProcessModule request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcessModule(resp *http.Response) (result GetProcessModuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
