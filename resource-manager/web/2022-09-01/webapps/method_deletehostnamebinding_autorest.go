package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHostNameBindingOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHostNameBinding ...
func (c WebAppsClient) DeleteHostNameBinding(ctx context.Context, id HostNameBindingId) (result DeleteHostNameBindingOperationResponse, err error) {
	req, err := c.preparerForDeleteHostNameBinding(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBinding", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBinding", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHostNameBinding(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBinding", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHostNameBinding prepares the DeleteHostNameBinding request.
func (c WebAppsClient) preparerForDeleteHostNameBinding(ctx context.Context, id HostNameBindingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteHostNameBinding handles the response to the DeleteHostNameBinding request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHostNameBinding(resp *http.Response) (result DeleteHostNameBindingOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
