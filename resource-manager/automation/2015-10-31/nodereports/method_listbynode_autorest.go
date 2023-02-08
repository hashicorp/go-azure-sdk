package nodereports

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

type ListByNodeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DscNodeReport

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByNodeOperationResponse, error)
}

type ListByNodeCompleteResult struct {
	Items []DscNodeReport
}

func (r ListByNodeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByNodeOperationResponse) LoadMore(ctx context.Context) (resp ListByNodeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByNodeOperationOptions struct {
	Filter *string
}

func DefaultListByNodeOperationOptions() ListByNodeOperationOptions {
	return ListByNodeOperationOptions{}
}

func (o ListByNodeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByNodeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByNode ...
func (c NodeReportsClient) ListByNode(ctx context.Context, id NodeId, options ListByNodeOperationOptions) (resp ListByNodeOperationResponse, err error) {
	req, err := c.preparerForListByNode(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByNode(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByNode prepares the ListByNode request.
func (c NodeReportsClient) preparerForListByNode(ctx context.Context, id NodeId, options ListByNodeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/reports", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByNodeWithNextLink prepares the ListByNode request with the given nextLink token.
func (c NodeReportsClient) preparerForListByNodeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByNode handles the response to the ListByNode request. The method always
// closes the http.Response Body.
func (c NodeReportsClient) responderForListByNode(resp *http.Response) (result ListByNodeOperationResponse, err error) {
	type page struct {
		Values   []DscNodeReport `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByNodeOperationResponse, err error) {
			req, err := c.preparerForListByNodeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByNode(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodereports.NodeReportsClient", "ListByNode", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByNodeComplete retrieves all of the results into a single object
func (c NodeReportsClient) ListByNodeComplete(ctx context.Context, id NodeId, options ListByNodeOperationOptions) (ListByNodeCompleteResult, error) {
	return c.ListByNodeCompleteMatchingPredicate(ctx, id, options, DscNodeReportOperationPredicate{})
}

// ListByNodeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c NodeReportsClient) ListByNodeCompleteMatchingPredicate(ctx context.Context, id NodeId, options ListByNodeOperationOptions, predicate DscNodeReportOperationPredicate) (resp ListByNodeCompleteResult, err error) {
	items := make([]DscNodeReport, 0)

	page, err := c.ListByNode(ctx, id, options)
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

	out := ListByNodeCompleteResult{
		Items: items,
	}
	return out, nil
}
