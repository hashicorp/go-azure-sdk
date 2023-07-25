package diagnostics

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

type ListSiteDetectorResponsesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DetectorResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteDetectorResponsesOperationResponse, error)
}

type ListSiteDetectorResponsesCompleteResult struct {
	Items []DetectorResponse
}

func (r ListSiteDetectorResponsesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteDetectorResponsesOperationResponse) LoadMore(ctx context.Context) (resp ListSiteDetectorResponsesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteDetectorResponses ...
func (c DiagnosticsClient) ListSiteDetectorResponses(ctx context.Context, id SiteId) (resp ListSiteDetectorResponsesOperationResponse, err error) {
	req, err := c.preparerForListSiteDetectorResponses(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteDetectorResponses(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteDetectorResponses prepares the ListSiteDetectorResponses request.
func (c DiagnosticsClient) preparerForListSiteDetectorResponses(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectors", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteDetectorResponsesWithNextLink prepares the ListSiteDetectorResponses request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteDetectorResponsesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteDetectorResponses handles the response to the ListSiteDetectorResponses request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteDetectorResponses(resp *http.Response) (result ListSiteDetectorResponsesOperationResponse, err error) {
	type page struct {
		Values   []DetectorResponse `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteDetectorResponsesOperationResponse, err error) {
			req, err := c.preparerForListSiteDetectorResponsesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteDetectorResponses(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectorResponses", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteDetectorResponsesComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteDetectorResponsesComplete(ctx context.Context, id SiteId) (ListSiteDetectorResponsesCompleteResult, error) {
	return c.ListSiteDetectorResponsesCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListSiteDetectorResponsesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteDetectorResponsesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate DetectorResponseOperationPredicate) (resp ListSiteDetectorResponsesCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	page, err := c.ListSiteDetectorResponses(ctx, id)
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

	out := ListSiteDetectorResponsesCompleteResult{
		Items: items,
	}
	return out, nil
}
