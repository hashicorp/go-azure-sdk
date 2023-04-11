package adaptivenetworkhardenings

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

type ListByExtendedResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AdaptiveNetworkHardening

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByExtendedResourceOperationResponse, error)
}

type ListByExtendedResourceCompleteResult struct {
	Items []AdaptiveNetworkHardening
}

func (r ListByExtendedResourceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByExtendedResourceOperationResponse) LoadMore(ctx context.Context) (resp ListByExtendedResourceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByExtendedResource ...
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResource(ctx context.Context, id ProviderId) (resp ListByExtendedResourceOperationResponse, err error) {
	req, err := c.preparerForListByExtendedResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByExtendedResource(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByExtendedResource prepares the ListByExtendedResource request.
func (c AdaptiveNetworkHardeningsClient) preparerForListByExtendedResource(ctx context.Context, id ProviderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/adaptiveNetworkHardenings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByExtendedResourceWithNextLink prepares the ListByExtendedResource request with the given nextLink token.
func (c AdaptiveNetworkHardeningsClient) preparerForListByExtendedResourceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByExtendedResource handles the response to the ListByExtendedResource request. The method always
// closes the http.Response Body.
func (c AdaptiveNetworkHardeningsClient) responderForListByExtendedResource(resp *http.Response) (result ListByExtendedResourceOperationResponse, err error) {
	type page struct {
		Values   []AdaptiveNetworkHardening `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByExtendedResourceOperationResponse, err error) {
			req, err := c.preparerForListByExtendedResourceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByExtendedResource(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "ListByExtendedResource", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByExtendedResourceComplete retrieves all of the results into a single object
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResourceComplete(ctx context.Context, id ProviderId) (ListByExtendedResourceCompleteResult, error) {
	return c.ListByExtendedResourceCompleteMatchingPredicate(ctx, id, AdaptiveNetworkHardeningOperationPredicate{})
}

// ListByExtendedResourceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AdaptiveNetworkHardeningsClient) ListByExtendedResourceCompleteMatchingPredicate(ctx context.Context, id ProviderId, predicate AdaptiveNetworkHardeningOperationPredicate) (resp ListByExtendedResourceCompleteResult, err error) {
	items := make([]AdaptiveNetworkHardening, 0)

	page, err := c.ListByExtendedResource(ctx, id)
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

	out := ListByExtendedResourceCompleteResult{
		Items: items,
	}
	return out, nil
}
