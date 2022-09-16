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

type SqlPoolWorkloadClassifierCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SqlPoolWorkloadClassifierCreateOrUpdate ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierCreateOrUpdate(ctx context.Context, id WorkloadClassifierId, input WorkloadClassifier) (result SqlPoolWorkloadClassifierCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadClassifierCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlPoolWorkloadClassifierCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlPoolWorkloadClassifierCreateOrUpdateThenPoll performs SqlPoolWorkloadClassifierCreateOrUpdate then polls until it's completed
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierCreateOrUpdateThenPoll(ctx context.Context, id WorkloadClassifierId, input WorkloadClassifier) error {
	result, err := c.SqlPoolWorkloadClassifierCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlPoolWorkloadClassifierCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlPoolWorkloadClassifierCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForSqlPoolWorkloadClassifierCreateOrUpdate prepares the SqlPoolWorkloadClassifierCreateOrUpdate request.
func (c SqlPoolsWorkloadClassifiersClient) preparerForSqlPoolWorkloadClassifierCreateOrUpdate(ctx context.Context, id WorkloadClassifierId, input WorkloadClassifier) (*http.Request, error) {
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

// senderForSqlPoolWorkloadClassifierCreateOrUpdate sends the SqlPoolWorkloadClassifierCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c SqlPoolsWorkloadClassifiersClient) senderForSqlPoolWorkloadClassifierCreateOrUpdate(ctx context.Context, req *http.Request) (future SqlPoolWorkloadClassifierCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
