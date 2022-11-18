package usersession

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

type ListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]UserSession

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByHostPoolOperationResponse, error)
}

type ListByHostPoolCompleteResult struct {
	Items []UserSession
}

func (r ListByHostPoolOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByHostPoolOperationResponse) LoadMore(ctx context.Context) (resp ListByHostPoolOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByHostPoolOperationOptions struct {
	Filter       *string
	InitialSkip  *int64
	IsDescending *bool
	PageSize     *int64
}

func DefaultListByHostPoolOperationOptions() ListByHostPoolOperationOptions {
	return ListByHostPoolOperationOptions{}
}

func (o ListByHostPoolOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByHostPoolOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.InitialSkip != nil {
		out["initialSkip"] = *o.InitialSkip
	}

	if o.IsDescending != nil {
		out["isDescending"] = *o.IsDescending
	}

	if o.PageSize != nil {
		out["pageSize"] = *o.PageSize
	}

	return out
}

// ListByHostPool ...
func (c UserSessionClient) ListByHostPool(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions) (resp ListByHostPoolOperationResponse, err error) {
	req, err := c.preparerForListByHostPool(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByHostPool(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByHostPool prepares the ListByHostPool request.
func (c UserSessionClient) preparerForListByHostPool(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/userSessions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByHostPoolWithNextLink prepares the ListByHostPool request with the given nextLink token.
func (c UserSessionClient) preparerForListByHostPoolWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByHostPool handles the response to the ListByHostPool request. The method always
// closes the http.Response Body.
func (c UserSessionClient) responderForListByHostPool(resp *http.Response) (result ListByHostPoolOperationResponse, err error) {
	type page struct {
		Values   []UserSession `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByHostPoolOperationResponse, err error) {
			req, err := c.preparerForListByHostPoolWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByHostPool(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "usersession.UserSessionClient", "ListByHostPool", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByHostPoolComplete retrieves all of the results into a single object
func (c UserSessionClient) ListByHostPoolComplete(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions) (ListByHostPoolCompleteResult, error) {
	return c.ListByHostPoolCompleteMatchingPredicate(ctx, id, options, UserSessionOperationPredicate{})
}

// ListByHostPoolCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c UserSessionClient) ListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, options ListByHostPoolOperationOptions, predicate UserSessionOperationPredicate) (resp ListByHostPoolCompleteResult, err error) {
	items := make([]UserSession, 0)

	page, err := c.ListByHostPool(ctx, id, options)
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

	out := ListByHostPoolCompleteResult{
		Items: items,
	}
	return out, nil
}
