package managedvmsizes

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

type ManagedUnsupportedVMSizesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ManagedVMSize

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ManagedUnsupportedVMSizesListOperationResponse, error)
}

type ManagedUnsupportedVMSizesListCompleteResult struct {
	Items []ManagedVMSize
}

func (r ManagedUnsupportedVMSizesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ManagedUnsupportedVMSizesListOperationResponse) LoadMore(ctx context.Context) (resp ManagedUnsupportedVMSizesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ManagedUnsupportedVMSizesList ...
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesList(ctx context.Context, id LocationId) (resp ManagedUnsupportedVMSizesListOperationResponse, err error) {
	req, err := c.preparerForManagedUnsupportedVMSizesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForManagedUnsupportedVMSizesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForManagedUnsupportedVMSizesList prepares the ManagedUnsupportedVMSizesList request.
func (c ManagedVMSizesClient) preparerForManagedUnsupportedVMSizesList(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/managedUnsupportedVMSizes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForManagedUnsupportedVMSizesListWithNextLink prepares the ManagedUnsupportedVMSizesList request with the given nextLink token.
func (c ManagedVMSizesClient) preparerForManagedUnsupportedVMSizesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForManagedUnsupportedVMSizesList handles the response to the ManagedUnsupportedVMSizesList request. The method always
// closes the http.Response Body.
func (c ManagedVMSizesClient) responderForManagedUnsupportedVMSizesList(resp *http.Response) (result ManagedUnsupportedVMSizesListOperationResponse, err error) {
	type page struct {
		Values   []ManagedVMSize `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ManagedUnsupportedVMSizesListOperationResponse, err error) {
			req, err := c.preparerForManagedUnsupportedVMSizesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForManagedUnsupportedVMSizesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ManagedUnsupportedVMSizesListComplete retrieves all of the results into a single object
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesListComplete(ctx context.Context, id LocationId) (ManagedUnsupportedVMSizesListCompleteResult, error) {
	return c.ManagedUnsupportedVMSizesListCompleteMatchingPredicate(ctx, id, ManagedVMSizeOperationPredicate{})
}

// ManagedUnsupportedVMSizesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate ManagedVMSizeOperationPredicate) (resp ManagedUnsupportedVMSizesListCompleteResult, err error) {
	items := make([]ManagedVMSize, 0)

	page, err := c.ManagedUnsupportedVMSizesList(ctx, id)
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

	out := ManagedUnsupportedVMSizesListCompleteResult{
		Items: items,
	}
	return out, nil
}
