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

type ListHostKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostKeys
}

// ListHostKeys ...
func (c WebAppsClient) ListHostKeys(ctx context.Context, id SiteId) (result ListHostKeysOperationResponse, err error) {
	req, err := c.preparerForListHostKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListHostKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListHostKeys prepares the ListHostKeys request.
func (c WebAppsClient) preparerForListHostKeys(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListHostKeys handles the response to the ListHostKeys request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHostKeys(resp *http.Response) (result ListHostKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
