package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceProcessModuleOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessModuleInfo
}

// GetInstanceProcessModule ...
func (c WebAppsClient) GetInstanceProcessModule(ctx context.Context, id InstanceProcessModuleId) (result GetInstanceProcessModuleOperationResponse, err error) {
	req, err := c.preparerForGetInstanceProcessModule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessModule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessModule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceProcessModule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessModule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceProcessModule prepares the GetInstanceProcessModule request.
func (c WebAppsClient) preparerForGetInstanceProcessModule(ctx context.Context, id InstanceProcessModuleId) (*http.Request, error) {
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

// responderForGetInstanceProcessModule handles the response to the GetInstanceProcessModule request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceProcessModule(resp *http.Response) (result GetInstanceProcessModuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
