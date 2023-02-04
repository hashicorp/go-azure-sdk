package migrates

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

type VMwareSitesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMwareSite

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (VMwareSitesListOperationResponse, error)
}

type VMwareSitesListCompleteResult struct {
	Items []VMwareSite
}

func (r VMwareSitesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r VMwareSitesListOperationResponse) LoadMore(ctx context.Context) (resp VMwareSitesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// VMwareSitesList ...
func (c MigratesClient) VMwareSitesList(ctx context.Context, id commonids.ResourceGroupId) (resp VMwareSitesListOperationResponse, err error) {
	req, err := c.preparerForVMwareSitesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForVMwareSitesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForVMwareSitesList prepares the VMwareSitesList request.
func (c MigratesClient) preparerForVMwareSitesList(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.OffAzure/vmwareSites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForVMwareSitesListWithNextLink prepares the VMwareSitesList request with the given nextLink token.
func (c MigratesClient) preparerForVMwareSitesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForVMwareSitesList handles the response to the VMwareSitesList request. The method always
// closes the http.Response Body.
func (c MigratesClient) responderForVMwareSitesList(resp *http.Response) (result VMwareSitesListOperationResponse, err error) {
	type page struct {
		Values   []VMwareSite `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result VMwareSitesListOperationResponse, err error) {
			req, err := c.preparerForVMwareSitesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForVMwareSitesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// VMwareSitesListComplete retrieves all of the results into a single object
func (c MigratesClient) VMwareSitesListComplete(ctx context.Context, id commonids.ResourceGroupId) (VMwareSitesListCompleteResult, error) {
	return c.VMwareSitesListCompleteMatchingPredicate(ctx, id, VMwareSiteOperationPredicate{})
}

// VMwareSitesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MigratesClient) VMwareSitesListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate VMwareSiteOperationPredicate) (resp VMwareSitesListCompleteResult, err error) {
	items := make([]VMwareSite, 0)

	page, err := c.VMwareSitesList(ctx, id)
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

	out := VMwareSitesListCompleteResult{
		Items: items,
	}
	return out, nil
}
