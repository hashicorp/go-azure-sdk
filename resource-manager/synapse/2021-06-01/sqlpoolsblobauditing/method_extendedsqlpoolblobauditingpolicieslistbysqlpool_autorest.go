package sqlpoolsblobauditing

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

type ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ExtendedSqlPoolBlobAuditingPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, error)
}

type ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult struct {
	Items []ExtendedSqlPoolBlobAuditingPolicy
}

func (r ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse) LoadMore(ctx context.Context) (resp ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool ...
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx context.Context, id SqlPoolId) (resp ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
	req, err := c.preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool prepares the ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool request.
func (c SqlPoolsBlobAuditingClient) preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extendedAuditingSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolWithNextLink prepares the ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool request with the given nextLink token.
func (c SqlPoolsBlobAuditingClient) preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool handles the response to the ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool request. The method always
// closes the http.Response Body.
func (c SqlPoolsBlobAuditingClient) responderForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(resp *http.Response) (result ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
	type page struct {
		Values   []ExtendedSqlPoolBlobAuditingPolicy `json:"value"`
		NextLink *string                             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
			req, err := c.preparerForExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsblobauditing.SqlPoolsBlobAuditingClient", "ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolComplete retrieves all of the results into a single object
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolComplete(ctx context.Context, id SqlPoolId) (ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, error) {
	return c.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx, id, ExtendedSqlPoolBlobAuditingPolicyOperationPredicate{})
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate ExtendedSqlPoolBlobAuditingPolicyOperationPredicate) (resp ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, err error) {
	items := make([]ExtendedSqlPoolBlobAuditingPolicy, 0)

	page, err := c.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx, id)
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

	out := ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult{
		Items: items,
	}
	return out, nil
}
