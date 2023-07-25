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

type ListPremierAddOnOffersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PremierAddOnOffer

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPremierAddOnOffersOperationResponse, error)
}

type ListPremierAddOnOffersCompleteResult struct {
	Items []PremierAddOnOffer
}

func (r ListPremierAddOnOffersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPremierAddOnOffersOperationResponse) LoadMore(ctx context.Context) (resp ListPremierAddOnOffersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListPremierAddOnOffers ...
func (c ResourceProvidersClient) ListPremierAddOnOffers(ctx context.Context, id commonids.SubscriptionId) (resp ListPremierAddOnOffersOperationResponse, err error) {
	req, err := c.preparerForListPremierAddOnOffers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPremierAddOnOffers(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPremierAddOnOffers prepares the ListPremierAddOnOffers request.
func (c ResourceProvidersClient) preparerForListPremierAddOnOffers(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/premieraddonoffers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListPremierAddOnOffersWithNextLink prepares the ListPremierAddOnOffers request with the given nextLink token.
func (c ResourceProvidersClient) preparerForListPremierAddOnOffersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPremierAddOnOffers handles the response to the ListPremierAddOnOffers request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForListPremierAddOnOffers(resp *http.Response) (result ListPremierAddOnOffersOperationResponse, err error) {
	type page struct {
		Values   []PremierAddOnOffer `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPremierAddOnOffersOperationResponse, err error) {
			req, err := c.preparerForListPremierAddOnOffersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPremierAddOnOffers(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListPremierAddOnOffers", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPremierAddOnOffersComplete retrieves all of the results into a single object
func (c ResourceProvidersClient) ListPremierAddOnOffersComplete(ctx context.Context, id commonids.SubscriptionId) (ListPremierAddOnOffersCompleteResult, error) {
	return c.ListPremierAddOnOffersCompleteMatchingPredicate(ctx, id, PremierAddOnOfferOperationPredicate{})
}

// ListPremierAddOnOffersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceProvidersClient) ListPremierAddOnOffersCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate PremierAddOnOfferOperationPredicate) (resp ListPremierAddOnOffersCompleteResult, err error) {
	items := make([]PremierAddOnOffer, 0)

	page, err := c.ListPremierAddOnOffers(ctx, id)
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

	out := ListPremierAddOnOffersCompleteResult{
		Items: items,
	}
	return out, nil
}
