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

type ListSiteDetectorsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DetectorDefinitionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteDetectorsOperationResponse, error)
}

type ListSiteDetectorsCompleteResult struct {
	Items []DetectorDefinitionResource
}

func (r ListSiteDetectorsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteDetectorsOperationResponse) LoadMore(ctx context.Context) (resp ListSiteDetectorsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteDetectors ...
func (c DiagnosticsClient) ListSiteDetectors(ctx context.Context, id DiagnosticId) (resp ListSiteDetectorsOperationResponse, err error) {
	req, err := c.preparerForListSiteDetectors(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteDetectors(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteDetectors prepares the ListSiteDetectors request.
func (c DiagnosticsClient) preparerForListSiteDetectors(ctx context.Context, id DiagnosticId) (*http.Request, error) {
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

// preparerForListSiteDetectorsWithNextLink prepares the ListSiteDetectors request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteDetectorsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteDetectors handles the response to the ListSiteDetectors request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteDetectors(resp *http.Response) (result ListSiteDetectorsOperationResponse, err error) {
	type page struct {
		Values   []DetectorDefinitionResource `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteDetectorsOperationResponse, err error) {
			req, err := c.preparerForListSiteDetectorsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteDetectors(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteDetectors", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteDetectorsComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteDetectorsComplete(ctx context.Context, id DiagnosticId) (ListSiteDetectorsCompleteResult, error) {
	return c.ListSiteDetectorsCompleteMatchingPredicate(ctx, id, DetectorDefinitionResourceOperationPredicate{})
}

// ListSiteDetectorsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteDetectorsCompleteMatchingPredicate(ctx context.Context, id DiagnosticId, predicate DetectorDefinitionResourceOperationPredicate) (resp ListSiteDetectorsCompleteResult, err error) {
	items := make([]DetectorDefinitionResource, 0)

	page, err := c.ListSiteDetectors(ctx, id)
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

	out := ListSiteDetectorsCompleteResult{
		Items: items,
	}
	return out, nil
}
