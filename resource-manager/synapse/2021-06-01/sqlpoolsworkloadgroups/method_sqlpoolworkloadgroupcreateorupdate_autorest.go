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

type SqlPoolWorkloadGroupCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlPoolWorkloadGroupCreateOrUpdate ...
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupCreateOrUpdate(ctx context.Context, id WorkloadGroupId, input WorkloadGroup) (result SqlPoolWorkloadGroupCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadGroupCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlPoolWorkloadGroupCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlPoolWorkloadGroupCreateOrUpdateThenPoll performs SqlPoolWorkloadGroupCreateOrUpdate then polls until it's completed
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupCreateOrUpdateThenPoll(ctx context.Context, id WorkloadGroupId, input WorkloadGroup) error {
	result, err := c.SqlPoolWorkloadGroupCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlPoolWorkloadGroupCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlPoolWorkloadGroupCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForSqlPoolWorkloadGroupCreateOrUpdate prepares the SqlPoolWorkloadGroupCreateOrUpdate request.
func (c SqlPoolsWorkloadGroupsClient) preparerForSqlPoolWorkloadGroupCreateOrUpdate(ctx context.Context, id WorkloadGroupId, input WorkloadGroup) (*http.Request, error) {
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

// senderForSqlPoolWorkloadGroupCreateOrUpdate sends the SqlPoolWorkloadGroupCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c SqlPoolsWorkloadGroupsClient) senderForSqlPoolWorkloadGroupCreateOrUpdate(ctx context.Context, req *http.Request) (future SqlPoolWorkloadGroupCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
