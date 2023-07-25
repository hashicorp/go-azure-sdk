package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHostNameBindingOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostNameBinding
}

// GetHostNameBinding ...
func (c WebAppsClient) GetHostNameBinding(ctx context.Context, id HostNameBindingId) (result GetHostNameBindingOperationResponse, err error) {
	req, err := c.preparerForGetHostNameBinding(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBinding", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBinding", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHostNameBinding(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBinding", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHostNameBinding prepares the GetHostNameBinding request.
func (c WebAppsClient) preparerForGetHostNameBinding(ctx context.Context, id HostNameBindingId) (*http.Request, error) {
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

// responderForGetHostNameBinding handles the response to the GetHostNameBinding request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetHostNameBinding(resp *http.Response) (result GetHostNameBindingOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
