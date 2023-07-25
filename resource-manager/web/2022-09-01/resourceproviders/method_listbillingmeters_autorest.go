package resourceproviders

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBillingMetersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BillingMeter

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBillingMetersOperationResponse, error)
}

type ListBillingMetersCompleteResult struct {
	Items []BillingMeter
}

func (r ListBillingMetersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBillingMetersOperationResponse) LoadMore(ctx context.Context) (resp ListBillingMetersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListBillingMetersOperationOptions struct {
	BillingLocation *string
	OsType          *string
}

func DefaultListBillingMetersOperationOptions() ListBillingMetersOperationOptions {
	return ListBillingMetersOperationOptions{}
}

func (o ListBillingMetersOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListBillingMetersOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.BillingLocation != nil {
		out["billingLocation"] = *o.BillingLocation
	}

	if o.OsType != nil {
		out["osType"] = *o.OsType
	}

	return out
}

// ListBillingMeters ...
func (c ResourceProvidersClient) ListBillingMeters(ctx context.Context, id commonids.SubscriptionId, options ListBillingMetersOperationOptions) (resp ListBillingMetersOperationResponse, err error) {
	req, err := c.preparerForListBillingMeters(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBillingMeters(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBillingMeters prepares the ListBillingMeters request.
func (c ResourceProvidersClient) preparerForListBillingMeters(ctx context.Context, id commonids.SubscriptionId, options ListBillingMetersOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/billingMeters", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBillingMetersWithNextLink prepares the ListBillingMeters request with the given nextLink token.
func (c ResourceProvidersClient) preparerForListBillingMetersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBillingMeters handles the response to the ListBillingMeters request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForListBillingMeters(resp *http.Response) (result ListBillingMetersOperationResponse, err error) {
	type page struct {
		Values   []BillingMeter `json:"value"`
		NextLink *string        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBillingMetersOperationResponse, err error) {
			req, err := c.preparerForListBillingMetersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBillingMeters(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListBillingMeters", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBillingMetersComplete retrieves all of the results into a single object
func (c ResourceProvidersClient) ListBillingMetersComplete(ctx context.Context, id commonids.SubscriptionId, options ListBillingMetersOperationOptions) (ListBillingMetersCompleteResult, error) {
	return c.ListBillingMetersCompleteMatchingPredicate(ctx, id, options, BillingMeterOperationPredicate{})
}

// ListBillingMetersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceProvidersClient) ListBillingMetersCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ListBillingMetersOperationOptions, predicate BillingMeterOperationPredicate) (resp ListBillingMetersCompleteResult, err error) {
	items := make([]BillingMeter, 0)

	page, err := c.ListBillingMeters(ctx, id, options)
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

	out := ListBillingMetersCompleteResult{
		Items: items,
	}
	return out, nil
}
