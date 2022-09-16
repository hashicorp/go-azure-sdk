package sqlpoolsworkloadclassifiers

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

type SqlPoolWorkloadClassifierDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlPoolWorkloadClassifierDelete ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierDelete(ctx context.Context, id WorkloadClassifierId) (result SqlPoolWorkloadClassifierDeleteOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadClassifierDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlPoolWorkloadClassifierDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlPoolWorkloadClassifierDeleteThenPoll performs SqlPoolWorkloadClassifierDelete then polls until it's completed
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierDeleteThenPoll(ctx context.Context, id WorkloadClassifierId) error {
	result, err := c.SqlPoolWorkloadClassifierDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SqlPoolWorkloadClassifierDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlPoolWorkloadClassifierDelete: %+v", err)
	}

	return nil
}

// preparerForSqlPoolWorkloadClassifierDelete prepares the SqlPoolWorkloadClassifierDelete request.
func (c SqlPoolsWorkloadClassifiersClient) preparerForSqlPoolWorkloadClassifierDelete(ctx context.Context, id WorkloadClassifierId) (*http.Request, error) {
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

// senderForSqlPoolWorkloadClassifierDelete sends the SqlPoolWorkloadClassifierDelete request. The method will close the
// http.Response Body if it receives an error.
func (c SqlPoolsWorkloadClassifiersClient) senderForSqlPoolWorkloadClassifierDelete(ctx context.Context, req *http.Request) (future SqlPoolWorkloadClassifierDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
