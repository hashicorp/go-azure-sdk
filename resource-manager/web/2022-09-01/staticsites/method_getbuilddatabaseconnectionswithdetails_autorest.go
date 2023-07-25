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

type GetBuildDatabaseConnectionsWithDetailsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DatabaseConnection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetBuildDatabaseConnectionsWithDetailsOperationResponse, error)
}

type GetBuildDatabaseConnectionsWithDetailsCompleteResult struct {
	Items []DatabaseConnection
}

func (r GetBuildDatabaseConnectionsWithDetailsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetBuildDatabaseConnectionsWithDetailsOperationResponse) LoadMore(ctx context.Context) (resp GetBuildDatabaseConnectionsWithDetailsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetBuildDatabaseConnectionsWithDetails ...
func (c StaticSitesClient) GetBuildDatabaseConnectionsWithDetails(ctx context.Context, id BuildId) (resp GetBuildDatabaseConnectionsWithDetailsOperationResponse, err error) {
	req, err := c.preparerForGetBuildDatabaseConnectionsWithDetails(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetBuildDatabaseConnectionsWithDetails(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetBuildDatabaseConnectionsWithDetails prepares the GetBuildDatabaseConnectionsWithDetails request.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnectionsWithDetails(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/showDatabaseConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetBuildDatabaseConnectionsWithDetailsWithNextLink prepares the GetBuildDatabaseConnectionsWithDetails request with the given nextLink token.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnectionsWithDetailsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBuildDatabaseConnectionsWithDetails handles the response to the GetBuildDatabaseConnectionsWithDetails request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetBuildDatabaseConnectionsWithDetails(resp *http.Response) (result GetBuildDatabaseConnectionsWithDetailsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetBuildDatabaseConnectionsWithDetailsOperationResponse, err error) {
			req, err := c.preparerForGetBuildDatabaseConnectionsWithDetailsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetBuildDatabaseConnectionsWithDetails(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnectionsWithDetails", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetBuildDatabaseConnectionsWithDetailsComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetBuildDatabaseConnectionsWithDetailsComplete(ctx context.Context, id BuildId) (GetBuildDatabaseConnectionsWithDetailsCompleteResult, error) {
	return c.GetBuildDatabaseConnectionsWithDetailsCompleteMatchingPredicate(ctx, id, DatabaseConnectionOperationPredicate{})
}

// GetBuildDatabaseConnectionsWithDetailsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetBuildDatabaseConnectionsWithDetailsCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate DatabaseConnectionOperationPredicate) (resp GetBuildDatabaseConnectionsWithDetailsCompleteResult, err error) {
	items := make([]DatabaseConnection, 0)

	page, err := c.GetBuildDatabaseConnectionsWithDetails(ctx, id)
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

	out := GetBuildDatabaseConnectionsWithDetailsCompleteResult{
		Items: items,
	}
	return out, nil
}
