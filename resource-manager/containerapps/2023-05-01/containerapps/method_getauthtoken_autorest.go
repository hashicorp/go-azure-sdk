package containerapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContainerAppAuthToken
}

// GetAuthToken ...
func (c ContainerAppsClient) GetAuthToken(ctx context.Context, id ContainerAppId) (result GetAuthTokenOperationResponse, err error) {
	req, err := c.preparerForGetAuthToken(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerapps.ContainerAppsClient", "GetAuthToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerapps.ContainerAppsClient", "GetAuthToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAuthToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerapps.ContainerAppsClient", "GetAuthToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAuthToken prepares the GetAuthToken request.
func (c ContainerAppsClient) preparerForGetAuthToken(ctx context.Context, id ContainerAppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getAuthtoken", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAuthToken handles the response to the GetAuthToken request. The method always
// closes the http.Response Body.
func (c ContainerAppsClient) responderForGetAuthToken(resp *http.Response) (result GetAuthTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
