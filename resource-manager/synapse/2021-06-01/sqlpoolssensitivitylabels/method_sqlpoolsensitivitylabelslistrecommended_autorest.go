package sqlpoolssensitivitylabels

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

type SqlPoolSensitivityLabelsListRecommendedOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SensitivityLabel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolSensitivityLabelsListRecommendedOperationResponse, error)
}

type SqlPoolSensitivityLabelsListRecommendedCompleteResult struct {
	Items []SensitivityLabel
}

func (r SqlPoolSensitivityLabelsListRecommendedOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolSensitivityLabelsListRecommendedOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolSensitivityLabelsListRecommendedOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type SqlPoolSensitivityLabelsListRecommendedOperationOptions struct {
	Filter                         *string
	IncludeDisabledRecommendations *bool
}

func DefaultSqlPoolSensitivityLabelsListRecommendedOperationOptions() SqlPoolSensitivityLabelsListRecommendedOperationOptions {
	return SqlPoolSensitivityLabelsListRecommendedOperationOptions{}
}

func (o SqlPoolSensitivityLabelsListRecommendedOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolSensitivityLabelsListRecommendedOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.IncludeDisabledRecommendations != nil {
		out["includeDisabledRecommendations"] = *o.IncludeDisabledRecommendations
	}

	return out
}

// SqlPoolSensitivityLabelsListRecommended ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommended(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions) (resp SqlPoolSensitivityLabelsListRecommendedOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsListRecommended(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolSensitivityLabelsListRecommended(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolSensitivityLabelsListRecommended prepares the SqlPoolSensitivityLabelsListRecommended request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsListRecommended(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/recommendedSensitivityLabels", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolSensitivityLabelsListRecommendedWithNextLink prepares the SqlPoolSensitivityLabelsListRecommended request with the given nextLink token.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsListRecommendedWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolSensitivityLabelsListRecommended handles the response to the SqlPoolSensitivityLabelsListRecommended request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsListRecommended(resp *http.Response) (result SqlPoolSensitivityLabelsListRecommendedOperationResponse, err error) {
	type page struct {
		Values   []SensitivityLabel `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolSensitivityLabelsListRecommendedOperationResponse, err error) {
			req, err := c.preparerForSqlPoolSensitivityLabelsListRecommendedWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolSensitivityLabelsListRecommended(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListRecommended", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolSensitivityLabelsListRecommendedComplete retrieves all of the results into a single object
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommendedComplete(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions) (SqlPoolSensitivityLabelsListRecommendedCompleteResult, error) {
	return c.SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListRecommendedCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListRecommendedOperationOptions, predicate SensitivityLabelOperationPredicate) (resp SqlPoolSensitivityLabelsListRecommendedCompleteResult, err error) {
	items := make([]SensitivityLabel, 0)

	page, err := c.SqlPoolSensitivityLabelsListRecommended(ctx, id, options)
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

	out := SqlPoolSensitivityLabelsListRecommendedCompleteResult{
		Items: items,
	}
	return out, nil
}
