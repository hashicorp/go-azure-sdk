package vmhost

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

type SubAccountListVMHostUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMResources

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SubAccountListVMHostUpdateOperationResponse, error)
}

type SubAccountListVMHostUpdateCompleteResult struct {
	Items []VMResources
}

func (r SubAccountListVMHostUpdateOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SubAccountListVMHostUpdateOperationResponse) LoadMore(ctx context.Context) (resp SubAccountListVMHostUpdateOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SubAccountListVMHostUpdate ...
func (c VMHostClient) SubAccountListVMHostUpdate(ctx context.Context, id AccountId, input VMHostUpdateRequest) (resp SubAccountListVMHostUpdateOperationResponse, err error) {
	req, err := c.preparerForSubAccountListVMHostUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSubAccountListVMHostUpdate(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSubAccountListVMHostUpdate prepares the SubAccountListVMHostUpdate request.
func (c VMHostClient) preparerForSubAccountListVMHostUpdate(ctx context.Context, id AccountId, input VMHostUpdateRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/vmHostUpdate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSubAccountListVMHostUpdateWithNextLink prepares the SubAccountListVMHostUpdate request with the given nextLink token.
func (c VMHostClient) preparerForSubAccountListVMHostUpdateWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSubAccountListVMHostUpdate handles the response to the SubAccountListVMHostUpdate request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForSubAccountListVMHostUpdate(resp *http.Response) (result SubAccountListVMHostUpdateOperationResponse, err error) {
	type page struct {
		Values   []VMResources `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SubAccountListVMHostUpdateOperationResponse, err error) {
			req, err := c.preparerForSubAccountListVMHostUpdateWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSubAccountListVMHostUpdate(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "SubAccountListVMHostUpdate", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SubAccountListVMHostUpdateComplete retrieves all of the results into a single object
func (c VMHostClient) SubAccountListVMHostUpdateComplete(ctx context.Context, id AccountId, input VMHostUpdateRequest) (SubAccountListVMHostUpdateCompleteResult, error) {
	return c.SubAccountListVMHostUpdateCompleteMatchingPredicate(ctx, id, input, VMResourcesOperationPredicate{})
}

// SubAccountListVMHostUpdateCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VMHostClient) SubAccountListVMHostUpdateCompleteMatchingPredicate(ctx context.Context, id AccountId, input VMHostUpdateRequest, predicate VMResourcesOperationPredicate) (resp SubAccountListVMHostUpdateCompleteResult, err error) {
	items := make([]VMResources, 0)

	page, err := c.SubAccountListVMHostUpdate(ctx, id, input)
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

	out := SubAccountListVMHostUpdateCompleteResult{
		Items: items,
	}
	return out, nil
}
