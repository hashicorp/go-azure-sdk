package externalsecuritysolutions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHomeRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ExternalSecuritySolution

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByHomeRegionOperationResponse, error)
}

type ListByHomeRegionCompleteResult struct {
	Items []ExternalSecuritySolution
}

func (r ListByHomeRegionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByHomeRegionOperationResponse) LoadMore(ctx context.Context) (resp ListByHomeRegionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByHomeRegion ...
func (c ExternalSecuritySolutionsClient) ListByHomeRegion(ctx context.Context, id LocationId) (resp ListByHomeRegionOperationResponse, err error) {
	req, err := c.preparerForListByHomeRegion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByHomeRegion(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByHomeRegion prepares the ListByHomeRegion request.
func (c ExternalSecuritySolutionsClient) preparerForListByHomeRegion(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/externalSecuritySolutions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByHomeRegionWithNextLink prepares the ListByHomeRegion request with the given nextLink token.
func (c ExternalSecuritySolutionsClient) preparerForListByHomeRegionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByHomeRegion handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (c ExternalSecuritySolutionsClient) responderForListByHomeRegion(resp *http.Response) (result ListByHomeRegionOperationResponse, err error) {
	type page struct {
		Values   []json.RawMessage `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	temp := make([]ExternalSecuritySolution, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalExternalSecuritySolutionImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for ExternalSecuritySolution (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByHomeRegionOperationResponse, err error) {
			req, err := c.preparerForListByHomeRegionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByHomeRegion(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "externalsecuritysolutions.ExternalSecuritySolutionsClient", "ListByHomeRegion", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByHomeRegionComplete retrieves all of the results into a single object
func (c ExternalSecuritySolutionsClient) ListByHomeRegionComplete(ctx context.Context, id LocationId) (ListByHomeRegionCompleteResult, error) {
	return c.ListByHomeRegionCompleteMatchingPredicate(ctx, id, ExternalSecuritySolutionOperationPredicate{})
}

// ListByHomeRegionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ExternalSecuritySolutionsClient) ListByHomeRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate ExternalSecuritySolutionOperationPredicate) (resp ListByHomeRegionCompleteResult, err error) {
	items := make([]ExternalSecuritySolution, 0)

	page, err := c.ListByHomeRegion(ctx, id)
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

	out := ListByHomeRegionCompleteResult{
		Items: items,
	}
	return out, nil
}