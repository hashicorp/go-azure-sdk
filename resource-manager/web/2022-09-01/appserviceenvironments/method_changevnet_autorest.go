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

type ChangeVnetOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *[]Site

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ChangeVnetOperationResponse, error)
}

type ChangeVnetCompleteResult struct {
	Items []Site
}

func (r ChangeVnetOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ChangeVnetOperationResponse) LoadMore(ctx context.Context) (resp ChangeVnetOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ChangeVnet ...
func (c AppServiceEnvironmentsClient) ChangeVnet(ctx context.Context, id HostingEnvironmentId, input VirtualNetworkProfile) (result ChangeVnetOperationResponse, err error) {
	req, err := c.preparerForChangeVnet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ChangeVnet", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForChangeVnet(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ChangeVnet", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ChangeVnetThenPoll performs ChangeVnet then polls until it's completed
func (c AppServiceEnvironmentsClient) ChangeVnetThenPoll(ctx context.Context, id HostingEnvironmentId, input VirtualNetworkProfile) error {
	result, err := c.ChangeVnet(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ChangeVnet: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ChangeVnet: %+v", err)
	}

	return nil
}

// preparerForChangeVnet prepares the ChangeVnet request.
func (c AppServiceEnvironmentsClient) preparerForChangeVnet(ctx context.Context, id HostingEnvironmentId, input VirtualNetworkProfile) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/changeVirtualNetwork", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForChangeVnetWithNextLink prepares the ChangeVnet request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForChangeVnetWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// senderForChangeVnet sends the ChangeVnet request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceEnvironmentsClient) senderForChangeVnet(ctx context.Context, req *http.Request) (future ChangeVnetOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
