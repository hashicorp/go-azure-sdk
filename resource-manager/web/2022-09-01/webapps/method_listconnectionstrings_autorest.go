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

type ListConnectionStringsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConnectionStringDictionary
}

// ListConnectionStrings ...
func (c WebAppsClient) ListConnectionStrings(ctx context.Context, id SiteId) (result ListConnectionStringsOperationResponse, err error) {
	req, err := c.preparerForListConnectionStrings(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStrings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStrings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListConnectionStrings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStrings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListConnectionStrings prepares the ListConnectionStrings request.
func (c WebAppsClient) preparerForListConnectionStrings(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/connectionStrings/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListConnectionStrings handles the response to the ListConnectionStrings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListConnectionStrings(resp *http.Response) (result ListConnectionStringsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
