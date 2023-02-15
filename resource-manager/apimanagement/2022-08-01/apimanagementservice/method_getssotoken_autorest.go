package apimanagementservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSsoTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiManagementServiceGetSsoTokenResult
}

// GetSsoToken ...
func (c ApiManagementServiceClient) GetSsoToken(ctx context.Context, id ServiceId) (result GetSsoTokenOperationResponse, err error) {
	req, err := c.preparerForGetSsoToken(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetSsoToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetSsoToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSsoToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagementservice.ApiManagementServiceClient", "GetSsoToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSsoToken prepares the GetSsoToken request.
func (c ApiManagementServiceClient) preparerForGetSsoToken(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getssotoken", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSsoToken handles the response to the GetSsoToken request. The method always
// closes the http.Response Body.
func (c ApiManagementServiceClient) responderForGetSsoToken(resp *http.Response) (result GetSsoTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
