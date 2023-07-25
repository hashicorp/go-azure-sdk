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

type GetAppSettingsKeyVaultReferencesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiKVReference

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAppSettingsKeyVaultReferencesOperationResponse, error)
}

type GetAppSettingsKeyVaultReferencesCompleteResult struct {
	Items []ApiKVReference
}

func (r GetAppSettingsKeyVaultReferencesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAppSettingsKeyVaultReferencesOperationResponse) LoadMore(ctx context.Context) (resp GetAppSettingsKeyVaultReferencesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetAppSettingsKeyVaultReferences ...
func (c WebAppsClient) GetAppSettingsKeyVaultReferences(ctx context.Context, id SiteId) (resp GetAppSettingsKeyVaultReferencesOperationResponse, err error) {
	req, err := c.preparerForGetAppSettingsKeyVaultReferences(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAppSettingsKeyVaultReferences(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAppSettingsKeyVaultReferences prepares the GetAppSettingsKeyVaultReferences request.
func (c WebAppsClient) preparerForGetAppSettingsKeyVaultReferences(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/configReferences/appSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAppSettingsKeyVaultReferencesWithNextLink prepares the GetAppSettingsKeyVaultReferences request with the given nextLink token.
func (c WebAppsClient) preparerForGetAppSettingsKeyVaultReferencesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAppSettingsKeyVaultReferences handles the response to the GetAppSettingsKeyVaultReferences request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetAppSettingsKeyVaultReferences(resp *http.Response) (result GetAppSettingsKeyVaultReferencesOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAppSettingsKeyVaultReferencesOperationResponse, err error) {
			req, err := c.preparerForGetAppSettingsKeyVaultReferencesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAppSettingsKeyVaultReferences(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetAppSettingsKeyVaultReferences", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAppSettingsKeyVaultReferencesComplete retrieves all of the results into a single object
func (c WebAppsClient) GetAppSettingsKeyVaultReferencesComplete(ctx context.Context, id SiteId) (GetAppSettingsKeyVaultReferencesCompleteResult, error) {
	return c.GetAppSettingsKeyVaultReferencesCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// GetAppSettingsKeyVaultReferencesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) GetAppSettingsKeyVaultReferencesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate ApiKVReferenceOperationPredicate) (resp GetAppSettingsKeyVaultReferencesCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	page, err := c.GetAppSettingsKeyVaultReferences(ctx, id)
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

	out := GetAppSettingsKeyVaultReferencesCompleteResult{
		Items: items,
	}
	return out, nil
}
