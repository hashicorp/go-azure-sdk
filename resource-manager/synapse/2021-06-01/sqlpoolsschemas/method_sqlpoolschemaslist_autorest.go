package sqlpoolsschemas

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

type SqlPoolSchemasListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Resource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolSchemasListOperationResponse, error)
}

type SqlPoolSchemasListCompleteResult struct {
	Items []Resource
}

func (r SqlPoolSchemasListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolSchemasListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolSchemasListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type SqlPoolSchemasListOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolSchemasListOperationOptions() SqlPoolSchemasListOperationOptions {
	return SqlPoolSchemasListOperationOptions{}
}

func (o SqlPoolSchemasListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolSchemasListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// SqlPoolSchemasList ...
func (c SqlPoolsSchemasClient) SqlPoolSchemasList(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions) (resp SqlPoolSchemasListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSchemasList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolSchemasList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolSchemasList prepares the SqlPoolSchemasList request.
func (c SqlPoolsSchemasClient) preparerForSqlPoolSchemasList(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/schemas", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolSchemasListWithNextLink prepares the SqlPoolSchemasList request with the given nextLink token.
func (c SqlPoolsSchemasClient) preparerForSqlPoolSchemasListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolSchemasList handles the response to the SqlPoolSchemasList request. The method always
// closes the http.Response Body.
func (c SqlPoolsSchemasClient) responderForSqlPoolSchemasList(resp *http.Response) (result SqlPoolSchemasListOperationResponse, err error) {
	type page struct {
		Values   []Resource `json:"value"`
		NextLink *string    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolSchemasListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolSchemasListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolSchemasList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsschemas.SqlPoolsSchemasClient", "SqlPoolSchemasList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolSchemasListComplete retrieves all of the results into a single object
func (c SqlPoolsSchemasClient) SqlPoolSchemasListComplete(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions) (SqlPoolSchemasListCompleteResult, error) {
	return c.SqlPoolSchemasListCompleteMatchingPredicate(ctx, id, options, ResourceOperationPredicate{})
}

// SqlPoolSchemasListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsSchemasClient) SqlPoolSchemasListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSchemasListOperationOptions, predicate ResourceOperationPredicate) (resp SqlPoolSchemasListCompleteResult, err error) {
	items := make([]Resource, 0)

	page, err := c.SqlPoolSchemasList(ctx, id, options)
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

	out := SqlPoolSchemasListCompleteResult{
		Items: items,
	}
	return out, nil
}
