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

type GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiKVReference

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, error)
}

type GetSiteConnectionStringKeyVaultReferencesSlotCompleteResult struct {
	Items []ApiKVReference
}

func (r GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse) LoadMore(ctx context.Context) (resp GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetSiteConnectionStringKeyVaultReferencesSlot ...
func (c WebAppsClient) GetSiteConnectionStringKeyVaultReferencesSlot(ctx context.Context, id SlotId) (resp GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, err error) {
	req, err := c.preparerForGetSiteConnectionStringKeyVaultReferencesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetSiteConnectionStringKeyVaultReferencesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetSiteConnectionStringKeyVaultReferencesSlot prepares the GetSiteConnectionStringKeyVaultReferencesSlot request.
func (c WebAppsClient) preparerForGetSiteConnectionStringKeyVaultReferencesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/configReferences/connectionStrings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetSiteConnectionStringKeyVaultReferencesSlotWithNextLink prepares the GetSiteConnectionStringKeyVaultReferencesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForGetSiteConnectionStringKeyVaultReferencesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetSiteConnectionStringKeyVaultReferencesSlot handles the response to the GetSiteConnectionStringKeyVaultReferencesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSiteConnectionStringKeyVaultReferencesSlot(resp *http.Response) (result GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, err error) {
	type page struct {
		Values   []ApiKVReference `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetSiteConnectionStringKeyVaultReferencesSlotOperationResponse, err error) {
			req, err := c.preparerForGetSiteConnectionStringKeyVaultReferencesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetSiteConnectionStringKeyVaultReferencesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSiteConnectionStringKeyVaultReferencesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetSiteConnectionStringKeyVaultReferencesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) GetSiteConnectionStringKeyVaultReferencesSlotComplete(ctx context.Context, id SlotId) (GetSiteConnectionStringKeyVaultReferencesSlotCompleteResult, error) {
	return c.GetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// GetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) GetSiteConnectionStringKeyVaultReferencesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ApiKVReferenceOperationPredicate) (resp GetSiteConnectionStringKeyVaultReferencesSlotCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	page, err := c.GetSiteConnectionStringKeyVaultReferencesSlot(ctx, id)
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

	out := GetSiteConnectionStringKeyVaultReferencesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
