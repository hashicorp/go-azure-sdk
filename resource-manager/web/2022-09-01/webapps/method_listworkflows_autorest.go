package webapps

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

type ListWorkflowsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkflowEnvelope

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWorkflowsOperationResponse, error)
}

type ListWorkflowsCompleteResult struct {
	Items []WorkflowEnvelope
}

func (r ListWorkflowsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWorkflowsOperationResponse) LoadMore(ctx context.Context) (resp ListWorkflowsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListWorkflows ...
func (c WebAppsClient) ListWorkflows(ctx context.Context, id SiteId) (resp ListWorkflowsOperationResponse, err error) {
	req, err := c.preparerForListWorkflows(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWorkflows(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWorkflows prepares the ListWorkflows request.
func (c WebAppsClient) preparerForListWorkflows(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workflows", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWorkflowsWithNextLink prepares the ListWorkflows request with the given nextLink token.
func (c WebAppsClient) preparerForListWorkflowsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWorkflows handles the response to the ListWorkflows request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListWorkflows(resp *http.Response) (result ListWorkflowsOperationResponse, err error) {
	type page struct {
		Values   []WorkflowEnvelope `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWorkflowsOperationResponse, err error) {
			req, err := c.preparerForListWorkflowsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWorkflows(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflows", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWorkflowsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListWorkflowsComplete(ctx context.Context, id SiteId) (ListWorkflowsCompleteResult, error) {
	return c.ListWorkflowsCompleteMatchingPredicate(ctx, id, WorkflowEnvelopeOperationPredicate{})
}

// ListWorkflowsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListWorkflowsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate WorkflowEnvelopeOperationPredicate) (resp ListWorkflowsCompleteResult, err error) {
	items := make([]WorkflowEnvelope, 0)

	page, err := c.ListWorkflows(ctx, id)
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

	out := ListWorkflowsCompleteResult{
		Items: items,
	}
	return out, nil
}
