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

type WorkflowRunActionRepetitionsListExpressionTracesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]interface{}

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, error)
}

type WorkflowRunActionRepetitionsListExpressionTracesCompleteResult struct {
	Items []interface{}
}

func (r WorkflowRunActionRepetitionsListExpressionTracesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkflowRunActionRepetitionsListExpressionTracesOperationResponse) LoadMore(ctx context.Context) (resp WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkflowRunActionRepetitionsListExpressionTraces ...
func (c WorkflowRunActionsClient) WorkflowRunActionRepetitionsListExpressionTraces(ctx context.Context, id RepetitionId) (resp WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, err error) {
	req, err := c.preparerForWorkflowRunActionRepetitionsListExpressionTraces(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkflowRunActionRepetitionsListExpressionTraces(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkflowRunActionRepetitionsListExpressionTraces prepares the WorkflowRunActionRepetitionsListExpressionTraces request.
func (c WorkflowRunActionsClient) preparerForWorkflowRunActionRepetitionsListExpressionTraces(ctx context.Context, id RepetitionId) (*http.Request, error) {
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

// preparerForWorkflowRunActionRepetitionsListExpressionTracesWithNextLink prepares the WorkflowRunActionRepetitionsListExpressionTraces request with the given nextLink token.
func (c WorkflowRunActionsClient) preparerForWorkflowRunActionRepetitionsListExpressionTracesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkflowRunActionRepetitionsListExpressionTraces handles the response to the WorkflowRunActionRepetitionsListExpressionTraces request. The method always
// closes the http.Response Body.
func (c WorkflowRunActionsClient) responderForWorkflowRunActionRepetitionsListExpressionTraces(resp *http.Response) (result WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkflowRunActionRepetitionsListExpressionTracesOperationResponse, err error) {
			req, err := c.preparerForWorkflowRunActionRepetitionsListExpressionTracesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkflowRunActionRepetitionsListExpressionTraces(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "WorkflowRunActionRepetitionsListExpressionTraces", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkflowRunActionRepetitionsListExpressionTracesComplete retrieves all of the results into a single object
func (c WorkflowRunActionsClient) WorkflowRunActionRepetitionsListExpressionTracesComplete(ctx context.Context, id RepetitionId) (result WorkflowRunActionRepetitionsListExpressionTracesCompleteResult, err error) {
	items := make([]interface{}, 0)

	page, err := c.WorkflowRunActionRepetitionsListExpressionTraces(ctx, id)
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

	out := WorkflowRunActionRepetitionsListExpressionTracesCompleteResult{
		Items: items,
	}
	return out, nil
}
