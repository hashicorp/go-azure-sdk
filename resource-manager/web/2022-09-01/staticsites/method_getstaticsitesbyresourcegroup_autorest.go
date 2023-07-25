package staticsites

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetStaticSitesByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetStaticSitesByResourceGroupOperationResponse, error)
}

type GetStaticSitesByResourceGroupCompleteResult struct {
	Items []StaticSiteARMResource
}

func (r GetStaticSitesByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetStaticSitesByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp GetStaticSitesByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetStaticSitesByResourceGroup ...
func (c StaticSitesClient) GetStaticSitesByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (resp GetStaticSitesByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForGetStaticSitesByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetStaticSitesByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetStaticSitesByResourceGroup prepares the GetStaticSitesByResourceGroup request.
func (c StaticSitesClient) preparerForGetStaticSitesByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/staticSites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetStaticSitesByResourceGroupWithNextLink prepares the GetStaticSitesByResourceGroup request with the given nextLink token.
func (c StaticSitesClient) preparerForGetStaticSitesByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetStaticSitesByResourceGroup handles the response to the GetStaticSitesByResourceGroup request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetStaticSitesByResourceGroup(resp *http.Response) (result GetStaticSitesByResourceGroupOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteARMResource `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetStaticSitesByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForGetStaticSitesByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetStaticSitesByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSitesByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetStaticSitesByResourceGroupComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetStaticSitesByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (GetStaticSitesByResourceGroupCompleteResult, error) {
	return c.GetStaticSitesByResourceGroupCompleteMatchingPredicate(ctx, id, StaticSiteARMResourceOperationPredicate{})
}

// GetStaticSitesByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetStaticSitesByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate StaticSiteARMResourceOperationPredicate) (resp GetStaticSitesByResourceGroupCompleteResult, err error) {
	items := make([]StaticSiteARMResource, 0)

	page, err := c.GetStaticSitesByResourceGroup(ctx, id)
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

	out := GetStaticSitesByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
