package migrations

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

type ListByTargetServerOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]MigrationResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByTargetServerOperationResponse, error)
}

type ListByTargetServerCompleteResult struct {
	Items []MigrationResource
}

func (r ListByTargetServerOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByTargetServerOperationResponse) LoadMore(ctx context.Context) (resp ListByTargetServerOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByTargetServerOperationOptions struct {
	MigrationListFilter *MigrationListFilter
}

func DefaultListByTargetServerOperationOptions() ListByTargetServerOperationOptions {
	return ListByTargetServerOperationOptions{}
}

func (o ListByTargetServerOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByTargetServerOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.MigrationListFilter != nil {
		out["migrationListFilter"] = *o.MigrationListFilter
	}

	return out
}

// ListByTargetServer ...
func (c MigrationsClient) ListByTargetServer(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions) (resp ListByTargetServerOperationResponse, err error) {
	req, err := c.preparerForListByTargetServer(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByTargetServer(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByTargetServer prepares the ListByTargetServer request.
func (c MigrationsClient) preparerForListByTargetServer(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/migrations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByTargetServerWithNextLink prepares the ListByTargetServer request with the given nextLink token.
func (c MigrationsClient) preparerForListByTargetServerWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByTargetServer handles the response to the ListByTargetServer request. The method always
// closes the http.Response Body.
func (c MigrationsClient) responderForListByTargetServer(resp *http.Response) (result ListByTargetServerOperationResponse, err error) {
	type page struct {
		Values   []MigrationResource `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByTargetServerOperationResponse, err error) {
			req, err := c.preparerForListByTargetServerWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByTargetServer(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrations.MigrationsClient", "ListByTargetServer", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByTargetServerComplete retrieves all of the results into a single object
func (c MigrationsClient) ListByTargetServerComplete(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions) (ListByTargetServerCompleteResult, error) {
	return c.ListByTargetServerCompleteMatchingPredicate(ctx, id, options, MigrationResourceOperationPredicate{})
}

// ListByTargetServerCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MigrationsClient) ListByTargetServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, options ListByTargetServerOperationOptions, predicate MigrationResourceOperationPredicate) (resp ListByTargetServerCompleteResult, err error) {
	items := make([]MigrationResource, 0)

	page, err := c.ListByTargetServer(ctx, id, options)
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

	out := ListByTargetServerCompleteResult{
		Items: items,
	}
	return out, nil
}
