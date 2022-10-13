package appplatform

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

type ApiPortalsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiPortalResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ApiPortalsListOperationResponse, error)
}

type ApiPortalsListCompleteResult struct {
	Items []ApiPortalResource
}

func (r ApiPortalsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ApiPortalsListOperationResponse) LoadMore(ctx context.Context) (resp ApiPortalsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ApiPortalsList ...
func (c AppPlatformClient) ApiPortalsList(ctx context.Context, id SpringId) (resp ApiPortalsListOperationResponse, err error) {
	req, err := c.preparerForApiPortalsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForApiPortalsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForApiPortalsList prepares the ApiPortalsList request.
func (c AppPlatformClient) preparerForApiPortalsList(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/apiPortals", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForApiPortalsListWithNextLink prepares the ApiPortalsList request with the given nextLink token.
func (c AppPlatformClient) preparerForApiPortalsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForApiPortalsList handles the response to the ApiPortalsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForApiPortalsList(resp *http.Response) (result ApiPortalsListOperationResponse, err error) {
	type page struct {
		Values   []ApiPortalResource `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ApiPortalsListOperationResponse, err error) {
			req, err := c.preparerForApiPortalsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForApiPortalsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ApiPortalsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) ApiPortalsListComplete(ctx context.Context, id SpringId) (ApiPortalsListCompleteResult, error) {
	return c.ApiPortalsListCompleteMatchingPredicate(ctx, id, ApiPortalResourceOperationPredicate{})
}

// ApiPortalsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) ApiPortalsListCompleteMatchingPredicate(ctx context.Context, id SpringId, predicate ApiPortalResourceOperationPredicate) (resp ApiPortalsListCompleteResult, err error) {
	items := make([]ApiPortalResource, 0)

	page, err := c.ApiPortalsList(ctx, id)
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

	out := ApiPortalsListCompleteResult{
		Items: items,
	}
	return out, nil
}
