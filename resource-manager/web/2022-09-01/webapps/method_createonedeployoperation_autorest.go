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

type CreateOneDeployOperationOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// CreateOneDeployOperation ...
func (c WebAppsClient) CreateOneDeployOperation(ctx context.Context, id SiteId) (result CreateOneDeployOperationOperationResponse, err error) {
	req, err := c.preparerForCreateOneDeployOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOneDeployOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOneDeployOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOneDeployOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOneDeployOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOneDeployOperation prepares the CreateOneDeployOperation request.
func (c WebAppsClient) preparerForCreateOneDeployOperation(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/onedeploy", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOneDeployOperation handles the response to the CreateOneDeployOperation request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOneDeployOperation(resp *http.Response) (result CreateOneDeployOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
