package appservicecertificateorders

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

type ListCertificatesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AppServiceCertificateResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListCertificatesOperationResponse, error)
}

type ListCertificatesCompleteResult struct {
	Items []AppServiceCertificateResource
}

func (r ListCertificatesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListCertificatesOperationResponse) LoadMore(ctx context.Context) (resp ListCertificatesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListCertificates ...
func (c AppServiceCertificateOrdersClient) ListCertificates(ctx context.Context, id CertificateOrderId) (resp ListCertificatesOperationResponse, err error) {
	req, err := c.preparerForListCertificates(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListCertificates(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListCertificates prepares the ListCertificates request.
func (c AppServiceCertificateOrdersClient) preparerForListCertificates(ctx context.Context, id CertificateOrderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/certificates", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListCertificatesWithNextLink prepares the ListCertificates request with the given nextLink token.
func (c AppServiceCertificateOrdersClient) preparerForListCertificatesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListCertificates handles the response to the ListCertificates request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForListCertificates(resp *http.Response) (result ListCertificatesOperationResponse, err error) {
	type page struct {
		Values   []AppServiceCertificateResource `json:"value"`
		NextLink *string                         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListCertificatesOperationResponse, err error) {
			req, err := c.preparerForListCertificatesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListCertificates(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "ListCertificates", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListCertificatesComplete retrieves all of the results into a single object
func (c AppServiceCertificateOrdersClient) ListCertificatesComplete(ctx context.Context, id CertificateOrderId) (ListCertificatesCompleteResult, error) {
	return c.ListCertificatesCompleteMatchingPredicate(ctx, id, AppServiceCertificateResourceOperationPredicate{})
}

// ListCertificatesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceCertificateOrdersClient) ListCertificatesCompleteMatchingPredicate(ctx context.Context, id CertificateOrderId, predicate AppServiceCertificateResourceOperationPredicate) (resp ListCertificatesCompleteResult, err error) {
	items := make([]AppServiceCertificateResource, 0)

	page, err := c.ListCertificates(ctx, id)
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

	out := ListCertificatesCompleteResult{
		Items: items,
	}
	return out, nil
}
