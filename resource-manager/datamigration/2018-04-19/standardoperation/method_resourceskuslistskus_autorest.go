package standardoperation

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

type ResourceSkusListSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResourceSku

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ResourceSkusListSkusOperationResponse, error)
}

type ResourceSkusListSkusCompleteResult struct {
	Items []ResourceSku
}

func (r ResourceSkusListSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ResourceSkusListSkusOperationResponse) LoadMore(ctx context.Context) (resp ResourceSkusListSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ResourceSkusListSkus ...
func (c StandardOperationClient) ResourceSkusListSkus(ctx context.Context, id commonids.SubscriptionId) (resp ResourceSkusListSkusOperationResponse, err error) {
	req, err := c.preparerForResourceSkusListSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForResourceSkusListSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForResourceSkusListSkus prepares the ResourceSkusListSkus request.
func (c StandardOperationClient) preparerForResourceSkusListSkus(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.DataMigration/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForResourceSkusListSkusWithNextLink prepares the ResourceSkusListSkus request with the given nextLink token.
func (c StandardOperationClient) preparerForResourceSkusListSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForResourceSkusListSkus handles the response to the ResourceSkusListSkus request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForResourceSkusListSkus(resp *http.Response) (result ResourceSkusListSkusOperationResponse, err error) {
	type page struct {
		Values   []ResourceSku `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ResourceSkusListSkusOperationResponse, err error) {
			req, err := c.preparerForResourceSkusListSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForResourceSkusListSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ResourceSkusListSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ResourceSkusListSkusComplete retrieves all of the results into a single object
func (c StandardOperationClient) ResourceSkusListSkusComplete(ctx context.Context, id commonids.SubscriptionId) (ResourceSkusListSkusCompleteResult, error) {
	return c.ResourceSkusListSkusCompleteMatchingPredicate(ctx, id, ResourceSkuOperationPredicate{})
}

// ResourceSkusListSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StandardOperationClient) ResourceSkusListSkusCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ResourceSkuOperationPredicate) (resp ResourceSkusListSkusCompleteResult, err error) {
	items := make([]ResourceSku, 0)

	page, err := c.ResourceSkusListSkus(ctx, id)
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

	out := ResourceSkusListSkusCompleteResult{
		Items: items,
	}
	return out, nil
}
