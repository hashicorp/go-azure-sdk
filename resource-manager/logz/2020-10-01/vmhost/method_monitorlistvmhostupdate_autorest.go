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

type MonitorListVMHostUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMResources

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (MonitorListVMHostUpdateOperationResponse, error)
}

type MonitorListVMHostUpdateCompleteResult struct {
	Items []VMResources
}

func (r MonitorListVMHostUpdateOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r MonitorListVMHostUpdateOperationResponse) LoadMore(ctx context.Context) (resp MonitorListVMHostUpdateOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// MonitorListVMHostUpdate ...
func (c VMHostClient) MonitorListVMHostUpdate(ctx context.Context, id MonitorId, input VMHostUpdateRequest) (resp MonitorListVMHostUpdateOperationResponse, err error) {
	req, err := c.preparerForMonitorListVMHostUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForMonitorListVMHostUpdate(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForMonitorListVMHostUpdate prepares the MonitorListVMHostUpdate request.
func (c VMHostClient) preparerForMonitorListVMHostUpdate(ctx context.Context, id MonitorId, input VMHostUpdateRequest) (*http.Request, error) {
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

// preparerForMonitorListVMHostUpdateWithNextLink prepares the MonitorListVMHostUpdate request with the given nextLink token.
func (c VMHostClient) preparerForMonitorListVMHostUpdateWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForMonitorListVMHostUpdate handles the response to the MonitorListVMHostUpdate request. The method always
// closes the http.Response Body.
func (c VMHostClient) responderForMonitorListVMHostUpdate(resp *http.Response) (result MonitorListVMHostUpdateOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result MonitorListVMHostUpdateOperationResponse, err error) {
			req, err := c.preparerForMonitorListVMHostUpdateWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForMonitorListVMHostUpdate(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmhost.VMHostClient", "MonitorListVMHostUpdate", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// MonitorListVMHostUpdateComplete retrieves all of the results into a single object
func (c VMHostClient) MonitorListVMHostUpdateComplete(ctx context.Context, id MonitorId, input VMHostUpdateRequest) (MonitorListVMHostUpdateCompleteResult, error) {
	return c.MonitorListVMHostUpdateCompleteMatchingPredicate(ctx, id, input, VMResourcesOperationPredicate{})
}

// MonitorListVMHostUpdateCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VMHostClient) MonitorListVMHostUpdateCompleteMatchingPredicate(ctx context.Context, id MonitorId, input VMHostUpdateRequest, predicate VMResourcesOperationPredicate) (resp MonitorListVMHostUpdateCompleteResult, err error) {
	items := make([]VMResources, 0)

	page, err := c.MonitorListVMHostUpdate(ctx, id, input)
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

	out := MonitorListVMHostUpdateCompleteResult{
		Items: items,
	}
	return out, nil
}
