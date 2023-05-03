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

type SqlResourcesCreateUpdateSqlRoleAssignmentOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlResourcesCreateUpdateSqlRoleAssignment ...
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId, input SqlRoleAssignmentCreateUpdateParameters) (result SqlResourcesCreateUpdateSqlRoleAssignmentOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesCreateUpdateSqlRoleAssignment(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesCreateUpdateSqlRoleAssignment", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlResourcesCreateUpdateSqlRoleAssignment(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesCreateUpdateSqlRoleAssignment", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlResourcesCreateUpdateSqlRoleAssignmentThenPoll performs SqlResourcesCreateUpdateSqlRoleAssignment then polls until it's completed
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleAssignmentThenPoll(ctx context.Context, id SqlRoleAssignmentId, input SqlRoleAssignmentCreateUpdateParameters) error {
	result, err := c.SqlResourcesCreateUpdateSqlRoleAssignment(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlResourcesCreateUpdateSqlRoleAssignment: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlResourcesCreateUpdateSqlRoleAssignment: %+v", err)
	}

	return nil
}

// preparerForSqlResourcesCreateUpdateSqlRoleAssignment prepares the SqlResourcesCreateUpdateSqlRoleAssignment request.
func (c RbacsClient) preparerForSqlResourcesCreateUpdateSqlRoleAssignment(ctx context.Context, id SqlRoleAssignmentId, input SqlRoleAssignmentCreateUpdateParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSqlResourcesCreateUpdateSqlRoleAssignment sends the SqlResourcesCreateUpdateSqlRoleAssignment request. The method will close the
// http.Response Body if it receives an error.
func (c RbacsClient) senderForSqlResourcesCreateUpdateSqlRoleAssignment(ctx context.Context, req *http.Request) (future SqlResourcesCreateUpdateSqlRoleAssignmentOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
