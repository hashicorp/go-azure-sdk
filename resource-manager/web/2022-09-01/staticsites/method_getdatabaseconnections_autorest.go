package staticsites

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

type GetDatabaseConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DatabaseConnection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetDatabaseConnectionsOperationResponse, error)
}

type GetDatabaseConnectionsCompleteResult struct {
	Items []DatabaseConnection
}

func (r GetDatabaseConnectionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetDatabaseConnectionsOperationResponse) LoadMore(ctx context.Context) (resp GetDatabaseConnectionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetDatabaseConnections ...
func (c StaticSitesClient) GetDatabaseConnections(ctx context.Context, id StaticSiteId) (resp GetDatabaseConnectionsOperationResponse, err error) {
	req, err := c.preparerForGetDatabaseConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetDatabaseConnections(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetDatabaseConnections prepares the GetDatabaseConnections request.
func (c StaticSitesClient) preparerForGetDatabaseConnections(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/databaseConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetDatabaseConnectionsWithNextLink prepares the GetDatabaseConnections request with the given nextLink token.
func (c StaticSitesClient) preparerForGetDatabaseConnectionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetDatabaseConnections handles the response to the GetDatabaseConnections request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetDatabaseConnections(resp *http.Response) (result GetDatabaseConnectionsOperationResponse, err error) {
	type page struct {
		Values   []DatabaseConnection `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetDatabaseConnectionsOperationResponse, err error) {
			req, err := c.preparerForGetDatabaseConnectionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetDatabaseConnections(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetDatabaseConnections", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetDatabaseConnectionsComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetDatabaseConnectionsComplete(ctx context.Context, id StaticSiteId) (GetDatabaseConnectionsCompleteResult, error) {
	return c.GetDatabaseConnectionsCompleteMatchingPredicate(ctx, id, DatabaseConnectionOperationPredicate{})
}

// GetDatabaseConnectionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetDatabaseConnectionsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate DatabaseConnectionOperationPredicate) (resp GetDatabaseConnectionsCompleteResult, err error) {
	items := make([]DatabaseConnection, 0)

	page, err := c.GetDatabaseConnections(ctx, id)
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

	out := GetDatabaseConnectionsCompleteResult{
		Items: items,
	}
	return out, nil
}
