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

type ListFieldsByModuleAndTypeOperationResponse struct {
	HttpResponse *http.Response
	Model        *TypeFieldListResult
}

// ListFieldsByModuleAndType ...
func (c ObjectDataTypesClient) ListFieldsByModuleAndType(ctx context.Context, id ModuleObjectDataTypeId) (result ListFieldsByModuleAndTypeOperationResponse, err error) {
	req, err := c.preparerForListFieldsByModuleAndType(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByModuleAndType", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByModuleAndType", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFieldsByModuleAndType(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "objectdatatypes.ObjectDataTypesClient", "ListFieldsByModuleAndType", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFieldsByModuleAndType prepares the ListFieldsByModuleAndType request.
func (c ObjectDataTypesClient) preparerForListFieldsByModuleAndType(ctx context.Context, id ModuleObjectDataTypeId) (*http.Request, error) {
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

// responderForListFieldsByModuleAndType handles the response to the ListFieldsByModuleAndType request. The method always
// closes the http.Response Body.
func (c ObjectDataTypesClient) responderForListFieldsByModuleAndType(resp *http.Response) (result ListFieldsByModuleAndTypeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
