package domains

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

type ListByEmailServiceResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DomainResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByEmailServiceResourceOperationResponse, error)
}

type ListByEmailServiceResourceCompleteResult struct {
	Items []DomainResource
}

func (r ListByEmailServiceResourceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByEmailServiceResourceOperationResponse) LoadMore(ctx context.Context) (resp ListByEmailServiceResourceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByEmailServiceResource ...
func (c DomainsClient) ListByEmailServiceResource(ctx context.Context, id EmailServiceId) (resp ListByEmailServiceResourceOperationResponse, err error) {
	req, err := c.preparerForListByEmailServiceResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByEmailServiceResource(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByEmailServiceResource prepares the ListByEmailServiceResource request.
func (c DomainsClient) preparerForListByEmailServiceResource(ctx context.Context, id EmailServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/domains", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByEmailServiceResourceWithNextLink prepares the ListByEmailServiceResource request with the given nextLink token.
func (c DomainsClient) preparerForListByEmailServiceResourceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByEmailServiceResource handles the response to the ListByEmailServiceResource request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForListByEmailServiceResource(resp *http.Response) (result ListByEmailServiceResourceOperationResponse, err error) {
	type page struct {
		Values   []DomainResource `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByEmailServiceResourceOperationResponse, err error) {
			req, err := c.preparerForListByEmailServiceResourceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByEmailServiceResource(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListByEmailServiceResource", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByEmailServiceResourceComplete retrieves all of the results into a single object
func (c DomainsClient) ListByEmailServiceResourceComplete(ctx context.Context, id EmailServiceId) (ListByEmailServiceResourceCompleteResult, error) {
	return c.ListByEmailServiceResourceCompleteMatchingPredicate(ctx, id, DomainResourceOperationPredicate{})
}

// ListByEmailServiceResourceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DomainsClient) ListByEmailServiceResourceCompleteMatchingPredicate(ctx context.Context, id EmailServiceId, predicate DomainResourceOperationPredicate) (resp ListByEmailServiceResourceCompleteResult, err error) {
	items := make([]DomainResource, 0)

	page, err := c.ListByEmailServiceResource(ctx, id)
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

	out := ListByEmailServiceResourceCompleteResult{
		Items: items,
	}
	return out, nil
}
