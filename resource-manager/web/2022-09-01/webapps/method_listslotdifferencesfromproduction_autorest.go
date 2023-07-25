package webapps

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

type ListSlotDifferencesFromProductionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SlotDifference

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSlotDifferencesFromProductionOperationResponse, error)
}

type ListSlotDifferencesFromProductionCompleteResult struct {
	Items []SlotDifference
}

func (r ListSlotDifferencesFromProductionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSlotDifferencesFromProductionOperationResponse) LoadMore(ctx context.Context) (resp ListSlotDifferencesFromProductionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSlotDifferencesFromProduction ...
func (c WebAppsClient) ListSlotDifferencesFromProduction(ctx context.Context, id SiteId, input CsmSlotEntity) (resp ListSlotDifferencesFromProductionOperationResponse, err error) {
	req, err := c.preparerForListSlotDifferencesFromProduction(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSlotDifferencesFromProduction(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSlotDifferencesFromProduction prepares the ListSlotDifferencesFromProduction request.
func (c WebAppsClient) preparerForListSlotDifferencesFromProduction(ctx context.Context, id SiteId, input CsmSlotEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/slotsdiffs", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSlotDifferencesFromProductionWithNextLink prepares the ListSlotDifferencesFromProduction request with the given nextLink token.
func (c WebAppsClient) preparerForListSlotDifferencesFromProductionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSlotDifferencesFromProduction handles the response to the ListSlotDifferencesFromProduction request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSlotDifferencesFromProduction(resp *http.Response) (result ListSlotDifferencesFromProductionOperationResponse, err error) {
	type page struct {
		Values   []SlotDifference `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSlotDifferencesFromProductionOperationResponse, err error) {
			req, err := c.preparerForListSlotDifferencesFromProductionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSlotDifferencesFromProduction(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotDifferencesFromProduction", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSlotDifferencesFromProductionComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSlotDifferencesFromProductionComplete(ctx context.Context, id SiteId, input CsmSlotEntity) (ListSlotDifferencesFromProductionCompleteResult, error) {
	return c.ListSlotDifferencesFromProductionCompleteMatchingPredicate(ctx, id, input, SlotDifferenceOperationPredicate{})
}

// ListSlotDifferencesFromProductionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSlotDifferencesFromProductionCompleteMatchingPredicate(ctx context.Context, id SiteId, input CsmSlotEntity, predicate SlotDifferenceOperationPredicate) (resp ListSlotDifferencesFromProductionCompleteResult, err error) {
	items := make([]SlotDifference, 0)

	page, err := c.ListSlotDifferencesFromProduction(ctx, id, input)
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

	out := ListSlotDifferencesFromProductionCompleteResult{
		Items: items,
	}
	return out, nil
}
