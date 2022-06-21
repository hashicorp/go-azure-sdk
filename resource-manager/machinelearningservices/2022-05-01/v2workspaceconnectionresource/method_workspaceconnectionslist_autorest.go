package v2workspaceconnectionresource

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

type WorkspaceConnectionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkspaceConnectionPropertiesV2BasicResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceConnectionsListOperationResponse, error)
}

type WorkspaceConnectionsListCompleteResult struct {
	Items []WorkspaceConnectionPropertiesV2BasicResource
}

func (r WorkspaceConnectionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceConnectionsListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceConnectionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type WorkspaceConnectionsListOperationOptions struct {
	Category *string
	Target   *string
}

func DefaultWorkspaceConnectionsListOperationOptions() WorkspaceConnectionsListOperationOptions {
	return WorkspaceConnectionsListOperationOptions{}
}

func (o WorkspaceConnectionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o WorkspaceConnectionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Category != nil {
		out["category"] = *o.Category
	}

	if o.Target != nil {
		out["target"] = *o.Target
	}

	return out
}

// WorkspaceConnectionsList ...
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsList(ctx context.Context, id WorkspaceId, options WorkspaceConnectionsListOperationOptions) (resp WorkspaceConnectionsListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceConnectionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceConnectionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// WorkspaceConnectionsListComplete retrieves all of the results into a single object
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsListComplete(ctx context.Context, id WorkspaceId, options WorkspaceConnectionsListOperationOptions) (WorkspaceConnectionsListCompleteResult, error) {
	return c.WorkspaceConnectionsListCompleteMatchingPredicate(ctx, id, options, WorkspaceConnectionPropertiesV2BasicResourceOperationPredicate{})
}

// WorkspaceConnectionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options WorkspaceConnectionsListOperationOptions, predicate WorkspaceConnectionPropertiesV2BasicResourceOperationPredicate) (resp WorkspaceConnectionsListCompleteResult, err error) {
	items := make([]WorkspaceConnectionPropertiesV2BasicResource, 0)

	page, err := c.WorkspaceConnectionsList(ctx, id, options)
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

	out := WorkspaceConnectionsListCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForWorkspaceConnectionsList prepares the WorkspaceConnectionsList request.
func (c V2WorkspaceConnectionResourceClient) preparerForWorkspaceConnectionsList(ctx context.Context, id WorkspaceId, options WorkspaceConnectionsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/connections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceConnectionsListWithNextLink prepares the WorkspaceConnectionsList request with the given nextLink token.
func (c V2WorkspaceConnectionResourceClient) preparerForWorkspaceConnectionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceConnectionsList handles the response to the WorkspaceConnectionsList request. The method always
// closes the http.Response Body.
func (c V2WorkspaceConnectionResourceClient) responderForWorkspaceConnectionsList(resp *http.Response) (result WorkspaceConnectionsListOperationResponse, err error) {
	type page struct {
		Values   []WorkspaceConnectionPropertiesV2BasicResource `json:"value"`
		NextLink *string                                        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceConnectionsListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceConnectionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceConnectionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
