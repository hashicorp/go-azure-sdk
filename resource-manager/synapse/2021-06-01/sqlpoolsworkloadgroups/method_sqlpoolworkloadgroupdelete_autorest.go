package sqlpoolsworkloadgroups

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

type SqlPoolWorkloadGroupDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlPoolWorkloadGroupDelete ...
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupDelete(ctx context.Context, id WorkloadGroupId) (result SqlPoolWorkloadGroupDeleteOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadGroupDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlPoolWorkloadGroupDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlPoolWorkloadGroupDeleteThenPoll performs SqlPoolWorkloadGroupDelete then polls until it's completed
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupDeleteThenPoll(ctx context.Context, id WorkloadGroupId) error {
	result, err := c.SqlPoolWorkloadGroupDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SqlPoolWorkloadGroupDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlPoolWorkloadGroupDelete: %+v", err)
	}

	return nil
}

// preparerForSqlPoolWorkloadGroupDelete prepares the SqlPoolWorkloadGroupDelete request.
func (c SqlPoolsWorkloadGroupsClient) preparerForSqlPoolWorkloadGroupDelete(ctx context.Context, id WorkloadGroupId) (*http.Request, error) {
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

// senderForSqlPoolWorkloadGroupDelete sends the SqlPoolWorkloadGroupDelete request. The method will close the
// http.Response Body if it receives an error.
func (c SqlPoolsWorkloadGroupsClient) senderForSqlPoolWorkloadGroupDelete(ctx context.Context, req *http.Request) (future SqlPoolWorkloadGroupDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
