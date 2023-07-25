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

type ListSlotSiteDeploymentStatusesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmDeploymentStatus

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSlotSiteDeploymentStatusesSlotOperationResponse, error)
}

type ListSlotSiteDeploymentStatusesSlotCompleteResult struct {
	Items []CsmDeploymentStatus
}

func (r ListSlotSiteDeploymentStatusesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSlotSiteDeploymentStatusesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListSlotSiteDeploymentStatusesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSlotSiteDeploymentStatusesSlot ...
func (c WebAppsClient) ListSlotSiteDeploymentStatusesSlot(ctx context.Context, id SlotId) (resp ListSlotSiteDeploymentStatusesSlotOperationResponse, err error) {
	req, err := c.preparerForListSlotSiteDeploymentStatusesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSlotSiteDeploymentStatusesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSlotSiteDeploymentStatusesSlot prepares the ListSlotSiteDeploymentStatusesSlot request.
func (c WebAppsClient) preparerForListSlotSiteDeploymentStatusesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deploymentStatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSlotSiteDeploymentStatusesSlotWithNextLink prepares the ListSlotSiteDeploymentStatusesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListSlotSiteDeploymentStatusesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSlotSiteDeploymentStatusesSlot handles the response to the ListSlotSiteDeploymentStatusesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSlotSiteDeploymentStatusesSlot(resp *http.Response) (result ListSlotSiteDeploymentStatusesSlotOperationResponse, err error) {
	type page struct {
		Values   []CsmDeploymentStatus `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSlotSiteDeploymentStatusesSlotOperationResponse, err error) {
			req, err := c.preparerForListSlotSiteDeploymentStatusesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSlotSiteDeploymentStatusesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSlotSiteDeploymentStatusesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSlotSiteDeploymentStatusesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSlotSiteDeploymentStatusesSlotComplete(ctx context.Context, id SlotId) (ListSlotSiteDeploymentStatusesSlotCompleteResult, error) {
	return c.ListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate(ctx, id, CsmDeploymentStatusOperationPredicate{})
}

// ListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSlotSiteDeploymentStatusesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate CsmDeploymentStatusOperationPredicate) (resp ListSlotSiteDeploymentStatusesSlotCompleteResult, err error) {
	items := make([]CsmDeploymentStatus, 0)

	page, err := c.ListSlotSiteDeploymentStatusesSlot(ctx, id)
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

	out := ListSlotSiteDeploymentStatusesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
