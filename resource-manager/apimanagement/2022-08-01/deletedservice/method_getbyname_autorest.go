package deletedservice

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeletedServiceContract
}

// GetByName ...
func (c DeletedServiceClient) GetByName(ctx context.Context, id DeletedServiceId) (result GetByNameOperationResponse, err error) {
	req, err := c.preparerForGetByName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedservice.DeletedServiceClient", "GetByName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedservice.DeletedServiceClient", "GetByName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deletedservice.DeletedServiceClient", "GetByName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByName prepares the GetByName request.
func (c DeletedServiceClient) preparerForGetByName(ctx context.Context, id DeletedServiceId) (*http.Request, error) {
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

// responderForGetByName handles the response to the GetByName request. The method always
// closes the http.Response Body.
func (c DeletedServiceClient) responderForGetByName(resp *http.Response) (result GetByNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
