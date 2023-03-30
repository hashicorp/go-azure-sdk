package senderusernames

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

type ListByDomainsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SenderUsernameResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByDomainsOperationResponse, error)
}

type ListByDomainsCompleteResult struct {
	Items []SenderUsernameResource
}

func (r ListByDomainsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByDomainsOperationResponse) LoadMore(ctx context.Context) (resp ListByDomainsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByDomains ...
func (c SenderUsernamesClient) ListByDomains(ctx context.Context, id DomainId) (resp ListByDomainsOperationResponse, err error) {
	req, err := c.preparerForListByDomains(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByDomains(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByDomains prepares the ListByDomains request.
func (c SenderUsernamesClient) preparerForListByDomains(ctx context.Context, id DomainId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/senderUsernames", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByDomainsWithNextLink prepares the ListByDomains request with the given nextLink token.
func (c SenderUsernamesClient) preparerForListByDomainsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByDomains handles the response to the ListByDomains request. The method always
// closes the http.Response Body.
func (c SenderUsernamesClient) responderForListByDomains(resp *http.Response) (result ListByDomainsOperationResponse, err error) {
	type page struct {
		Values   []SenderUsernameResource `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByDomainsOperationResponse, err error) {
			req, err := c.preparerForListByDomainsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByDomains(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "senderusernames.SenderUsernamesClient", "ListByDomains", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByDomainsComplete retrieves all of the results into a single object
func (c SenderUsernamesClient) ListByDomainsComplete(ctx context.Context, id DomainId) (ListByDomainsCompleteResult, error) {
	return c.ListByDomainsCompleteMatchingPredicate(ctx, id, SenderUsernameResourceOperationPredicate{})
}

// ListByDomainsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SenderUsernamesClient) ListByDomainsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate SenderUsernameResourceOperationPredicate) (resp ListByDomainsCompleteResult, err error) {
	items := make([]SenderUsernameResource, 0)

	page, err := c.ListByDomains(ctx, id)
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

	out := ListByDomainsCompleteResult{
		Items: items,
	}
	return out, nil
}
