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

type BuildServiceBuilderListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BuilderResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BuildServiceBuilderListOperationResponse, error)
}

type BuildServiceBuilderListCompleteResult struct {
	Items []BuilderResource
}

func (r BuildServiceBuilderListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BuildServiceBuilderListOperationResponse) LoadMore(ctx context.Context) (resp BuildServiceBuilderListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// BuildServiceBuilderList ...
func (c AppPlatformClient) BuildServiceBuilderList(ctx context.Context, id BuildServiceId) (resp BuildServiceBuilderListOperationResponse, err error) {
	req, err := c.preparerForBuildServiceBuilderList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBuildServiceBuilderList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBuildServiceBuilderList prepares the BuildServiceBuilderList request.
func (c AppPlatformClient) preparerForBuildServiceBuilderList(ctx context.Context, id BuildServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/builders", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBuildServiceBuilderListWithNextLink prepares the BuildServiceBuilderList request with the given nextLink token.
func (c AppPlatformClient) preparerForBuildServiceBuilderListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBuildServiceBuilderList handles the response to the BuildServiceBuilderList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceBuilderList(resp *http.Response) (result BuildServiceBuilderListOperationResponse, err error) {
	type page struct {
		Values   []BuilderResource `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BuildServiceBuilderListOperationResponse, err error) {
			req, err := c.preparerForBuildServiceBuilderListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBuildServiceBuilderList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BuildServiceBuilderListComplete retrieves all of the results into a single object
func (c AppPlatformClient) BuildServiceBuilderListComplete(ctx context.Context, id BuildServiceId) (BuildServiceBuilderListCompleteResult, error) {
	return c.BuildServiceBuilderListCompleteMatchingPredicate(ctx, id, BuilderResourceOperationPredicate{})
}

// BuildServiceBuilderListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) BuildServiceBuilderListCompleteMatchingPredicate(ctx context.Context, id BuildServiceId, predicate BuilderResourceOperationPredicate) (resp BuildServiceBuilderListCompleteResult, err error) {
	items := make([]BuilderResource, 0)

	page, err := c.BuildServiceBuilderList(ctx, id)
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

	out := BuildServiceBuilderListCompleteResult{
		Items: items,
	}
	return out, nil
}
