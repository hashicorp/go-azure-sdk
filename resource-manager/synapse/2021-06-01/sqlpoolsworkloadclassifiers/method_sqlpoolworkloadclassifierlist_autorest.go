package sqlpoolsworkloadclassifiers

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

type SqlPoolWorkloadClassifierListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadClassifier

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolWorkloadClassifierListOperationResponse, error)
}

type SqlPoolWorkloadClassifierListCompleteResult struct {
	Items []WorkloadClassifier
}

func (r SqlPoolWorkloadClassifierListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolWorkloadClassifierListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolWorkloadClassifierListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolWorkloadClassifierList ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierList(ctx context.Context, id WorkloadGroupId) (resp SqlPoolWorkloadClassifierListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadClassifierList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolWorkloadClassifierList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolWorkloadClassifierList prepares the SqlPoolWorkloadClassifierList request.
func (c SqlPoolsWorkloadClassifiersClient) preparerForSqlPoolWorkloadClassifierList(ctx context.Context, id WorkloadGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadClassifiers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolWorkloadClassifierListWithNextLink prepares the SqlPoolWorkloadClassifierList request with the given nextLink token.
func (c SqlPoolsWorkloadClassifiersClient) preparerForSqlPoolWorkloadClassifierListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolWorkloadClassifierList handles the response to the SqlPoolWorkloadClassifierList request. The method always
// closes the http.Response Body.
func (c SqlPoolsWorkloadClassifiersClient) responderForSqlPoolWorkloadClassifierList(resp *http.Response) (result SqlPoolWorkloadClassifierListOperationResponse, err error) {
	type page struct {
		Values   []WorkloadClassifier `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolWorkloadClassifierListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolWorkloadClassifierListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolWorkloadClassifierList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient", "SqlPoolWorkloadClassifierList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolWorkloadClassifierListComplete retrieves all of the results into a single object
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierListComplete(ctx context.Context, id WorkloadGroupId) (SqlPoolWorkloadClassifierListCompleteResult, error) {
	return c.SqlPoolWorkloadClassifierListCompleteMatchingPredicate(ctx, id, WorkloadClassifierOperationPredicate{})
}

// SqlPoolWorkloadClassifierListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierListCompleteMatchingPredicate(ctx context.Context, id WorkloadGroupId, predicate WorkloadClassifierOperationPredicate) (resp SqlPoolWorkloadClassifierListCompleteResult, err error) {
	items := make([]WorkloadClassifier, 0)

	page, err := c.SqlPoolWorkloadClassifierList(ctx, id)
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

	out := SqlPoolWorkloadClassifierListCompleteResult{
		Items: items,
	}
	return out, nil
}
