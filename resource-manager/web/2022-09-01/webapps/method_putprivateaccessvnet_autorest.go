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

type PutPrivateAccessVnetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateAccess
}

// PutPrivateAccessVnet ...
func (c WebAppsClient) PutPrivateAccessVnet(ctx context.Context, id SiteId, input PrivateAccess) (result PutPrivateAccessVnetOperationResponse, err error) {
	req, err := c.preparerForPutPrivateAccessVnet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutPrivateAccessVnet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutPrivateAccessVnet prepares the PutPrivateAccessVnet request.
func (c WebAppsClient) preparerForPutPrivateAccessVnet(ctx context.Context, id SiteId, input PrivateAccess) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateAccess/virtualNetworks", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPutPrivateAccessVnet handles the response to the PutPrivateAccessVnet request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForPutPrivateAccessVnet(resp *http.Response) (result PutPrivateAccessVnetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
