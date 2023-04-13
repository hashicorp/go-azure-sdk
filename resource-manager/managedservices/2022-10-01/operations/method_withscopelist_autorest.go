package operations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WithScopeListOperationResponse struct {
	HttpResponse *http.Response
	Model        *OperationList
}

// WithScopeList ...
func (c OperationsClient) WithScopeList(ctx context.Context, id commonids.ScopeId) (result WithScopeListOperationResponse, err error) {
	req, err := c.preparerForWithScopeList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "WithScopeList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "WithScopeList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWithScopeList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "WithScopeList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWithScopeList prepares the WithScopeList request.
func (c OperationsClient) preparerForWithScopeList(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.ManagedServices/operations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWithScopeList handles the response to the WithScopeList request. The method always
// closes the http.Response Body.
func (c OperationsClient) responderForWithScopeList(resp *http.Response) (result WithScopeListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
