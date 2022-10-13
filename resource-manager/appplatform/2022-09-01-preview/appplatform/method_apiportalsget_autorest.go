package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiPortalsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiPortalResource
}

// ApiPortalsGet ...
func (c AppPlatformClient) ApiPortalsGet(ctx context.Context, id ApiPortalId) (result ApiPortalsGetOperationResponse, err error) {
	req, err := c.preparerForApiPortalsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApiPortalsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApiPortalsGet prepares the ApiPortalsGet request.
func (c AppPlatformClient) preparerForApiPortalsGet(ctx context.Context, id ApiPortalId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForApiPortalsGet handles the response to the ApiPortalsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForApiPortalsGet(resp *http.Response) (result ApiPortalsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
