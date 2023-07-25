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

type ListProductionSiteDeploymentStatusesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmDeploymentStatus

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListProductionSiteDeploymentStatusesOperationResponse, error)
}

type ListProductionSiteDeploymentStatusesCompleteResult struct {
	Items []CsmDeploymentStatus
}

func (r ListProductionSiteDeploymentStatusesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListProductionSiteDeploymentStatusesOperationResponse) LoadMore(ctx context.Context) (resp ListProductionSiteDeploymentStatusesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListProductionSiteDeploymentStatuses ...
func (c WebAppsClient) ListProductionSiteDeploymentStatuses(ctx context.Context, id SiteId) (resp ListProductionSiteDeploymentStatusesOperationResponse, err error) {
	req, err := c.preparerForListProductionSiteDeploymentStatuses(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListProductionSiteDeploymentStatuses(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListProductionSiteDeploymentStatuses prepares the ListProductionSiteDeploymentStatuses request.
func (c WebAppsClient) preparerForListProductionSiteDeploymentStatuses(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListProductionSiteDeploymentStatusesWithNextLink prepares the ListProductionSiteDeploymentStatuses request with the given nextLink token.
func (c WebAppsClient) preparerForListProductionSiteDeploymentStatusesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListProductionSiteDeploymentStatuses handles the response to the ListProductionSiteDeploymentStatuses request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListProductionSiteDeploymentStatuses(resp *http.Response) (result ListProductionSiteDeploymentStatusesOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListProductionSiteDeploymentStatusesOperationResponse, err error) {
			req, err := c.preparerForListProductionSiteDeploymentStatusesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListProductionSiteDeploymentStatuses(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListProductionSiteDeploymentStatuses", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListProductionSiteDeploymentStatusesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListProductionSiteDeploymentStatusesComplete(ctx context.Context, id SiteId) (ListProductionSiteDeploymentStatusesCompleteResult, error) {
	return c.ListProductionSiteDeploymentStatusesCompleteMatchingPredicate(ctx, id, CsmDeploymentStatusOperationPredicate{})
}

// ListProductionSiteDeploymentStatusesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListProductionSiteDeploymentStatusesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate CsmDeploymentStatusOperationPredicate) (resp ListProductionSiteDeploymentStatusesCompleteResult, err error) {
	items := make([]CsmDeploymentStatus, 0)

	page, err := c.ListProductionSiteDeploymentStatuses(ctx, id)
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

	out := ListProductionSiteDeploymentStatusesCompleteResult{
		Items: items,
	}
	return out, nil
}
