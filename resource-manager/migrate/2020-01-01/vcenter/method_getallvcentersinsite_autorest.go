package vcenter

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

type GetAllVCentersInSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VCenter

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAllVCentersInSiteOperationResponse, error)
}

type GetAllVCentersInSiteCompleteResult struct {
	Items []VCenter
}

func (r GetAllVCentersInSiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAllVCentersInSiteOperationResponse) LoadMore(ctx context.Context) (resp GetAllVCentersInSiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetAllVCentersInSiteOperationOptions struct {
	Filter *string
}

func DefaultGetAllVCentersInSiteOperationOptions() GetAllVCentersInSiteOperationOptions {
	return GetAllVCentersInSiteOperationOptions{}
}

func (o GetAllVCentersInSiteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAllVCentersInSiteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// GetAllVCentersInSite ...
func (c VCenterClient) GetAllVCentersInSite(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions) (resp GetAllVCentersInSiteOperationResponse, err error) {
	req, err := c.preparerForGetAllVCentersInSite(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAllVCentersInSite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAllVCentersInSite prepares the GetAllVCentersInSite request.
func (c VCenterClient) preparerForGetAllVCentersInSite(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/vCenters", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAllVCentersInSiteWithNextLink prepares the GetAllVCentersInSite request with the given nextLink token.
func (c VCenterClient) preparerForGetAllVCentersInSiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAllVCentersInSite handles the response to the GetAllVCentersInSite request. The method always
// closes the http.Response Body.
func (c VCenterClient) responderForGetAllVCentersInSite(resp *http.Response) (result GetAllVCentersInSiteOperationResponse, err error) {
	type page struct {
		Values   []VCenter `json:"value"`
		NextLink *string   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAllVCentersInSiteOperationResponse, err error) {
			req, err := c.preparerForGetAllVCentersInSiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAllVCentersInSite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetAllVCentersInSite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAllVCentersInSiteComplete retrieves all of the results into a single object
func (c VCenterClient) GetAllVCentersInSiteComplete(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions) (GetAllVCentersInSiteCompleteResult, error) {
	return c.GetAllVCentersInSiteCompleteMatchingPredicate(ctx, id, options, VCenterOperationPredicate{})
}

// GetAllVCentersInSiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VCenterClient) GetAllVCentersInSiteCompleteMatchingPredicate(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions, predicate VCenterOperationPredicate) (resp GetAllVCentersInSiteCompleteResult, err error) {
	items := make([]VCenter, 0)

	page, err := c.GetAllVCentersInSite(ctx, id, options)
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

	out := GetAllVCentersInSiteCompleteResult{
		Items: items,
	}
	return out, nil
}
