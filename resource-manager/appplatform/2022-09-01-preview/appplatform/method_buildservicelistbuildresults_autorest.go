package appplatform

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

type BuildServiceListBuildResultsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BuildResult

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BuildServiceListBuildResultsOperationResponse, error)
}

type BuildServiceListBuildResultsCompleteResult struct {
	Items []BuildResult
}

func (r BuildServiceListBuildResultsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BuildServiceListBuildResultsOperationResponse) LoadMore(ctx context.Context) (resp BuildServiceListBuildResultsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// BuildServiceListBuildResults ...
func (c AppPlatformClient) BuildServiceListBuildResults(ctx context.Context, id BuildId) (resp BuildServiceListBuildResultsOperationResponse, err error) {
	req, err := c.preparerForBuildServiceListBuildResults(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBuildServiceListBuildResults(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBuildServiceListBuildResults prepares the BuildServiceListBuildResults request.
func (c AppPlatformClient) preparerForBuildServiceListBuildResults(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/results", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBuildServiceListBuildResultsWithNextLink prepares the BuildServiceListBuildResults request with the given nextLink token.
func (c AppPlatformClient) preparerForBuildServiceListBuildResultsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBuildServiceListBuildResults handles the response to the BuildServiceListBuildResults request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceListBuildResults(resp *http.Response) (result BuildServiceListBuildResultsOperationResponse, err error) {
	type page struct {
		Values   []BuildResult `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BuildServiceListBuildResultsOperationResponse, err error) {
			req, err := c.preparerForBuildServiceListBuildResultsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBuildServiceListBuildResults(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildResults", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BuildServiceListBuildResultsComplete retrieves all of the results into a single object
func (c AppPlatformClient) BuildServiceListBuildResultsComplete(ctx context.Context, id BuildId) (BuildServiceListBuildResultsCompleteResult, error) {
	return c.BuildServiceListBuildResultsCompleteMatchingPredicate(ctx, id, BuildResultOperationPredicate{})
}

// BuildServiceListBuildResultsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) BuildServiceListBuildResultsCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate BuildResultOperationPredicate) (resp BuildServiceListBuildResultsCompleteResult, err error) {
	items := make([]BuildResult, 0)

	page, err := c.BuildServiceListBuildResults(ctx, id)
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

	out := BuildServiceListBuildResultsCompleteResult{
		Items: items,
	}
	return out, nil
}
