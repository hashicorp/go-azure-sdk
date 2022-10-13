package appplatform

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

type GatewayCustomDomainsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GatewayCustomDomainResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GatewayCustomDomainsListOperationResponse, error)
}

type GatewayCustomDomainsListCompleteResult struct {
	Items []GatewayCustomDomainResource
}

func (r GatewayCustomDomainsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GatewayCustomDomainsListOperationResponse) LoadMore(ctx context.Context) (resp GatewayCustomDomainsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GatewayCustomDomainsList ...
func (c AppPlatformClient) GatewayCustomDomainsList(ctx context.Context, id GatewayId) (resp GatewayCustomDomainsListOperationResponse, err error) {
	req, err := c.preparerForGatewayCustomDomainsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGatewayCustomDomainsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGatewayCustomDomainsList prepares the GatewayCustomDomainsList request.
func (c AppPlatformClient) preparerForGatewayCustomDomainsList(ctx context.Context, id GatewayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/domains", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGatewayCustomDomainsListWithNextLink prepares the GatewayCustomDomainsList request with the given nextLink token.
func (c AppPlatformClient) preparerForGatewayCustomDomainsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGatewayCustomDomainsList handles the response to the GatewayCustomDomainsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewayCustomDomainsList(resp *http.Response) (result GatewayCustomDomainsListOperationResponse, err error) {
	type page struct {
		Values   []GatewayCustomDomainResource `json:"value"`
		NextLink *string                       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GatewayCustomDomainsListOperationResponse, err error) {
			req, err := c.preparerForGatewayCustomDomainsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGatewayCustomDomainsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GatewayCustomDomainsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) GatewayCustomDomainsListComplete(ctx context.Context, id GatewayId) (GatewayCustomDomainsListCompleteResult, error) {
	return c.GatewayCustomDomainsListCompleteMatchingPredicate(ctx, id, GatewayCustomDomainResourceOperationPredicate{})
}

// GatewayCustomDomainsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) GatewayCustomDomainsListCompleteMatchingPredicate(ctx context.Context, id GatewayId, predicate GatewayCustomDomainResourceOperationPredicate) (resp GatewayCustomDomainsListCompleteResult, err error) {
	items := make([]GatewayCustomDomainResource, 0)

	page, err := c.GatewayCustomDomainsList(ctx, id)
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

	out := GatewayCustomDomainsListCompleteResult{
		Items: items,
	}
	return out, nil
}
