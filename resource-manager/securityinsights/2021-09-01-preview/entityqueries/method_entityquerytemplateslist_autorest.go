package entityqueries

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

type EntityQueryTemplatesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]EntityQueryTemplate

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (EntityQueryTemplatesListOperationResponse, error)
}

type EntityQueryTemplatesListCompleteResult struct {
	Items []EntityQueryTemplate
}

func (r EntityQueryTemplatesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r EntityQueryTemplatesListOperationResponse) LoadMore(ctx context.Context) (resp EntityQueryTemplatesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type EntityQueryTemplatesListOperationOptions struct {
	Kind *Kind
}

func DefaultEntityQueryTemplatesListOperationOptions() EntityQueryTemplatesListOperationOptions {
	return EntityQueryTemplatesListOperationOptions{}
}

func (o EntityQueryTemplatesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o EntityQueryTemplatesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Kind != nil {
		out["kind"] = *o.Kind
	}

	return out
}

// EntityQueryTemplatesList ...
func (c EntityQueriesClient) EntityQueryTemplatesList(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions) (resp EntityQueryTemplatesListOperationResponse, err error) {
	req, err := c.preparerForEntityQueryTemplatesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForEntityQueryTemplatesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForEntityQueryTemplatesList prepares the EntityQueryTemplatesList request.
func (c EntityQueriesClient) preparerForEntityQueryTemplatesList(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/entityQueryTemplates", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForEntityQueryTemplatesListWithNextLink prepares the EntityQueryTemplatesList request with the given nextLink token.
func (c EntityQueriesClient) preparerForEntityQueryTemplatesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForEntityQueryTemplatesList handles the response to the EntityQueryTemplatesList request. The method always
// closes the http.Response Body.
func (c EntityQueriesClient) responderForEntityQueryTemplatesList(resp *http.Response) (result EntityQueryTemplatesListOperationResponse, err error) {
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
	temp := make([]EntityQueryTemplate, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalEntityQueryTemplateImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for EntityQueryTemplate (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result EntityQueryTemplatesListOperationResponse, err error) {
			req, err := c.preparerForEntityQueryTemplatesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForEntityQueryTemplatesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// EntityQueryTemplatesListComplete retrieves all of the results into a single object
func (c EntityQueriesClient) EntityQueryTemplatesListComplete(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions) (EntityQueryTemplatesListCompleteResult, error) {
	return c.EntityQueryTemplatesListCompleteMatchingPredicate(ctx, id, options, EntityQueryTemplateOperationPredicate{})
}

// EntityQueryTemplatesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c EntityQueriesClient) EntityQueryTemplatesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options EntityQueryTemplatesListOperationOptions, predicate EntityQueryTemplateOperationPredicate) (resp EntityQueryTemplatesListCompleteResult, err error) {
	items := make([]EntityQueryTemplate, 0)

	page, err := c.EntityQueryTemplatesList(ctx, id, options)
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

	out := EntityQueryTemplatesListCompleteResult{
		Items: items,
	}
	return out, nil
}
