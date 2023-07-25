package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Operation
}

// ListOperations ...
func (c AppServiceEnvironmentsClient) ListOperations(ctx context.Context, id HostingEnvironmentId) (result ListOperationsOperationResponse, err error) {
	req, err := c.preparerForListOperations(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListOperations", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListOperations(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListOperations", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListOperations prepares the ListOperations request.
func (c AppServiceEnvironmentsClient) preparerForListOperations(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/operations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListOperations handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListOperations(resp *http.Response) (result ListOperationsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
