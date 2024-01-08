package contentproducttemplates

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

type ProductTemplatesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProductTemplateModel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ProductTemplatesListOperationResponse, error)
}

type ProductTemplatesListCompleteResult struct {
	Items []ProductTemplateModel
}

func (r ProductTemplatesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ProductTemplatesListOperationResponse) LoadMore(ctx context.Context) (resp ProductTemplatesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ProductTemplatesListOperationOptions struct {
	Count   *bool
	Filter  *string
	Orderby *string
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultProductTemplatesListOperationOptions() ProductTemplatesListOperationOptions {
	return ProductTemplatesListOperationOptions{}
}

func (o ProductTemplatesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ProductTemplatesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Count != nil {
		out["$count"] = *o.Count
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Search != nil {
		out["$search"] = *o.Search
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ProductTemplatesList ...
func (c ContentProductTemplatesClient) ProductTemplatesList(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions) (resp ProductTemplatesListOperationResponse, err error) {
	req, err := c.preparerForProductTemplatesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForProductTemplatesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForProductTemplatesList prepares the ProductTemplatesList request.
func (c ContentProductTemplatesClient) preparerForProductTemplatesList(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/contentProductTemplates", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForProductTemplatesListWithNextLink prepares the ProductTemplatesList request with the given nextLink token.
func (c ContentProductTemplatesClient) preparerForProductTemplatesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForProductTemplatesList handles the response to the ProductTemplatesList request. The method always
// closes the http.Response Body.
func (c ContentProductTemplatesClient) responderForProductTemplatesList(resp *http.Response) (result ProductTemplatesListOperationResponse, err error) {
	type page struct {
		Values   []ProductTemplateModel `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ProductTemplatesListOperationResponse, err error) {
			req, err := c.preparerForProductTemplatesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForProductTemplatesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproducttemplates.ContentProductTemplatesClient", "ProductTemplatesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ProductTemplatesListComplete retrieves all of the results into a single object
func (c ContentProductTemplatesClient) ProductTemplatesListComplete(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions) (ProductTemplatesListCompleteResult, error) {
	return c.ProductTemplatesListCompleteMatchingPredicate(ctx, id, options, ProductTemplateModelOperationPredicate{})
}

// ProductTemplatesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ContentProductTemplatesClient) ProductTemplatesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ProductTemplatesListOperationOptions, predicate ProductTemplateModelOperationPredicate) (resp ProductTemplatesListCompleteResult, err error) {
	items := make([]ProductTemplateModel, 0)

	page, err := c.ProductTemplatesList(ctx, id, options)
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

	out := ProductTemplatesListCompleteResult{
		Items: items,
	}
	return out, nil
}
