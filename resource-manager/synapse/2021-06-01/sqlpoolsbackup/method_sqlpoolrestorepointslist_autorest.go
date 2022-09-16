package sqlpoolsbackup

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

type SqlPoolRestorePointsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RestorePoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolRestorePointsListOperationResponse, error)
}

type SqlPoolRestorePointsListCompleteResult struct {
	Items []RestorePoint
}

func (r SqlPoolRestorePointsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolRestorePointsListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolRestorePointsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolRestorePointsList ...
func (c SqlPoolsBackupClient) SqlPoolRestorePointsList(ctx context.Context, id SqlPoolId) (resp SqlPoolRestorePointsListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolRestorePointsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolRestorePointsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolRestorePointsList prepares the SqlPoolRestorePointsList request.
func (c SqlPoolsBackupClient) preparerForSqlPoolRestorePointsList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restorePoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolRestorePointsListWithNextLink prepares the SqlPoolRestorePointsList request with the given nextLink token.
func (c SqlPoolsBackupClient) preparerForSqlPoolRestorePointsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolRestorePointsList handles the response to the SqlPoolRestorePointsList request. The method always
// closes the http.Response Body.
func (c SqlPoolsBackupClient) responderForSqlPoolRestorePointsList(resp *http.Response) (result SqlPoolRestorePointsListOperationResponse, err error) {
	type page struct {
		Values   []RestorePoint `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolRestorePointsListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolRestorePointsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolRestorePointsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsbackup.SqlPoolsBackupClient", "SqlPoolRestorePointsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolRestorePointsListComplete retrieves all of the results into a single object
func (c SqlPoolsBackupClient) SqlPoolRestorePointsListComplete(ctx context.Context, id SqlPoolId) (SqlPoolRestorePointsListCompleteResult, error) {
	return c.SqlPoolRestorePointsListCompleteMatchingPredicate(ctx, id, RestorePointOperationPredicate{})
}

// SqlPoolRestorePointsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsBackupClient) SqlPoolRestorePointsListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate RestorePointOperationPredicate) (resp SqlPoolRestorePointsListCompleteResult, err error) {
	items := make([]RestorePoint, 0)

	page, err := c.SqlPoolRestorePointsList(ctx, id)
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

	out := SqlPoolRestorePointsListCompleteResult{
		Items: items,
	}
	return out, nil
}
