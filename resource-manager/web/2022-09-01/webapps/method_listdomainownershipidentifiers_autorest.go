package webapps

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

type ListDomainOwnershipIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Identifier

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDomainOwnershipIdentifiersOperationResponse, error)
}

type ListDomainOwnershipIdentifiersCompleteResult struct {
	Items []Identifier
}

func (r ListDomainOwnershipIdentifiersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDomainOwnershipIdentifiersOperationResponse) LoadMore(ctx context.Context) (resp ListDomainOwnershipIdentifiersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDomainOwnershipIdentifiers ...
func (c WebAppsClient) ListDomainOwnershipIdentifiers(ctx context.Context, id SiteId) (resp ListDomainOwnershipIdentifiersOperationResponse, err error) {
	req, err := c.preparerForListDomainOwnershipIdentifiers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDomainOwnershipIdentifiers(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDomainOwnershipIdentifiers prepares the ListDomainOwnershipIdentifiers request.
func (c WebAppsClient) preparerForListDomainOwnershipIdentifiers(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/domainOwnershipIdentifiers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDomainOwnershipIdentifiersWithNextLink prepares the ListDomainOwnershipIdentifiers request with the given nextLink token.
func (c WebAppsClient) preparerForListDomainOwnershipIdentifiersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListDomainOwnershipIdentifiers handles the response to the ListDomainOwnershipIdentifiers request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListDomainOwnershipIdentifiers(resp *http.Response) (result ListDomainOwnershipIdentifiersOperationResponse, err error) {
	type page struct {
		Values   []Identifier `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDomainOwnershipIdentifiersOperationResponse, err error) {
			req, err := c.preparerForListDomainOwnershipIdentifiersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDomainOwnershipIdentifiers(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDomainOwnershipIdentifiers", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDomainOwnershipIdentifiersComplete retrieves all of the results into a single object
func (c WebAppsClient) ListDomainOwnershipIdentifiersComplete(ctx context.Context, id SiteId) (ListDomainOwnershipIdentifiersCompleteResult, error) {
	return c.ListDomainOwnershipIdentifiersCompleteMatchingPredicate(ctx, id, IdentifierOperationPredicate{})
}

// ListDomainOwnershipIdentifiersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListDomainOwnershipIdentifiersCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate IdentifierOperationPredicate) (resp ListDomainOwnershipIdentifiersCompleteResult, err error) {
	items := make([]Identifier, 0)

	page, err := c.ListDomainOwnershipIdentifiers(ctx, id)
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

	out := ListDomainOwnershipIdentifiersCompleteResult{
		Items: items,
	}
	return out, nil
}
