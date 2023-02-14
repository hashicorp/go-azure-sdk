package trigger

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

type ListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Trigger

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByShareSubscriptionOperationResponse, error)
}

type ListByShareSubscriptionCompleteResult struct {
	Items []Trigger
}

func (r ListByShareSubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByShareSubscriptionOperationResponse) LoadMore(ctx context.Context) (resp ListByShareSubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByShareSubscription ...
func (c TriggerClient) ListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (resp ListByShareSubscriptionOperationResponse, err error) {
	req, err := c.preparerForListByShareSubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByShareSubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByShareSubscription prepares the ListByShareSubscription request.
func (c TriggerClient) preparerForListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/triggers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByShareSubscriptionWithNextLink prepares the ListByShareSubscription request with the given nextLink token.
func (c TriggerClient) preparerForListByShareSubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByShareSubscription handles the response to the ListByShareSubscription request. The method always
// closes the http.Response Body.
func (c TriggerClient) responderForListByShareSubscription(resp *http.Response) (result ListByShareSubscriptionOperationResponse, err error) {
	type page struct {
		Values   []Trigger `json:"value"`
		NextLink *string   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByShareSubscriptionOperationResponse, err error) {
			req, err := c.preparerForListByShareSubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByShareSubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "trigger.TriggerClient", "ListByShareSubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByShareSubscriptionComplete retrieves all of the results into a single object
func (c TriggerClient) ListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId) (ListByShareSubscriptionCompleteResult, error) {
	return c.ListByShareSubscriptionCompleteMatchingPredicate(ctx, id, TriggerOperationPredicate{})
}

// ListByShareSubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c TriggerClient) ListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate TriggerOperationPredicate) (resp ListByShareSubscriptionCompleteResult, err error) {
	items := make([]Trigger, 0)

	page, err := c.ListByShareSubscription(ctx, id)
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

	out := ListByShareSubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
