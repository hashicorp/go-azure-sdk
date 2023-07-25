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

type GetPrivateAccessOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateAccess
}

// GetPrivateAccess ...
func (c WebAppsClient) GetPrivateAccess(ctx context.Context, id SiteId) (result GetPrivateAccessOperationResponse, err error) {
	req, err := c.preparerForGetPrivateAccess(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateAccess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateAccess prepares the GetPrivateAccess request.
func (c WebAppsClient) preparerForGetPrivateAccess(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateAccess/virtualNetworks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetPrivateAccess handles the response to the GetPrivateAccess request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateAccess(resp *http.Response) (result GetPrivateAccessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
