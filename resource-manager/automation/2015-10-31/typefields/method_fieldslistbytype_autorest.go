package typefields

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FieldsListByTypeOperationResponse struct {
	HttpResponse *http.Response
	Model        *TypeFieldListResult
}

// FieldsListByType ...
func (c TypeFieldsClient) FieldsListByType(ctx context.Context, id TypeId) (result FieldsListByTypeOperationResponse, err error) {
	req, err := c.preparerForFieldsListByType(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "typefields.TypeFieldsClient", "FieldsListByType", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "typefields.TypeFieldsClient", "FieldsListByType", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFieldsListByType(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "typefields.TypeFieldsClient", "FieldsListByType", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFieldsListByType prepares the FieldsListByType request.
func (c TypeFieldsClient) preparerForFieldsListByType(ctx context.Context, id TypeId) (*http.Request, error) {
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

// responderForFieldsListByType handles the response to the FieldsListByType request. The method always
// closes the http.Response Body.
func (c TypeFieldsClient) responderForFieldsListByType(resp *http.Response) (result FieldsListByTypeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
