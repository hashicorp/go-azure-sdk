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

type GetOneDeployStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// GetOneDeployStatus ...
func (c WebAppsClient) GetOneDeployStatus(ctx context.Context, id SiteId) (result GetOneDeployStatusOperationResponse, err error) {
	req, err := c.preparerForGetOneDeployStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetOneDeployStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetOneDeployStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetOneDeployStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetOneDeployStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetOneDeployStatus prepares the GetOneDeployStatus request.
func (c WebAppsClient) preparerForGetOneDeployStatus(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/onedeploy", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetOneDeployStatus handles the response to the GetOneDeployStatus request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetOneDeployStatus(resp *http.Response) (result GetOneDeployStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
