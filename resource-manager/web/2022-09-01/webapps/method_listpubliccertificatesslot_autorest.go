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

type ListPublicCertificatesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PublicCertificate

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPublicCertificatesSlotOperationResponse, error)
}

type ListPublicCertificatesSlotCompleteResult struct {
	Items []PublicCertificate
}

func (r ListPublicCertificatesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPublicCertificatesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListPublicCertificatesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListPublicCertificatesSlot ...
func (c WebAppsClient) ListPublicCertificatesSlot(ctx context.Context, id SlotId) (resp ListPublicCertificatesSlotOperationResponse, err error) {
	req, err := c.preparerForListPublicCertificatesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPublicCertificatesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPublicCertificatesSlot prepares the ListPublicCertificatesSlot request.
func (c WebAppsClient) preparerForListPublicCertificatesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/publicCertificates", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListPublicCertificatesSlotWithNextLink prepares the ListPublicCertificatesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListPublicCertificatesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPublicCertificatesSlot handles the response to the ListPublicCertificatesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPublicCertificatesSlot(resp *http.Response) (result ListPublicCertificatesSlotOperationResponse, err error) {
	type page struct {
		Values   []PublicCertificate `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPublicCertificatesSlotOperationResponse, err error) {
			req, err := c.preparerForListPublicCertificatesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPublicCertificatesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificatesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPublicCertificatesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListPublicCertificatesSlotComplete(ctx context.Context, id SlotId) (ListPublicCertificatesSlotCompleteResult, error) {
	return c.ListPublicCertificatesSlotCompleteMatchingPredicate(ctx, id, PublicCertificateOperationPredicate{})
}

// ListPublicCertificatesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListPublicCertificatesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate PublicCertificateOperationPredicate) (resp ListPublicCertificatesSlotCompleteResult, err error) {
	items := make([]PublicCertificate, 0)

	page, err := c.ListPublicCertificatesSlot(ctx, id)
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

	out := ListPublicCertificatesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
