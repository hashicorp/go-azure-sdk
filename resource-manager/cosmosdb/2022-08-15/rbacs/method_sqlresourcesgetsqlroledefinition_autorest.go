package rbacs

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlResourcesGetSqlRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlRoleDefinitionGetResults
}

// SqlResourcesGetSqlRoleDefinition ...
func (c RbacsClient) SqlResourcesGetSqlRoleDefinition(ctx context.Context, id SqlRoleDefinitionId) (result SqlResourcesGetSqlRoleDefinitionOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesGetSqlRoleDefinition(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleDefinition", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleDefinition", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlResourcesGetSqlRoleDefinition(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleDefinition", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlResourcesGetSqlRoleDefinition prepares the SqlResourcesGetSqlRoleDefinition request.
func (c RbacsClient) preparerForSqlResourcesGetSqlRoleDefinition(ctx context.Context, id SqlRoleDefinitionId) (*http.Request, error) {
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

// responderForSqlResourcesGetSqlRoleDefinition handles the response to the SqlResourcesGetSqlRoleDefinition request. The method always
// closes the http.Response Body.
func (c RbacsClient) responderForSqlResourcesGetSqlRoleDefinition(resp *http.Response) (result SqlResourcesGetSqlRoleDefinitionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
