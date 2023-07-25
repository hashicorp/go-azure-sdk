package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *WebSiteInstanceStatus
}

// GetInstanceInfo ...
func (c WebAppsClient) GetInstanceInfo(ctx context.Context, id InstanceId) (result GetInstanceInfoOperationResponse, err error) {
	req, err := c.preparerForGetInstanceInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceInfo prepares the GetInstanceInfo request.
func (c WebAppsClient) preparerForGetInstanceInfo(ctx context.Context, id InstanceId) (*http.Request, error) {
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

// responderForGetInstanceInfo handles the response to the GetInstanceInfo request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceInfo(resp *http.Response) (result GetInstanceInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
