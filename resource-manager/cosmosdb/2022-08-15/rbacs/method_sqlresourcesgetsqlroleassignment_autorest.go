package rbacs

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlResourcesGetSqlRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	Model        *SqlRoleAssignmentGetResults
}

// SqlResourcesGetSqlRoleAssignment ...
func (c RbacsClient) SqlResourcesGetSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId) (result SqlResourcesGetSqlRoleAssignmentOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesGetSqlRoleAssignment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleAssignment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleAssignment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlResourcesGetSqlRoleAssignment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesGetSqlRoleAssignment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlResourcesGetSqlRoleAssignment prepares the SqlResourcesGetSqlRoleAssignment request.
func (c RbacsClient) preparerForSqlResourcesGetSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId) (*http.Request, error) {
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

// responderForSqlResourcesGetSqlRoleAssignment handles the response to the SqlResourcesGetSqlRoleAssignment request. The method always
// closes the http.Response Body.
func (c RbacsClient) responderForSqlResourcesGetSqlRoleAssignment(resp *http.Response) (result SqlResourcesGetSqlRoleAssignmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
