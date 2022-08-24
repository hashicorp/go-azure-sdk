package rbacs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlResourcesDeleteSqlRoleAssignmentOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlResourcesDeleteSqlRoleAssignment ...
func (c RbacsClient) SqlResourcesDeleteSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId) (result SqlResourcesDeleteSqlRoleAssignmentOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesDeleteSqlRoleAssignment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesDeleteSqlRoleAssignment", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlResourcesDeleteSqlRoleAssignment(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesDeleteSqlRoleAssignment", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlResourcesDeleteSqlRoleAssignmentThenPoll performs SqlResourcesDeleteSqlRoleAssignment then polls until it's completed
func (c RbacsClient) SqlResourcesDeleteSqlRoleAssignmentThenPoll(ctx context.Context, id SqlRoleAssignmentId) error {
	result, err := c.SqlResourcesDeleteSqlRoleAssignment(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SqlResourcesDeleteSqlRoleAssignment: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlResourcesDeleteSqlRoleAssignment: %+v", err)
	}

	return nil
}

// preparerForSqlResourcesDeleteSqlRoleAssignment prepares the SqlResourcesDeleteSqlRoleAssignment request.
func (c RbacsClient) preparerForSqlResourcesDeleteSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSqlResourcesDeleteSqlRoleAssignment sends the SqlResourcesDeleteSqlRoleAssignment request. The method will close the
// http.Response Body if it receives an error.
func (c RbacsClient) senderForSqlResourcesDeleteSqlRoleAssignment(ctx context.Context, req *http.Request) (future SqlResourcesDeleteSqlRoleAssignmentOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
