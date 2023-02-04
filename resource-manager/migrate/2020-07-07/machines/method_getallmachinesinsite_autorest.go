package machines

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

type GetAllMachinesInSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMwareMachine

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAllMachinesInSiteOperationResponse, error)
}

type GetAllMachinesInSiteCompleteResult struct {
	Items []VMwareMachine
}

func (r GetAllMachinesInSiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAllMachinesInSiteOperationResponse) LoadMore(ctx context.Context) (resp GetAllMachinesInSiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetAllMachinesInSiteOperationOptions struct {
	ContinuationToken *string
	Filter            *string
	Top               *int64
	TotalRecordCount  *int64
}

func DefaultGetAllMachinesInSiteOperationOptions() GetAllMachinesInSiteOperationOptions {
	return GetAllMachinesInSiteOperationOptions{}
}

func (o GetAllMachinesInSiteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAllMachinesInSiteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ContinuationToken != nil {
		out["continuationToken"] = *o.ContinuationToken
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	if o.TotalRecordCount != nil {
		out["totalRecordCount"] = *o.TotalRecordCount
	}

	return out
}

// GetAllMachinesInSite ...
func (c MachinesClient) GetAllMachinesInSite(ctx context.Context, id VMwareSiteId, options GetAllMachinesInSiteOperationOptions) (resp GetAllMachinesInSiteOperationResponse, err error) {
	req, err := c.preparerForGetAllMachinesInSite(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAllMachinesInSite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAllMachinesInSite prepares the GetAllMachinesInSite request.
func (c MachinesClient) preparerForGetAllMachinesInSite(ctx context.Context, id VMwareSiteId, options GetAllMachinesInSiteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/machines", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAllMachinesInSiteWithNextLink prepares the GetAllMachinesInSite request with the given nextLink token.
func (c MachinesClient) preparerForGetAllMachinesInSiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAllMachinesInSite handles the response to the GetAllMachinesInSite request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForGetAllMachinesInSite(resp *http.Response) (result GetAllMachinesInSiteOperationResponse, err error) {
	type page struct {
		Values   []VMwareMachine `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAllMachinesInSiteOperationResponse, err error) {
			req, err := c.preparerForGetAllMachinesInSiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAllMachinesInSite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetAllMachinesInSite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAllMachinesInSiteComplete retrieves all of the results into a single object
func (c MachinesClient) GetAllMachinesInSiteComplete(ctx context.Context, id VMwareSiteId, options GetAllMachinesInSiteOperationOptions) (GetAllMachinesInSiteCompleteResult, error) {
	return c.GetAllMachinesInSiteCompleteMatchingPredicate(ctx, id, options, VMwareMachineOperationPredicate{})
}

// GetAllMachinesInSiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MachinesClient) GetAllMachinesInSiteCompleteMatchingPredicate(ctx context.Context, id VMwareSiteId, options GetAllMachinesInSiteOperationOptions, predicate VMwareMachineOperationPredicate) (resp GetAllMachinesInSiteCompleteResult, err error) {
	items := make([]VMwareMachine, 0)

	page, err := c.GetAllMachinesInSite(ctx, id, options)
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

	out := GetAllMachinesInSiteCompleteResult{
		Items: items,
	}
	return out, nil
}
