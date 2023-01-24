package standardoperation

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

type ProjectsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Project

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ProjectsListOperationResponse, error)
}

type ProjectsListCompleteResult struct {
	Items []Project
}

func (r ProjectsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ProjectsListOperationResponse) LoadMore(ctx context.Context) (resp ProjectsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ProjectsList ...
func (c StandardOperationClient) ProjectsList(ctx context.Context, id ServiceId) (resp ProjectsListOperationResponse, err error) {
	req, err := c.preparerForProjectsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForProjectsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForProjectsList prepares the ProjectsList request.
func (c StandardOperationClient) preparerForProjectsList(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/projects", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForProjectsListWithNextLink prepares the ProjectsList request with the given nextLink token.
func (c StandardOperationClient) preparerForProjectsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForProjectsList handles the response to the ProjectsList request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForProjectsList(resp *http.Response) (result ProjectsListOperationResponse, err error) {
	type page struct {
		Values   []Project `json:"value"`
		NextLink *string   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ProjectsListOperationResponse, err error) {
			req, err := c.preparerForProjectsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForProjectsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ProjectsListComplete retrieves all of the results into a single object
func (c StandardOperationClient) ProjectsListComplete(ctx context.Context, id ServiceId) (ProjectsListCompleteResult, error) {
	return c.ProjectsListCompleteMatchingPredicate(ctx, id, ProjectOperationPredicate{})
}

// ProjectsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StandardOperationClient) ProjectsListCompleteMatchingPredicate(ctx context.Context, id ServiceId, predicate ProjectOperationPredicate) (resp ProjectsListCompleteResult, err error) {
	items := make([]Project, 0)

	page, err := c.ProjectsList(ctx, id)
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

	out := ProjectsListCompleteResult{
		Items: items,
	}
	return out, nil
}
