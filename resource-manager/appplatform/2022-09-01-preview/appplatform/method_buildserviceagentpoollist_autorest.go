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

type BuildServiceAgentPoolListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BuildServiceAgentPoolResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BuildServiceAgentPoolListOperationResponse, error)
}

type BuildServiceAgentPoolListCompleteResult struct {
	Items []BuildServiceAgentPoolResource
}

func (r BuildServiceAgentPoolListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BuildServiceAgentPoolListOperationResponse) LoadMore(ctx context.Context) (resp BuildServiceAgentPoolListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// BuildServiceAgentPoolList ...
func (c AppPlatformClient) BuildServiceAgentPoolList(ctx context.Context, id BuildServiceId) (resp BuildServiceAgentPoolListOperationResponse, err error) {
	req, err := c.preparerForBuildServiceAgentPoolList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBuildServiceAgentPoolList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBuildServiceAgentPoolList prepares the BuildServiceAgentPoolList request.
func (c AppPlatformClient) preparerForBuildServiceAgentPoolList(ctx context.Context, id BuildServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/agentPools", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBuildServiceAgentPoolListWithNextLink prepares the BuildServiceAgentPoolList request with the given nextLink token.
func (c AppPlatformClient) preparerForBuildServiceAgentPoolListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBuildServiceAgentPoolList handles the response to the BuildServiceAgentPoolList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceAgentPoolList(resp *http.Response) (result BuildServiceAgentPoolListOperationResponse, err error) {
	type page struct {
		Values   []BuildServiceAgentPoolResource `json:"value"`
		NextLink *string                         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BuildServiceAgentPoolListOperationResponse, err error) {
			req, err := c.preparerForBuildServiceAgentPoolListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBuildServiceAgentPoolList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceAgentPoolList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BuildServiceAgentPoolListComplete retrieves all of the results into a single object
func (c AppPlatformClient) BuildServiceAgentPoolListComplete(ctx context.Context, id BuildServiceId) (BuildServiceAgentPoolListCompleteResult, error) {
	return c.BuildServiceAgentPoolListCompleteMatchingPredicate(ctx, id, BuildServiceAgentPoolResourceOperationPredicate{})
}

// BuildServiceAgentPoolListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) BuildServiceAgentPoolListCompleteMatchingPredicate(ctx context.Context, id BuildServiceId, predicate BuildServiceAgentPoolResourceOperationPredicate) (resp BuildServiceAgentPoolListCompleteResult, err error) {
	items := make([]BuildServiceAgentPoolResource, 0)

	page, err := c.BuildServiceAgentPoolList(ctx, id)
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

	out := BuildServiceAgentPoolListCompleteResult{
		Items: items,
	}
	return out, nil
}
