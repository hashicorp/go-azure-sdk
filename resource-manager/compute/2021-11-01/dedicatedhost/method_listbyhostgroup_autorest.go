package dedicatedhost

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

type ListByHostGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DedicatedHost

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByHostGroupOperationResponse, error)
}

type ListByHostGroupCompleteResult struct {
	Items []DedicatedHost
}

func (r ListByHostGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByHostGroupOperationResponse) LoadMore(ctx context.Context) (resp ListByHostGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByHostGroup ...
func (c DedicatedHostClient) ListByHostGroup(ctx context.Context, id HostGroupId) (resp ListByHostGroupOperationResponse, err error) {
	req, err := c.preparerForListByHostGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByHostGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListByHostGroupComplete retrieves all of the results into a single object
func (c DedicatedHostClient) ListByHostGroupComplete(ctx context.Context, id HostGroupId) (ListByHostGroupCompleteResult, error) {
	return c.ListByHostGroupCompleteMatchingPredicate(ctx, id, DedicatedHostOperationPredicate{})
}

// ListByHostGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DedicatedHostClient) ListByHostGroupCompleteMatchingPredicate(ctx context.Context, id HostGroupId, predicate DedicatedHostOperationPredicate) (resp ListByHostGroupCompleteResult, err error) {
	items := make([]DedicatedHost, 0)

	page, err := c.ListByHostGroup(ctx, id)
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

	out := ListByHostGroupCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListByHostGroup prepares the ListByHostGroup request.
func (c DedicatedHostClient) preparerForListByHostGroup(ctx context.Context, id HostGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hosts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByHostGroupWithNextLink prepares the ListByHostGroup request with the given nextLink token.
func (c DedicatedHostClient) preparerForListByHostGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByHostGroup handles the response to the ListByHostGroup request. The method always
// closes the http.Response Body.
func (c DedicatedHostClient) responderForListByHostGroup(resp *http.Response) (result ListByHostGroupOperationResponse, err error) {
	type page struct {
		Values   []DedicatedHost `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByHostGroupOperationResponse, err error) {
			req, err := c.preparerForListByHostGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByHostGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dedicatedhost.DedicatedHostClient", "ListByHostGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
