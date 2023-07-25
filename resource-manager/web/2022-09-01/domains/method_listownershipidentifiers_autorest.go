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

type ListOwnershipIdentifiersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DomainOwnershipIdentifier

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListOwnershipIdentifiersOperationResponse, error)
}

type ListOwnershipIdentifiersCompleteResult struct {
	Items []DomainOwnershipIdentifier
}

func (r ListOwnershipIdentifiersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListOwnershipIdentifiersOperationResponse) LoadMore(ctx context.Context) (resp ListOwnershipIdentifiersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListOwnershipIdentifiers ...
func (c DomainsClient) ListOwnershipIdentifiers(ctx context.Context, id DomainId) (resp ListOwnershipIdentifiersOperationResponse, err error) {
	req, err := c.preparerForListOwnershipIdentifiers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListOwnershipIdentifiers(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListOwnershipIdentifiers prepares the ListOwnershipIdentifiers request.
func (c DomainsClient) preparerForListOwnershipIdentifiers(ctx context.Context, id DomainId) (*http.Request, error) {
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

// preparerForListOwnershipIdentifiersWithNextLink prepares the ListOwnershipIdentifiers request with the given nextLink token.
func (c DomainsClient) preparerForListOwnershipIdentifiersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListOwnershipIdentifiers handles the response to the ListOwnershipIdentifiers request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForListOwnershipIdentifiers(resp *http.Response) (result ListOwnershipIdentifiersOperationResponse, err error) {
	type page struct {
		Values   []DomainOwnershipIdentifier `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListOwnershipIdentifiersOperationResponse, err error) {
			req, err := c.preparerForListOwnershipIdentifiersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListOwnershipIdentifiers(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "domains.DomainsClient", "ListOwnershipIdentifiers", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListOwnershipIdentifiersComplete retrieves all of the results into a single object
func (c DomainsClient) ListOwnershipIdentifiersComplete(ctx context.Context, id DomainId) (ListOwnershipIdentifiersCompleteResult, error) {
	return c.ListOwnershipIdentifiersCompleteMatchingPredicate(ctx, id, DomainOwnershipIdentifierOperationPredicate{})
}

// ListOwnershipIdentifiersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DomainsClient) ListOwnershipIdentifiersCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate DomainOwnershipIdentifierOperationPredicate) (resp ListOwnershipIdentifiersCompleteResult, err error) {
	items := make([]DomainOwnershipIdentifier, 0)

	page, err := c.ListOwnershipIdentifiers(ctx, id)
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

	out := ListOwnershipIdentifiersCompleteResult{
		Items: items,
	}
	return out, nil
}
