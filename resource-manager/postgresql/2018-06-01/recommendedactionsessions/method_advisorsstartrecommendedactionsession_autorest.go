package recommendedactionsessions

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

type AdvisorsStartRecommendedActionSessionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type AdvisorsStartRecommendedActionSessionOperationOptions struct {
	DatabaseName *string
}

func DefaultAdvisorsStartRecommendedActionSessionOperationOptions() AdvisorsStartRecommendedActionSessionOperationOptions {
	return AdvisorsStartRecommendedActionSessionOperationOptions{}
}

func (o AdvisorsStartRecommendedActionSessionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AdvisorsStartRecommendedActionSessionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DatabaseName != nil {
		out["databaseName"] = *o.DatabaseName
	}

	return out
}

// AdvisorsStartRecommendedActionSession ...
func (c RecommendedActionSessionsClient) AdvisorsStartRecommendedActionSession(ctx context.Context, id AdvisorId, options AdvisorsStartRecommendedActionSessionOperationOptions) (result AdvisorsStartRecommendedActionSessionOperationResponse, err error) {
	req, err := c.preparerForAdvisorsStartRecommendedActionSession(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendedactionsessions.RecommendedActionSessionsClient", "AdvisorsStartRecommendedActionSession", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAdvisorsStartRecommendedActionSession(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendedactionsessions.RecommendedActionSessionsClient", "AdvisorsStartRecommendedActionSession", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AdvisorsStartRecommendedActionSessionThenPoll performs AdvisorsStartRecommendedActionSession then polls until it's completed
func (c RecommendedActionSessionsClient) AdvisorsStartRecommendedActionSessionThenPoll(ctx context.Context, id AdvisorId, options AdvisorsStartRecommendedActionSessionOperationOptions) error {
	result, err := c.AdvisorsStartRecommendedActionSession(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing AdvisorsStartRecommendedActionSession: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AdvisorsStartRecommendedActionSession: %+v", err)
	}

	return nil
}

// preparerForAdvisorsStartRecommendedActionSession prepares the AdvisorsStartRecommendedActionSession request.
func (c RecommendedActionSessionsClient) preparerForAdvisorsStartRecommendedActionSession(ctx context.Context, id AdvisorId, options AdvisorsStartRecommendedActionSessionOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/recommendedActionSessions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAdvisorsStartRecommendedActionSession sends the AdvisorsStartRecommendedActionSession request. The method will close the
// http.Response Body if it receives an error.
func (c RecommendedActionSessionsClient) senderForAdvisorsStartRecommendedActionSession(ctx context.Context, req *http.Request) (future AdvisorsStartRecommendedActionSessionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
