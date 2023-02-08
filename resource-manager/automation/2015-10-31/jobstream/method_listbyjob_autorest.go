package jobstream

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

type ListByJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]JobStream

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByJobOperationResponse, error)
}

type ListByJobCompleteResult struct {
	Items []JobStream
}

func (r ListByJobOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByJobOperationResponse) LoadMore(ctx context.Context) (resp ListByJobOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByJobOperationOptions struct {
	Filter *string
}

func DefaultListByJobOperationOptions() ListByJobOperationOptions {
	return ListByJobOperationOptions{}
}

func (o ListByJobOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByJobOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByJob ...
func (c JobStreamClient) ListByJob(ctx context.Context, id JobId, options ListByJobOperationOptions) (resp ListByJobOperationResponse, err error) {
	req, err := c.preparerForListByJob(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByJob(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByJob prepares the ListByJob request.
func (c JobStreamClient) preparerForListByJob(ctx context.Context, id JobId, options ListByJobOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/streams", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByJobWithNextLink prepares the ListByJob request with the given nextLink token.
func (c JobStreamClient) preparerForListByJobWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByJob handles the response to the ListByJob request. The method always
// closes the http.Response Body.
func (c JobStreamClient) responderForListByJob(resp *http.Response) (result ListByJobOperationResponse, err error) {
	type page struct {
		Values   []JobStream `json:"value"`
		NextLink *string     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByJobOperationResponse, err error) {
			req, err := c.preparerForListByJobWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByJob(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobstream.JobStreamClient", "ListByJob", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByJobComplete retrieves all of the results into a single object
func (c JobStreamClient) ListByJobComplete(ctx context.Context, id JobId, options ListByJobOperationOptions) (ListByJobCompleteResult, error) {
	return c.ListByJobCompleteMatchingPredicate(ctx, id, options, JobStreamOperationPredicate{})
}

// ListByJobCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c JobStreamClient) ListByJobCompleteMatchingPredicate(ctx context.Context, id JobId, options ListByJobOperationOptions, predicate JobStreamOperationPredicate) (resp ListByJobCompleteResult, err error) {
	items := make([]JobStream, 0)

	page, err := c.ListByJob(ctx, id, options)
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

	out := ListByJobCompleteResult{
		Items: items,
	}
	return out, nil
}
