package repositories

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

type SourceControllistRepositoriesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Repo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SourceControllistRepositoriesOperationResponse, error)
}

type SourceControllistRepositoriesCompleteResult struct {
	Items []Repo
}

func (r SourceControllistRepositoriesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SourceControllistRepositoriesOperationResponse) LoadMore(ctx context.Context) (resp SourceControllistRepositoriesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SourceControllistRepositories ...
func (c RepositoriesClient) SourceControllistRepositories(ctx context.Context, id WorkspaceId, input RepositoryAccessProperties) (resp SourceControllistRepositoriesOperationResponse, err error) {
	req, err := c.preparerForSourceControllistRepositories(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSourceControllistRepositories(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSourceControllistRepositories prepares the SourceControllistRepositories request.
func (c RepositoriesClient) preparerForSourceControllistRepositories(ctx context.Context, id WorkspaceId, input RepositoryAccessProperties) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/listRepositories", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSourceControllistRepositoriesWithNextLink prepares the SourceControllistRepositories request with the given nextLink token.
func (c RepositoriesClient) preparerForSourceControllistRepositoriesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSourceControllistRepositories handles the response to the SourceControllistRepositories request. The method always
// closes the http.Response Body.
func (c RepositoriesClient) responderForSourceControllistRepositories(resp *http.Response) (result SourceControllistRepositoriesOperationResponse, err error) {
	type page struct {
		Values   []Repo  `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SourceControllistRepositoriesOperationResponse, err error) {
			req, err := c.preparerForSourceControllistRepositoriesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSourceControllistRepositories(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "repositories.RepositoriesClient", "SourceControllistRepositories", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SourceControllistRepositoriesComplete retrieves all of the results into a single object
func (c RepositoriesClient) SourceControllistRepositoriesComplete(ctx context.Context, id WorkspaceId, input RepositoryAccessProperties) (SourceControllistRepositoriesCompleteResult, error) {
	return c.SourceControllistRepositoriesCompleteMatchingPredicate(ctx, id, input, RepoOperationPredicate{})
}

// SourceControllistRepositoriesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RepositoriesClient) SourceControllistRepositoriesCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, input RepositoryAccessProperties, predicate RepoOperationPredicate) (resp SourceControllistRepositoriesCompleteResult, err error) {
	items := make([]Repo, 0)

	page, err := c.SourceControllistRepositories(ctx, id, input)
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

	out := SourceControllistRepositoriesCompleteResult{
		Items: items,
	}
	return out, nil
}
