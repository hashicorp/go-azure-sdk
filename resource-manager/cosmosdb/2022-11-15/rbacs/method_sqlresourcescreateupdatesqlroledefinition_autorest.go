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

type SqlResourcesCreateUpdateSqlRoleDefinitionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *SqlRoleDefinitionGetResults
}

// SqlResourcesCreateUpdateSqlRoleDefinition ...
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleDefinition(ctx context.Context, id SqlRoleDefinitionId, input SqlRoleDefinitionCreateUpdateParameters) (result SqlResourcesCreateUpdateSqlRoleDefinitionOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesCreateUpdateSqlRoleDefinition(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesCreateUpdateSqlRoleDefinition", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlResourcesCreateUpdateSqlRoleDefinition(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rbacs.RbacsClient", "SqlResourcesCreateUpdateSqlRoleDefinition", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlResourcesCreateUpdateSqlRoleDefinitionThenPoll performs SqlResourcesCreateUpdateSqlRoleDefinition then polls until it's completed
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleDefinitionThenPoll(ctx context.Context, id SqlRoleDefinitionId, input SqlRoleDefinitionCreateUpdateParameters) error {
	result, err := c.SqlResourcesCreateUpdateSqlRoleDefinition(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlResourcesCreateUpdateSqlRoleDefinition: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlResourcesCreateUpdateSqlRoleDefinition: %+v", err)
	}

	return nil
}

// preparerForSqlResourcesCreateUpdateSqlRoleDefinition prepares the SqlResourcesCreateUpdateSqlRoleDefinition request.
func (c RbacsClient) preparerForSqlResourcesCreateUpdateSqlRoleDefinition(ctx context.Context, id SqlRoleDefinitionId, input SqlRoleDefinitionCreateUpdateParameters) (*http.Request, error) {
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

// senderForSqlResourcesCreateUpdateSqlRoleDefinition sends the SqlResourcesCreateUpdateSqlRoleDefinition request. The method will close the
// http.Response Body if it receives an error.
func (c RbacsClient) senderForSqlResourcesCreateUpdateSqlRoleDefinition(ctx context.Context, req *http.Request) (future SqlResourcesCreateUpdateSqlRoleDefinitionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
