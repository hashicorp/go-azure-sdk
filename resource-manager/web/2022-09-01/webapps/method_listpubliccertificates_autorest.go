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

type ListPublicCertificatesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PublicCertificate

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPublicCertificatesOperationResponse, error)
}

type ListPublicCertificatesCompleteResult struct {
	Items []PublicCertificate
}

func (r ListPublicCertificatesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPublicCertificatesOperationResponse) LoadMore(ctx context.Context) (resp ListPublicCertificatesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListPublicCertificates ...
func (c WebAppsClient) ListPublicCertificates(ctx context.Context, id SiteId) (resp ListPublicCertificatesOperationResponse, err error) {
	req, err := c.preparerForListPublicCertificates(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPublicCertificates(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPublicCertificates prepares the ListPublicCertificates request.
func (c WebAppsClient) preparerForListPublicCertificates(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListPublicCertificatesWithNextLink prepares the ListPublicCertificates request with the given nextLink token.
func (c WebAppsClient) preparerForListPublicCertificatesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPublicCertificates handles the response to the ListPublicCertificates request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListPublicCertificates(resp *http.Response) (result ListPublicCertificatesOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPublicCertificatesOperationResponse, err error) {
			req, err := c.preparerForListPublicCertificatesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPublicCertificates(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublicCertificates", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPublicCertificatesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListPublicCertificatesComplete(ctx context.Context, id SiteId) (ListPublicCertificatesCompleteResult, error) {
	return c.ListPublicCertificatesCompleteMatchingPredicate(ctx, id, PublicCertificateOperationPredicate{})
}

// ListPublicCertificatesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListPublicCertificatesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate PublicCertificateOperationPredicate) (resp ListPublicCertificatesCompleteResult, err error) {
	items := make([]PublicCertificate, 0)

	page, err := c.ListPublicCertificates(ctx, id)
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

	out := ListPublicCertificatesCompleteResult{
		Items: items,
	}
	return out, nil
}
