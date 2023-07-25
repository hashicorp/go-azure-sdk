package global

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeletedWebAppOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeletedSite
}

// GetDeletedWebApp ...
func (c GlobalClient) GetDeletedWebApp(ctx context.Context, id DeletedSiteId) (result GetDeletedWebAppOperationResponse, err error) {
	req, err := c.preparerForGetDeletedWebApp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebApp", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebApp", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeletedWebApp(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "global.GlobalClient", "GetDeletedWebApp", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeletedWebApp prepares the GetDeletedWebApp request.
func (c GlobalClient) preparerForGetDeletedWebApp(ctx context.Context, id DeletedSiteId) (*http.Request, error) {
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

// responderForGetDeletedWebApp handles the response to the GetDeletedWebApp request. The method always
// closes the http.Response Body.
func (c GlobalClient) responderForGetDeletedWebApp(resp *http.Response) (result GetDeletedWebAppOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
