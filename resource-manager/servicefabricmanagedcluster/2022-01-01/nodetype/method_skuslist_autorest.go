package nodetype

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

type SkusListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NodeTypeAvailableSku

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SkusListOperationResponse, error)
}

type SkusListCompleteResult struct {
	Items []NodeTypeAvailableSku
}

func (r SkusListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SkusListOperationResponse) LoadMore(ctx context.Context) (resp SkusListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SkusList ...
func (c NodeTypeClient) SkusList(ctx context.Context, id NodeTypeId) (resp SkusListOperationResponse, err error) {
	req, err := c.preparerForSkusList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSkusList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSkusList prepares the SkusList request.
func (c NodeTypeClient) preparerForSkusList(ctx context.Context, id NodeTypeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSkusListWithNextLink prepares the SkusList request with the given nextLink token.
func (c NodeTypeClient) preparerForSkusListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSkusList handles the response to the SkusList request. The method always
// closes the http.Response Body.
func (c NodeTypeClient) responderForSkusList(resp *http.Response) (result SkusListOperationResponse, err error) {
	type page struct {
		Values   []NodeTypeAvailableSku `json:"value"`
		NextLink *string                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SkusListOperationResponse, err error) {
			req, err := c.preparerForSkusListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSkusList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "nodetype.NodeTypeClient", "SkusList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SkusListComplete retrieves all of the results into a single object
func (c NodeTypeClient) SkusListComplete(ctx context.Context, id NodeTypeId) (SkusListCompleteResult, error) {
	return c.SkusListCompleteMatchingPredicate(ctx, id, NodeTypeAvailableSkuOperationPredicate{})
}

// SkusListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c NodeTypeClient) SkusListCompleteMatchingPredicate(ctx context.Context, id NodeTypeId, predicate NodeTypeAvailableSkuOperationPredicate) (resp SkusListCompleteResult, err error) {
	items := make([]NodeTypeAvailableSku, 0)

	page, err := c.SkusList(ctx, id)
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

	out := SkusListCompleteResult{
		Items: items,
	}
	return out, nil
}
