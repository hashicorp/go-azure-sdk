package appserviceenvironments

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

type ListWebAppsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Site

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWebAppsOperationResponse, error)
}

type ListWebAppsCompleteResult struct {
	Items []Site
}

func (r ListWebAppsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWebAppsOperationResponse) LoadMore(ctx context.Context) (resp ListWebAppsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListWebAppsOperationOptions struct {
	PropertiesToInclude *string
}

func DefaultListWebAppsOperationOptions() ListWebAppsOperationOptions {
	return ListWebAppsOperationOptions{}
}

func (o ListWebAppsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListWebAppsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.PropertiesToInclude != nil {
		out["propertiesToInclude"] = *o.PropertiesToInclude
	}

	return out
}

// ListWebApps ...
func (c AppServiceEnvironmentsClient) ListWebApps(ctx context.Context, id HostingEnvironmentId, options ListWebAppsOperationOptions) (resp ListWebAppsOperationResponse, err error) {
	req, err := c.preparerForListWebApps(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWebApps(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWebApps prepares the ListWebApps request.
func (c AppServiceEnvironmentsClient) preparerForListWebApps(ctx context.Context, id HostingEnvironmentId, options ListWebAppsOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/sites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWebAppsWithNextLink prepares the ListWebApps request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListWebAppsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWebApps handles the response to the ListWebApps request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListWebApps(resp *http.Response) (result ListWebAppsOperationResponse, err error) {
	type page struct {
		Values   []Site  `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWebAppsOperationResponse, err error) {
			req, err := c.preparerForListWebAppsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWebApps(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebApps", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWebAppsComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListWebAppsComplete(ctx context.Context, id HostingEnvironmentId, options ListWebAppsOperationOptions) (ListWebAppsCompleteResult, error) {
	return c.ListWebAppsCompleteMatchingPredicate(ctx, id, options, SiteOperationPredicate{})
}

// ListWebAppsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListWebAppsCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, options ListWebAppsOperationOptions, predicate SiteOperationPredicate) (resp ListWebAppsCompleteResult, err error) {
	items := make([]Site, 0)

	page, err := c.ListWebApps(ctx, id, options)
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

	out := ListWebAppsCompleteResult{
		Items: items,
	}
	return out, nil
}
