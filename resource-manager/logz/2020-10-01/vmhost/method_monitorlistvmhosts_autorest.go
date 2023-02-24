package vmhost

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorListVMHostsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMResources

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (MonitorListVMHostsOperationResponse, error)
}

type MonitorListVMHostsCompleteResult struct {
	Items []VMResources
}

func (r MonitorListVMHostsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r MonitorListVMHostsOperationResponse) LoadMore(ctx context.Context) (resp MonitorListVMHostsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// MonitorListVMHosts ...
func (c VMHostClient) MonitorListVMHosts(ctx context.Context, id MonitorId) (resp MonitorListVMHostsOperationResponse, err error) {
	req, err := c.preparerForMonitorListVMHosts(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForMonitorListVMHosts(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForMonitorListVMHosts prepares the MonitorListVMHosts request.
func (c VMHostClient) preparerForMonitorListVMHosts(ctx context.Context, id MonitorId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listVMHosts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForMonitorListVMHostsWithNextLink prepares the MonitorListVMHosts request with the given nextLink token.
func (c VMHostClient) preparerForMonitorListVMHostsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForMonitorListVMHosts handles the response to the MonitorListVMHosts request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForMonitorListVMHosts(resp *http.Response) (result MonitorListVMHostsOperationResponse, err error) {
	type page struct {
		Values   []VMResources `json:"value"`
		NextLink *string       `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result MonitorListVMHostsOperationResponse, err error) {
			req, err := c.preparerForMonitorListVMHostsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForMonitorListVMHosts(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHosts", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// MonitorListVMHostsComplete retrieves all of the results into a single object
func (c VMHostClient) MonitorListVMHostsComplete(ctx context.Context, id MonitorId) (MonitorListVMHostsCompleteResult, error) {
	return c.MonitorListVMHostsCompleteMatchingPredicate(ctx, id, VMResourcesOperationPredicate{})
}

// MonitorListVMHostsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VMHostClient) MonitorListVMHostsCompleteMatchingPredicate(ctx context.Context, id MonitorId, predicate VMResourcesOperationPredicate) (resp MonitorListVMHostsCompleteResult, err error) {
	items := make([]VMResources, 0)

	page, err := c.MonitorListVMHosts(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := MonitorListVMHostsCompleteResult{
		Items: items,
	}
	return out, nil
}
