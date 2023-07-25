package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SuspendOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *[]Site

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SuspendOperationResponse, error)
}

type SuspendCompleteResult struct {
	Items []Site
}

func (r SuspendOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SuspendOperationResponse) LoadMore(ctx context.Context) (resp SuspendOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// Suspend ...
func (c AppServiceEnvironmentsClient) Suspend(ctx context.Context, id HostingEnvironmentId) (result SuspendOperationResponse, err error) {
	req, err := c.preparerForSuspend(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Suspend", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSuspend(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Suspend", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SuspendThenPoll performs Suspend then polls until it's completed
func (c AppServiceEnvironmentsClient) SuspendThenPoll(ctx context.Context, id HostingEnvironmentId) error {
	result, err := c.Suspend(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Suspend: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Suspend: %+v", err)
	}

	return nil
}

// preparerForSuspend prepares the Suspend request.
func (c AppServiceEnvironmentsClient) preparerForSuspend(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/suspend", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSuspendWithNextLink prepares the Suspend request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForSuspendWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSuspend sends the Suspend request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceEnvironmentsClient) senderForSuspend(ctx context.Context, req *http.Request) (future SuspendOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
