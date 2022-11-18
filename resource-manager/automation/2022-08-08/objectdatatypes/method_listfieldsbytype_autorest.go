package objectdatatypes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListFieldsByTypeOperationResponse struct {
	HttpResponse *http.Response
	Model        *TypeFieldListResult
}

// ListFieldsByType ...
func (c ObjectDataTypesClient) ListFieldsByType(ctx context.Context, id ObjectDataTypeId) (result ListFieldsByTypeOperationResponse, err error) {
	req, err := c.preparerForListFieldsByType(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByType", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByType", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFieldsByType(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByType", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFieldsByType prepares the ListFieldsByType request.
func (c ObjectDataTypesClient) preparerForListFieldsByType(ctx context.Context, id ObjectDataTypeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/fields", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListFieldsByType handles the response to the ListFieldsByType request. The method always
// closes the http.Response Body.
func (c ObjectDataTypesClient) responderForListFieldsByType(resp *http.Response) (result ListFieldsByTypeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
