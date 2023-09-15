package offboardfromd4api

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APICollectionsOffboardAzureApiManagementApiOperationResponse struct {
	HttpResponse *http.Response
}

// APICollectionsOffboardAzureApiManagementApi ...
func (c OffboardFromD4APIClient) APICollectionsOffboardAzureApiManagementApi(ctx context.Context, id ApiCollectionId) (result APICollectionsOffboardAzureApiManagementApiOperationResponse, err error) {
	req, err := c.preparerForAPICollectionsOffboardAzureApiManagementApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "offboardfromd4api.OffboardFromD4APIClient", "APICollectionsOffboardAzureApiManagementApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "offboardfromd4api.OffboardFromD4APIClient", "APICollectionsOffboardAzureApiManagementApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAPICollectionsOffboardAzureApiManagementApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "offboardfromd4api.OffboardFromD4APIClient", "APICollectionsOffboardAzureApiManagementApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAPICollectionsOffboardAzureApiManagementApi prepares the APICollectionsOffboardAzureApiManagementApi request.
func (c OffboardFromD4APIClient) preparerForAPICollectionsOffboardAzureApiManagementApi(ctx context.Context, id ApiCollectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAPICollectionsOffboardAzureApiManagementApi handles the response to the APICollectionsOffboardAzureApiManagementApi request. The method always
// closes the http.Response Body.
func (c OffboardFromD4APIClient) responderForAPICollectionsOffboardAzureApiManagementApi(resp *http.Response) (result APICollectionsOffboardAzureApiManagementApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
