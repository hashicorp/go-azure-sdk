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

type UpdateConnectionStringsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConnectionStringDictionary
}

// UpdateConnectionStrings ...
func (c WebAppsClient) UpdateConnectionStrings(ctx context.Context, id SiteId, input ConnectionStringDictionary) (result UpdateConnectionStringsOperationResponse, err error) {
	req, err := c.preparerForUpdateConnectionStrings(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStrings", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStrings", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateConnectionStrings(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStrings", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateConnectionStrings prepares the UpdateConnectionStrings request.
func (c WebAppsClient) preparerForUpdateConnectionStrings(ctx context.Context, id SiteId, input ConnectionStringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/connectionStrings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateConnectionStrings handles the response to the UpdateConnectionStrings request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateConnectionStrings(resp *http.Response) (result UpdateConnectionStringsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
