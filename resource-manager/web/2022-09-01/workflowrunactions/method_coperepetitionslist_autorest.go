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

type CopeRepetitionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkflowRunActionRepetitionDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (CopeRepetitionsListOperationResponse, error)
}

type CopeRepetitionsListCompleteResult struct {
	Items []WorkflowRunActionRepetitionDefinition
}

func (r CopeRepetitionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r CopeRepetitionsListOperationResponse) LoadMore(ctx context.Context) (resp CopeRepetitionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// CopeRepetitionsList ...
func (c WorkflowRunActionsClient) CopeRepetitionsList(ctx context.Context, id ActionId) (resp CopeRepetitionsListOperationResponse, err error) {
	req, err := c.preparerForCopeRepetitionsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForCopeRepetitionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForCopeRepetitionsList prepares the CopeRepetitionsList request.
func (c WorkflowRunActionsClient) preparerForCopeRepetitionsList(ctx context.Context, id ActionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/scopeRepetitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForCopeRepetitionsListWithNextLink prepares the CopeRepetitionsList request with the given nextLink token.
func (c WorkflowRunActionsClient) preparerForCopeRepetitionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForCopeRepetitionsList handles the response to the CopeRepetitionsList request. The method always
// closes the http.Response Body.
func (c WorkflowRunActionsClient) responderForCopeRepetitionsList(resp *http.Response) (result CopeRepetitionsListOperationResponse, err error) {
	type page struct {
		Values   []WorkflowRunActionRepetitionDefinition `json:"value"`
		NextLink *string                                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result CopeRepetitionsListOperationResponse, err error) {
			req, err := c.preparerForCopeRepetitionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForCopeRepetitionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workflowrunactions.WorkflowRunActionsClient", "CopeRepetitionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// CopeRepetitionsListComplete retrieves all of the results into a single object
func (c WorkflowRunActionsClient) CopeRepetitionsListComplete(ctx context.Context, id ActionId) (CopeRepetitionsListCompleteResult, error) {
	return c.CopeRepetitionsListCompleteMatchingPredicate(ctx, id, WorkflowRunActionRepetitionDefinitionOperationPredicate{})
}

// CopeRepetitionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkflowRunActionsClient) CopeRepetitionsListCompleteMatchingPredicate(ctx context.Context, id ActionId, predicate WorkflowRunActionRepetitionDefinitionOperationPredicate) (resp CopeRepetitionsListCompleteResult, err error) {
	items := make([]WorkflowRunActionRepetitionDefinition, 0)

	page, err := c.CopeRepetitionsList(ctx, id)
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

	out := CopeRepetitionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
