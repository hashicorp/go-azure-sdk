package dataversionregistry

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

type RegistryDataVersionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DataVersionBaseResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryDataVersionsListOperationResponse, error)
}

type RegistryDataVersionsListCompleteResult struct {
	Items []DataVersionBaseResource
}

func (r RegistryDataVersionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryDataVersionsListOperationResponse) LoadMore(ctx context.Context) (resp RegistryDataVersionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryDataVersionsListOperationOptions struct {
	ListViewType *ListViewType
	OrderBy      *string
	Skip         *string
	Tags         *string
	Top          *int64
}

func DefaultRegistryDataVersionsListOperationOptions() RegistryDataVersionsListOperationOptions {
	return RegistryDataVersionsListOperationOptions{}
}

func (o RegistryDataVersionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryDataVersionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Tags != nil {
		out["$tags"] = *o.Tags
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// RegistryDataVersionsList ...
func (c DataVersionRegistryClient) RegistryDataVersionsList(ctx context.Context, id DataId, options RegistryDataVersionsListOperationOptions) (resp RegistryDataVersionsListOperationResponse, err error) {
	req, err := c.preparerForRegistryDataVersionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryDataVersionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryDataVersionsList prepares the RegistryDataVersionsList request.
func (c DataVersionRegistryClient) preparerForRegistryDataVersionsList(ctx context.Context, id DataId, options RegistryDataVersionsListOperationOptions) (*http.Request, error) {
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

// preparerForRegistryDataVersionsListWithNextLink prepares the RegistryDataVersionsList request with the given nextLink token.
func (c DataVersionRegistryClient) preparerForRegistryDataVersionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryDataVersionsList handles the response to the RegistryDataVersionsList request. The method always
// closes the http.Response Body.
func (c DataVersionRegistryClient) responderForRegistryDataVersionsList(resp *http.Response) (result RegistryDataVersionsListOperationResponse, err error) {
	type page struct {
		Values   []DataVersionBaseResource `json:"value"`
		NextLink *string                   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryDataVersionsListOperationResponse, err error) {
			req, err := c.preparerForRegistryDataVersionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryDataVersionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryDataVersionsListComplete retrieves all of the results into a single object
func (c DataVersionRegistryClient) RegistryDataVersionsListComplete(ctx context.Context, id DataId, options RegistryDataVersionsListOperationOptions) (RegistryDataVersionsListCompleteResult, error) {
	return c.RegistryDataVersionsListCompleteMatchingPredicate(ctx, id, options, DataVersionBaseResourceOperationPredicate{})
}

// RegistryDataVersionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DataVersionRegistryClient) RegistryDataVersionsListCompleteMatchingPredicate(ctx context.Context, id DataId, options RegistryDataVersionsListOperationOptions, predicate DataVersionBaseResourceOperationPredicate) (resp RegistryDataVersionsListCompleteResult, err error) {
	items := make([]DataVersionBaseResource, 0)

	page, err := c.RegistryDataVersionsList(ctx, id, options)
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

	out := RegistryDataVersionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
