package sqlpoolssecurityalertpolicies

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

type SqlPoolSecurityAlertPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SqlPoolSecurityAlertPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolSecurityAlertPoliciesListOperationResponse, error)
}

type SqlPoolSecurityAlertPoliciesListCompleteResult struct {
	Items []SqlPoolSecurityAlertPolicy
}

func (r SqlPoolSecurityAlertPoliciesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolSecurityAlertPoliciesListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolSecurityAlertPoliciesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolSecurityAlertPoliciesList ...
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesList(ctx context.Context, id SqlPoolId) (resp SqlPoolSecurityAlertPoliciesListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSecurityAlertPoliciesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolSecurityAlertPoliciesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolSecurityAlertPoliciesList prepares the SqlPoolSecurityAlertPoliciesList request.
func (c SqlPoolsSecurityAlertPoliciesClient) preparerForSqlPoolSecurityAlertPoliciesList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securityAlertPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolSecurityAlertPoliciesListWithNextLink prepares the SqlPoolSecurityAlertPoliciesList request with the given nextLink token.
func (c SqlPoolsSecurityAlertPoliciesClient) preparerForSqlPoolSecurityAlertPoliciesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolSecurityAlertPoliciesList handles the response to the SqlPoolSecurityAlertPoliciesList request. The method always
// closes the http.Response Body.
func (c SqlPoolsSecurityAlertPoliciesClient) responderForSqlPoolSecurityAlertPoliciesList(resp *http.Response) (result SqlPoolSecurityAlertPoliciesListOperationResponse, err error) {
	type page struct {
		Values   []SqlPoolSecurityAlertPolicy `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolSecurityAlertPoliciesListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolSecurityAlertPoliciesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolSecurityAlertPoliciesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient", "SqlPoolSecurityAlertPoliciesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolSecurityAlertPoliciesListComplete retrieves all of the results into a single object
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesListComplete(ctx context.Context, id SqlPoolId) (SqlPoolSecurityAlertPoliciesListCompleteResult, error) {
	return c.SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate(ctx, id, SqlPoolSecurityAlertPolicyOperationPredicate{})
}

// SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsSecurityAlertPoliciesClient) SqlPoolSecurityAlertPoliciesListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolSecurityAlertPolicyOperationPredicate) (resp SqlPoolSecurityAlertPoliciesListCompleteResult, err error) {
	items := make([]SqlPoolSecurityAlertPolicy, 0)

	page, err := c.SqlPoolSecurityAlertPoliciesList(ctx, id)
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

	out := SqlPoolSecurityAlertPoliciesListCompleteResult{
		Items: items,
	}
	return out, nil
}
