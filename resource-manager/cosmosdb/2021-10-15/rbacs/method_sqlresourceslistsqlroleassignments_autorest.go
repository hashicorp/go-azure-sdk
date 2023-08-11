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

type SqlResourcesListSqlRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlRoleAssignmentListResult
}

// SqlResourcesListSqlRoleAssignments ...
func (c RbacsClient) SqlResourcesListSqlRoleAssignments(ctx context.Context, id DatabaseAccountId) (result SqlResourcesListSqlRoleAssignmentsOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesListSqlRoleAssignments(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleAssignments", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleAssignments", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlResourcesListSqlRoleAssignments(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesListSqlRoleAssignments", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlResourcesListSqlRoleAssignments prepares the SqlResourcesListSqlRoleAssignments request.
func (c RbacsClient) preparerForSqlResourcesListSqlRoleAssignments(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlRoleAssignments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlResourcesListSqlRoleAssignments handles the response to the SqlResourcesListSqlRoleAssignments request. The method always
// closes the http.Response Body.
func (c RbacsClient) responderForSqlResourcesListSqlRoleAssignments(resp *http.Response) (result SqlResourcesListSqlRoleAssignmentsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
