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

type GetFunctionsAdminTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *string
}

// GetFunctionsAdminToken ...
func (c WebAppsClient) GetFunctionsAdminToken(ctx context.Context, id SiteId) (result GetFunctionsAdminTokenOperationResponse, err error) {
	req, err := c.preparerForGetFunctionsAdminToken(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFunctionsAdminToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFunctionsAdminToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFunctionsAdminToken prepares the GetFunctionsAdminToken request.
func (c WebAppsClient) preparerForGetFunctionsAdminToken(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/functions/admin/token", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFunctionsAdminToken handles the response to the GetFunctionsAdminToken request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetFunctionsAdminToken(resp *http.Response) (result GetFunctionsAdminTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
