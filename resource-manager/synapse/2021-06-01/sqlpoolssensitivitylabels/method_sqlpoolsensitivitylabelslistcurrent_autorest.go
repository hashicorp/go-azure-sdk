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

type SqlPoolSensitivityLabelsListCurrentOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SensitivityLabel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolSensitivityLabelsListCurrentOperationResponse, error)
}

type SqlPoolSensitivityLabelsListCurrentCompleteResult struct {
	Items []SensitivityLabel
}

func (r SqlPoolSensitivityLabelsListCurrentOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolSensitivityLabelsListCurrentOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolSensitivityLabelsListCurrentOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type SqlPoolSensitivityLabelsListCurrentOperationOptions struct {
	Filter *string
}

func DefaultSqlPoolSensitivityLabelsListCurrentOperationOptions() SqlPoolSensitivityLabelsListCurrentOperationOptions {
	return SqlPoolSensitivityLabelsListCurrentOperationOptions{}
}

func (o SqlPoolSensitivityLabelsListCurrentOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolSensitivityLabelsListCurrentOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// SqlPoolSensitivityLabelsListCurrent ...
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrent(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions) (resp SqlPoolSensitivityLabelsListCurrentOperationResponse, err error) {
	req, err := c.preparerForSqlPoolSensitivityLabelsListCurrent(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolSensitivityLabelsListCurrent(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolSensitivityLabelsListCurrent prepares the SqlPoolSensitivityLabelsListCurrent request.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsListCurrent(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/currentSensitivityLabels", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolSensitivityLabelsListCurrentWithNextLink prepares the SqlPoolSensitivityLabelsListCurrent request with the given nextLink token.
func (c SqlPoolsSensitivityLabelsClient) preparerForSqlPoolSensitivityLabelsListCurrentWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolSensitivityLabelsListCurrent handles the response to the SqlPoolSensitivityLabelsListCurrent request. The method always
// closes the http.Response Body.
func (c SqlPoolsSensitivityLabelsClient) responderForSqlPoolSensitivityLabelsListCurrent(resp *http.Response) (result SqlPoolSensitivityLabelsListCurrentOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolSensitivityLabelsListCurrentOperationResponse, err error) {
			req, err := c.preparerForSqlPoolSensitivityLabelsListCurrentWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolSensitivityLabelsListCurrent(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient", "SqlPoolSensitivityLabelsListCurrent", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolSensitivityLabelsListCurrentComplete retrieves all of the results into a single object
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrentComplete(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions) (SqlPoolSensitivityLabelsListCurrentCompleteResult, error) {
	return c.SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsSensitivityLabelsClient) SqlPoolSensitivityLabelsListCurrentCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, options SqlPoolSensitivityLabelsListCurrentOperationOptions, predicate SensitivityLabelOperationPredicate) (resp SqlPoolSensitivityLabelsListCurrentCompleteResult, err error) {
	items := make([]SensitivityLabel, 0)

	page, err := c.SqlPoolSensitivityLabelsListCurrent(ctx, id, options)
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

	out := SqlPoolSensitivityLabelsListCurrentCompleteResult{
		Items: items,
	}
	return out, nil
}
