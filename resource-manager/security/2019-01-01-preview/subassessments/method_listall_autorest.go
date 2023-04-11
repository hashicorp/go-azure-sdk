package subassessments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAllOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecuritySubAssessment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAllOperationResponse, error)
}

type ListAllCompleteResult struct {
	Items []SecuritySubAssessment
}

func (r ListAllOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAllOperationResponse) LoadMore(ctx context.Context) (resp ListAllOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListAll ...
func (c SubAssessmentsClient) ListAll(ctx context.Context, id commonids.ScopeId) (resp ListAllOperationResponse, err error) {
	req, err := c.preparerForListAll(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAll(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAll prepares the ListAll request.
func (c SubAssessmentsClient) preparerForListAll(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/subAssessments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAllWithNextLink prepares the ListAll request with the given nextLink token.
func (c SubAssessmentsClient) preparerForListAllWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAll handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (c SubAssessmentsClient) responderForListAll(resp *http.Response) (result ListAllOperationResponse, err error) {
	type page struct {
		Values   []SecuritySubAssessment `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAllOperationResponse, err error) {
			req, err := c.preparerForListAllWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAll(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subassessments.SubAssessmentsClient", "ListAll", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAllComplete retrieves all of the results into a single object
func (c SubAssessmentsClient) ListAllComplete(ctx context.Context, id commonids.ScopeId) (ListAllCompleteResult, error) {
	return c.ListAllCompleteMatchingPredicate(ctx, id, SecuritySubAssessmentOperationPredicate{})
}

// ListAllCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SubAssessmentsClient) ListAllCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate SecuritySubAssessmentOperationPredicate) (resp ListAllCompleteResult, err error) {
	items := make([]SecuritySubAssessment, 0)

	page, err := c.ListAll(ctx, id)
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

	out := ListAllCompleteResult{
		Items: items,
	}
	return out, nil
}
