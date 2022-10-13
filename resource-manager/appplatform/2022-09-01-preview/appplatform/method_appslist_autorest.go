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

type AppsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AppResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AppsListOperationResponse, error)
}

type AppsListCompleteResult struct {
	Items []AppResource
}

func (r AppsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AppsListOperationResponse) LoadMore(ctx context.Context) (resp AppsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AppsList ...
func (c AppPlatformClient) AppsList(ctx context.Context, id SpringId) (resp AppsListOperationResponse, err error) {
	req, err := c.preparerForAppsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAppsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAppsList prepares the AppsList request.
func (c AppPlatformClient) preparerForAppsList(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/apps", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAppsListWithNextLink prepares the AppsList request with the given nextLink token.
func (c AppPlatformClient) preparerForAppsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAppsList handles the response to the AppsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForAppsList(resp *http.Response) (result AppsListOperationResponse, err error) {
	type page struct {
		Values   []AppResource `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AppsListOperationResponse, err error) {
			req, err := c.preparerForAppsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAppsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AppsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) AppsListComplete(ctx context.Context, id SpringId) (AppsListCompleteResult, error) {
	return c.AppsListCompleteMatchingPredicate(ctx, id, AppResourceOperationPredicate{})
}

// AppsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) AppsListCompleteMatchingPredicate(ctx context.Context, id SpringId, predicate AppResourceOperationPredicate) (resp AppsListCompleteResult, err error) {
	items := make([]AppResource, 0)

	page, err := c.AppsList(ctx, id)
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

	out := AppsListCompleteResult{
		Items: items,
	}
	return out, nil
}
