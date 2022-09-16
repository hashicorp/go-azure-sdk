package sqlpoolsusages

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

type SqlPoolUsagesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SqlPoolUsage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolUsagesListOperationResponse, error)
}

type SqlPoolUsagesListCompleteResult struct {
	Items []SqlPoolUsage
}

func (r SqlPoolUsagesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolUsagesListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolUsagesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolUsagesList ...
func (c SqlPoolsUsagesClient) SqlPoolUsagesList(ctx context.Context, id SqlPoolId) (resp SqlPoolUsagesListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolUsagesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolUsagesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolUsagesList prepares the SqlPoolUsagesList request.
func (c SqlPoolsUsagesClient) preparerForSqlPoolUsagesList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolUsagesListWithNextLink prepares the SqlPoolUsagesList request with the given nextLink token.
func (c SqlPoolsUsagesClient) preparerForSqlPoolUsagesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolUsagesList handles the response to the SqlPoolUsagesList request. The method always
// closes the http.Response Body.
func (c SqlPoolsUsagesClient) responderForSqlPoolUsagesList(resp *http.Response) (result SqlPoolUsagesListOperationResponse, err error) {
	type page struct {
		Values   []SqlPoolUsage `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolUsagesListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolUsagesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolUsagesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsusages.SqlPoolsUsagesClient", "SqlPoolUsagesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolUsagesListComplete retrieves all of the results into a single object
func (c SqlPoolsUsagesClient) SqlPoolUsagesListComplete(ctx context.Context, id SqlPoolId) (SqlPoolUsagesListCompleteResult, error) {
	return c.SqlPoolUsagesListCompleteMatchingPredicate(ctx, id, SqlPoolUsageOperationPredicate{})
}

// SqlPoolUsagesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsUsagesClient) SqlPoolUsagesListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolUsageOperationPredicate) (resp SqlPoolUsagesListCompleteResult, err error) {
	items := make([]SqlPoolUsage, 0)

	page, err := c.SqlPoolUsagesList(ctx, id)
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

	out := SqlPoolUsagesListCompleteResult{
		Items: items,
	}
	return out, nil
}
