package rbacs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlResourcesListSqlRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlRoleDefinitionListResult
}

// SqlResourcesListSqlRoleDefinitions ...
func (c RbacsClient) SqlResourcesListSqlRoleDefinitions(ctx context.Context, id DatabaseAccountId) (result SqlResourcesListSqlRoleDefinitionsOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesListSqlRoleDefinitions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleDefinitions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleDefinitions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlResourcesListSqlRoleDefinitions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleDefinitions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlResourcesListSqlRoleDefinitions prepares the SqlResourcesListSqlRoleDefinitions request.
func (c RbacsClient) preparerForSqlResourcesListSqlRoleDefinitions(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlRoleDefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlResourcesListSqlRoleDefinitions handles the response to the SqlResourcesListSqlRoleDefinitions request. The method always
// closes the http.Response Body.
func (c RbacsClient) responderForSqlResourcesListSqlRoleDefinitions(resp *http.Response) (result SqlResourcesListSqlRoleDefinitionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
