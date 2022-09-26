package virtualmachinescalesets

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

type ListSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineScaleSetSku

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSkusOperationResponse, error)
}

type ListSkusCompleteResult struct {
	Items []VirtualMachineScaleSetSku
}

func (r ListSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSkusOperationResponse) LoadMore(ctx context.Context) (resp ListSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSkus ...
func (c VirtualMachineScaleSetsClient) ListSkus(ctx context.Context, id VirtualMachineScaleSetId) (resp ListSkusOperationResponse, err error) {
	req, err := c.preparerForListSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSkus prepares the ListSkus request.
func (c VirtualMachineScaleSetsClient) preparerForListSkus(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSkusWithNextLink prepares the ListSkus request with the given nextLink token.
func (c VirtualMachineScaleSetsClient) preparerForListSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSkus handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (c VirtualMachineScaleSetsClient) responderForListSkus(resp *http.Response) (result ListSkusOperationResponse, err error) {
	type page struct {
		Values   []VirtualMachineScaleSetSku `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSkusOperationResponse, err error) {
			req, err := c.preparerForListSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "ListSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSkusComplete retrieves all of the results into a single object
func (c VirtualMachineScaleSetsClient) ListSkusComplete(ctx context.Context, id VirtualMachineScaleSetId) (ListSkusCompleteResult, error) {
	return c.ListSkusCompleteMatchingPredicate(ctx, id, VirtualMachineScaleSetSkuOperationPredicate{})
}

// ListSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VirtualMachineScaleSetsClient) ListSkusCompleteMatchingPredicate(ctx context.Context, id VirtualMachineScaleSetId, predicate VirtualMachineScaleSetSkuOperationPredicate) (resp ListSkusCompleteResult, err error) {
	items := make([]VirtualMachineScaleSetSku, 0)

	page, err := c.ListSkus(ctx, id)
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

	out := ListSkusCompleteResult{
		Items: items,
	}
	return out, nil
}
