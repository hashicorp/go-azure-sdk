package sharesubscription

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

type ConsumerSourceDataSetsListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ConsumerSourceDataSet

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, error)
}

type ConsumerSourceDataSetsListByShareSubscriptionCompleteResult struct {
	Items []ConsumerSourceDataSet
}

func (r ConsumerSourceDataSetsListByShareSubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ConsumerSourceDataSetsListByShareSubscriptionOperationResponse) LoadMore(ctx context.Context) (resp ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ConsumerSourceDataSetsListByShareSubscription ...
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (resp ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
	req, err := c.preparerForConsumerSourceDataSetsListByShareSubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForConsumerSourceDataSetsListByShareSubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForConsumerSourceDataSetsListByShareSubscription prepares the ConsumerSourceDataSetsListByShareSubscription request.
func (c ShareSubscriptionClient) preparerForConsumerSourceDataSetsListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/consumerSourceDataSets", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForConsumerSourceDataSetsListByShareSubscriptionWithNextLink prepares the ConsumerSourceDataSetsListByShareSubscription request with the given nextLink token.
func (c ShareSubscriptionClient) preparerForConsumerSourceDataSetsListByShareSubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForConsumerSourceDataSetsListByShareSubscription handles the response to the ConsumerSourceDataSetsListByShareSubscription request. The method always
// closes the http.Response Body.
func (c ShareSubscriptionClient) responderForConsumerSourceDataSetsListByShareSubscription(resp *http.Response) (result ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
	type page struct {
		Values   []ConsumerSourceDataSet `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ConsumerSourceDataSetsListByShareSubscriptionOperationResponse, err error) {
			req, err := c.preparerForConsumerSourceDataSetsListByShareSubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForConsumerSourceDataSetsListByShareSubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sharesubscription.ShareSubscriptionClient", "ConsumerSourceDataSetsListByShareSubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ConsumerSourceDataSetsListByShareSubscriptionComplete retrieves all of the results into a single object
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId) (ConsumerSourceDataSetsListByShareSubscriptionCompleteResult, error) {
	return c.ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate(ctx, id, ConsumerSourceDataSetOperationPredicate{})
}

// ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ShareSubscriptionClient) ConsumerSourceDataSetsListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate ConsumerSourceDataSetOperationPredicate) (resp ConsumerSourceDataSetsListByShareSubscriptionCompleteResult, err error) {
	items := make([]ConsumerSourceDataSet, 0)

	page, err := c.ConsumerSourceDataSetsListByShareSubscription(ctx, id)
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

	out := ConsumerSourceDataSetsListByShareSubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
