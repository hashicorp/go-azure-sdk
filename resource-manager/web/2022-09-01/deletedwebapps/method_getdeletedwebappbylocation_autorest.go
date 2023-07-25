package deletedwebapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeletedWebAppByLocationOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeletedSite
}

// GetDeletedWebAppByLocation ...
func (c DeletedWebAppsClient) GetDeletedWebAppByLocation(ctx context.Context, id LocationDeletedSiteId) (result GetDeletedWebAppByLocationOperationResponse, err error) {
	req, err := c.preparerForGetDeletedWebAppByLocation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedwebapps.DeletedWebAppsClient", "GetDeletedWebAppByLocation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedwebapps.DeletedWebAppsClient", "GetDeletedWebAppByLocation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeletedWebAppByLocation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedwebapps.DeletedWebAppsClient", "GetDeletedWebAppByLocation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeletedWebAppByLocation prepares the GetDeletedWebAppByLocation request.
func (c DeletedWebAppsClient) preparerForGetDeletedWebAppByLocation(ctx context.Context, id LocationDeletedSiteId) (*http.Request, error) {
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

// responderForGetDeletedWebAppByLocation handles the response to the GetDeletedWebAppByLocation request. The method always
// closes the http.Response Body.
func (c DeletedWebAppsClient) responderForGetDeletedWebAppByLocation(resp *http.Response) (result GetDeletedWebAppByLocationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
