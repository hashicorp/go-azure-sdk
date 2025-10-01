package bms

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

type FetchTieringCostPostOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *TieringCostInfo
}

// FetchTieringCostPost ...
func (c BmsClient) FetchTieringCostPost(ctx context.Context, id VaultId, input FetchTieringCostInfoRequest) (result FetchTieringCostPostOperationResponse, err error) {
	req, err := c.preparerForFetchTieringCostPost(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "FetchTieringCostPost", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForFetchTieringCostPost(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "FetchTieringCostPost", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// FetchTieringCostPostThenPoll performs FetchTieringCostPost then polls until it's completed
func (c BmsClient) FetchTieringCostPostThenPoll(ctx context.Context, id VaultId, input FetchTieringCostInfoRequest) error {
	result, err := c.FetchTieringCostPost(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing FetchTieringCostPost: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after FetchTieringCostPost: %+v", err)
	}

	return nil
}

// preparerForFetchTieringCostPost prepares the FetchTieringCostPost request.
func (c BmsClient) preparerForFetchTieringCostPost(ctx context.Context, id VaultId, input FetchTieringCostInfoRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupTieringCost/default/fetchTieringCost", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForFetchTieringCostPost sends the FetchTieringCostPost request. The method will close the
// http.Response Body if it receives an error.
func (c BmsClient) senderForFetchTieringCostPost(ctx context.Context, req *http.Request) (future FetchTieringCostPostOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
