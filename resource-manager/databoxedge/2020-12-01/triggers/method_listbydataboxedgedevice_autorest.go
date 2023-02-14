package triggers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Trigger

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByDataBoxEdgeDeviceOperationResponse, error)
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	Items []Trigger
}

func (r ListByDataBoxEdgeDeviceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByDataBoxEdgeDeviceOperationResponse) LoadMore(ctx context.Context) (resp ListByDataBoxEdgeDeviceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByDataBoxEdgeDeviceOperationOptions struct {
	Filter *string
}

func DefaultListByDataBoxEdgeDeviceOperationOptions() ListByDataBoxEdgeDeviceOperationOptions {
	return ListByDataBoxEdgeDeviceOperationOptions{}
}

func (o ListByDataBoxEdgeDeviceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByDataBoxEdgeDeviceOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByDataBoxEdgeDevice ...
func (c TriggersClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (resp ListByDataBoxEdgeDeviceOperationResponse, err error) {
	req, err := c.preparerForListByDataBoxEdgeDevice(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByDataBoxEdgeDevice(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByDataBoxEdgeDevice prepares the ListByDataBoxEdgeDevice request.
func (c TriggersClient) preparerForListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/triggers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByDataBoxEdgeDeviceWithNextLink prepares the ListByDataBoxEdgeDevice request with the given nextLink token.
func (c TriggersClient) preparerForListByDataBoxEdgeDeviceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByDataBoxEdgeDevice handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (c TriggersClient) responderForListByDataBoxEdgeDevice(resp *http.Response) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	type page struct {
		Values   []json.RawMessage `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	temp := make([]Trigger, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalTriggerImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for Trigger (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
			req, err := c.preparerForListByDataBoxEdgeDeviceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByDataBoxEdgeDevice(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "ListByDataBoxEdgeDevice", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all of the results into a single object
func (c TriggersClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, options, TriggerOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c TriggersClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions, predicate TriggerOperationPredicate) (resp ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]Trigger, 0)

	page, err := c.ListByDataBoxEdgeDevice(ctx, id, options)
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

	out := ListByDataBoxEdgeDeviceCompleteResult{
		Items: items,
	}
	return out, nil
}
