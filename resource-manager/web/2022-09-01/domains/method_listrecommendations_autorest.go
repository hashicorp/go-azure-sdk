package domains

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

type ListRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NameIdentifier

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListRecommendationsOperationResponse, error)
}

type ListRecommendationsCompleteResult struct {
	Items []NameIdentifier
}

func (r ListRecommendationsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListRecommendationsOperationResponse) LoadMore(ctx context.Context) (resp ListRecommendationsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListRecommendations ...
func (c DomainsClient) ListRecommendations(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters) (resp ListRecommendationsOperationResponse, err error) {
	req, err := c.preparerForListRecommendations(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListRecommendations(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListRecommendations prepares the ListRecommendations request.
func (c DomainsClient) preparerForListRecommendations(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.DomainRegistration/listDomainRecommendations", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListRecommendationsWithNextLink prepares the ListRecommendations request with the given nextLink token.
func (c DomainsClient) preparerForListRecommendationsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListRecommendations handles the response to the ListRecommendations request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForListRecommendations(resp *http.Response) (result ListRecommendationsOperationResponse, err error) {
	type page struct {
		Values   []NameIdentifier `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListRecommendationsOperationResponse, err error) {
			req, err := c.preparerForListRecommendationsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListRecommendations(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListRecommendations", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListRecommendationsComplete retrieves all of the results into a single object
func (c DomainsClient) ListRecommendationsComplete(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters) (ListRecommendationsCompleteResult, error) {
	return c.ListRecommendationsCompleteMatchingPredicate(ctx, id, input, NameIdentifierOperationPredicate{})
}

// ListRecommendationsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DomainsClient) ListRecommendationsCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters, predicate NameIdentifierOperationPredicate) (resp ListRecommendationsCompleteResult, err error) {
	items := make([]NameIdentifier, 0)

	page, err := c.ListRecommendations(ctx, id, input)
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

	out := ListRecommendationsCompleteResult{
		Items: items,
	}
	return out, nil
}
