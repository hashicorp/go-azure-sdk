package workflowrunactions

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

type ListExpressionTracesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]interface{}

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListExpressionTracesOperationResponse, error)
}

type ListExpressionTracesCompleteResult struct {
	Items []interface{}
}

func (r ListExpressionTracesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListExpressionTracesOperationResponse) LoadMore(ctx context.Context) (resp ListExpressionTracesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListExpressionTraces ...
func (c WorkflowRunActionsClient) ListExpressionTraces(ctx context.Context, id ActionId) (resp ListExpressionTracesOperationResponse, err error) {
	req, err := c.preparerForListExpressionTraces(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListExpressionTraces(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListExpressionTraces prepares the ListExpressionTraces request.
func (c WorkflowRunActionsClient) preparerForListExpressionTraces(ctx context.Context, id ActionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listExpressionTraces", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListExpressionTracesWithNextLink prepares the ListExpressionTraces request with the given nextLink token.
func (c WorkflowRunActionsClient) preparerForListExpressionTracesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListExpressionTraces handles the response to the ListExpressionTraces request. The method always
// closes the http.Response Body.
func (c WorkflowRunActionsClient) responderForListExpressionTraces(resp *http.Response) (result ListExpressionTracesOperationResponse, err error) {
	type page struct {
		Values   []interface{} `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListExpressionTracesOperationResponse, err error) {
			req, err := c.preparerForListExpressionTracesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListExpressionTraces(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "ListExpressionTraces", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListExpressionTracesComplete retrieves all of the results into a single object
func (c WorkflowRunActionsClient) ListExpressionTracesComplete(ctx context.Context, id ActionId) (result ListExpressionTracesCompleteResult, err error) {
	items := make([]interface{}, 0)

	page, err := c.ListExpressionTraces(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			items = append(items, v)
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
				items = append(items, v)
			}
		}
	}

	out := ListExpressionTracesCompleteResult{
		Items: items,
	}
	return out, nil
}
