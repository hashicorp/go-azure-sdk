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

type BuildServiceListBuildServicesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BuildService

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BuildServiceListBuildServicesOperationResponse, error)
}

type BuildServiceListBuildServicesCompleteResult struct {
	Items []BuildService
}

func (r BuildServiceListBuildServicesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BuildServiceListBuildServicesOperationResponse) LoadMore(ctx context.Context) (resp BuildServiceListBuildServicesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// BuildServiceListBuildServices ...
func (c AppPlatformClient) BuildServiceListBuildServices(ctx context.Context, id SpringId) (resp BuildServiceListBuildServicesOperationResponse, err error) {
	req, err := c.preparerForBuildServiceListBuildServices(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBuildServiceListBuildServices(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBuildServiceListBuildServices prepares the BuildServiceListBuildServices request.
func (c AppPlatformClient) preparerForBuildServiceListBuildServices(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/buildServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBuildServiceListBuildServicesWithNextLink prepares the BuildServiceListBuildServices request with the given nextLink token.
func (c AppPlatformClient) preparerForBuildServiceListBuildServicesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBuildServiceListBuildServices handles the response to the BuildServiceListBuildServices request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBuildServiceListBuildServices(resp *http.Response) (result BuildServiceListBuildServicesOperationResponse, err error) {
	type page struct {
		Values   []BuildService `json:"value"`
		NextLink *string        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BuildServiceListBuildServicesOperationResponse, err error) {
			req, err := c.preparerForBuildServiceListBuildServicesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBuildServiceListBuildServices(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceListBuildServices", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BuildServiceListBuildServicesComplete retrieves all of the results into a single object
func (c AppPlatformClient) BuildServiceListBuildServicesComplete(ctx context.Context, id SpringId) (BuildServiceListBuildServicesCompleteResult, error) {
	return c.BuildServiceListBuildServicesCompleteMatchingPredicate(ctx, id, BuildServiceOperationPredicate{})
}

// BuildServiceListBuildServicesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) BuildServiceListBuildServicesCompleteMatchingPredicate(ctx context.Context, id SpringId, predicate BuildServiceOperationPredicate) (resp BuildServiceListBuildServicesCompleteResult, err error) {
	items := make([]BuildService, 0)

	page, err := c.BuildServiceListBuildServices(ctx, id)
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

	out := BuildServiceListBuildServicesCompleteResult{
		Items: items,
	}
	return out, nil
}
