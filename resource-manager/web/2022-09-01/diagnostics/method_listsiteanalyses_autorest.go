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

type ListSiteAnalysesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AnalysisDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSiteAnalysesOperationResponse, error)
}

type ListSiteAnalysesCompleteResult struct {
	Items []AnalysisDefinition
}

func (r ListSiteAnalysesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSiteAnalysesOperationResponse) LoadMore(ctx context.Context) (resp ListSiteAnalysesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSiteAnalyses ...
func (c DiagnosticsClient) ListSiteAnalyses(ctx context.Context, id DiagnosticId) (resp ListSiteAnalysesOperationResponse, err error) {
	req, err := c.preparerForListSiteAnalyses(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSiteAnalyses(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSiteAnalyses prepares the ListSiteAnalyses request.
func (c DiagnosticsClient) preparerForListSiteAnalyses(ctx context.Context, id DiagnosticId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/analyses", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSiteAnalysesWithNextLink prepares the ListSiteAnalyses request with the given nextLink token.
func (c DiagnosticsClient) preparerForListSiteAnalysesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSiteAnalyses handles the response to the ListSiteAnalyses request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListSiteAnalyses(resp *http.Response) (result ListSiteAnalysesOperationResponse, err error) {
	type page struct {
		Values   []AnalysisDefinition `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSiteAnalysesOperationResponse, err error) {
			req, err := c.preparerForListSiteAnalysesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSiteAnalyses(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListSiteAnalyses", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSiteAnalysesComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListSiteAnalysesComplete(ctx context.Context, id DiagnosticId) (ListSiteAnalysesCompleteResult, error) {
	return c.ListSiteAnalysesCompleteMatchingPredicate(ctx, id, AnalysisDefinitionOperationPredicate{})
}

// ListSiteAnalysesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListSiteAnalysesCompleteMatchingPredicate(ctx context.Context, id DiagnosticId, predicate AnalysisDefinitionOperationPredicate) (resp ListSiteAnalysesCompleteResult, err error) {
	items := make([]AnalysisDefinition, 0)

	page, err := c.ListSiteAnalyses(ctx, id)
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

	out := ListSiteAnalysesCompleteResult{
		Items: items,
	}
	return out, nil
}
