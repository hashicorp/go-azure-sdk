package componentversion

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

type RegistryComponentVersionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ComponentVersionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryComponentVersionsListOperationResponse, error)
}

type RegistryComponentVersionsListCompleteResult struct {
	Items []ComponentVersionResource
}

func (r RegistryComponentVersionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryComponentVersionsListOperationResponse) LoadMore(ctx context.Context) (resp RegistryComponentVersionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryComponentVersionsListOperationOptions struct {
	OrderBy *string
	Skip    *string
	Top     *int64
}

func DefaultRegistryComponentVersionsListOperationOptions() RegistryComponentVersionsListOperationOptions {
	return RegistryComponentVersionsListOperationOptions{}
}

func (o RegistryComponentVersionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryComponentVersionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// RegistryComponentVersionsList ...
func (c ComponentVersionClient) RegistryComponentVersionsList(ctx context.Context, id RegistryComponentId, options RegistryComponentVersionsListOperationOptions) (resp RegistryComponentVersionsListOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentVersionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryComponentVersionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryComponentVersionsList prepares the RegistryComponentVersionsList request.
func (c ComponentVersionClient) preparerForRegistryComponentVersionsList(ctx context.Context, id RegistryComponentId, options RegistryComponentVersionsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryComponentVersionsListWithNextLink prepares the RegistryComponentVersionsList request with the given nextLink token.
func (c ComponentVersionClient) preparerForRegistryComponentVersionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryComponentVersionsList handles the response to the RegistryComponentVersionsList request. The method always
// closes the http.Response Body.
func (c ComponentVersionClient) responderForRegistryComponentVersionsList(resp *http.Response) (result RegistryComponentVersionsListOperationResponse, err error) {
	type page struct {
		Values   []ComponentVersionResource `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryComponentVersionsListOperationResponse, err error) {
			req, err := c.preparerForRegistryComponentVersionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryComponentVersionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryComponentVersionsListComplete retrieves all of the results into a single object
func (c ComponentVersionClient) RegistryComponentVersionsListComplete(ctx context.Context, id RegistryComponentId, options RegistryComponentVersionsListOperationOptions) (RegistryComponentVersionsListCompleteResult, error) {
	return c.RegistryComponentVersionsListCompleteMatchingPredicate(ctx, id, options, ComponentVersionResourceOperationPredicate{})
}

// RegistryComponentVersionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ComponentVersionClient) RegistryComponentVersionsListCompleteMatchingPredicate(ctx context.Context, id RegistryComponentId, options RegistryComponentVersionsListOperationOptions, predicate ComponentVersionResourceOperationPredicate) (resp RegistryComponentVersionsListCompleteResult, err error) {
	items := make([]ComponentVersionResource, 0)

	page, err := c.RegistryComponentVersionsList(ctx, id, options)
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

	out := RegistryComponentVersionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
