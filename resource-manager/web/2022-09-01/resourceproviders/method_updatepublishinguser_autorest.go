package resourceproviders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePublishingUserOperationResponse struct {
	HttpResponse *http.Response
	Model        *User
}

// UpdatePublishingUser ...
func (c ResourceProvidersClient) UpdatePublishingUser(ctx context.Context, input User) (result UpdatePublishingUserOperationResponse, err error) {
	req, err := c.preparerForUpdatePublishingUser(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdatePublishingUser", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdatePublishingUser", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdatePublishingUser(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "UpdatePublishingUser", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdatePublishingUser prepares the UpdatePublishingUser request.
func (c ResourceProvidersClient) preparerForUpdatePublishingUser(ctx context.Context, input User) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdatePublishingUser handles the response to the UpdatePublishingUser request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForUpdatePublishingUser(resp *http.Response) (result UpdatePublishingUserOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
