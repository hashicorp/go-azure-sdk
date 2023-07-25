package resourceproviders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPublishingUserOperationResponse struct {
	HttpResponse *http.Response
	Model        *User
}

// GetPublishingUser ...
func (c ResourceProvidersClient) GetPublishingUser(ctx context.Context) (result GetPublishingUserOperationResponse, err error) {
	req, err := c.preparerForGetPublishingUser(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetPublishingUser", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetPublishingUser", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPublishingUser(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "GetPublishingUser", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPublishingUser prepares the GetPublishingUser request.
func (c ResourceProvidersClient) preparerForGetPublishingUser(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetPublishingUser handles the response to the GetPublishingUser request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForGetPublishingUser(resp *http.Response) (result GetPublishingUserOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
