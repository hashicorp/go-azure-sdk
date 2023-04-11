package iotsecuritysolutions

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

type IoTSecuritySolutionsResourceGroupListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IoTSecuritySolutionModel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IoTSecuritySolutionsResourceGroupListOperationResponse, error)
}

type IoTSecuritySolutionsResourceGroupListCompleteResult struct {
	Items []IoTSecuritySolutionModel
}

func (r IoTSecuritySolutionsResourceGroupListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IoTSecuritySolutionsResourceGroupListOperationResponse) LoadMore(ctx context.Context) (resp IoTSecuritySolutionsResourceGroupListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type IoTSecuritySolutionsResourceGroupListOperationOptions struct {
	Filter *string
}

func DefaultIoTSecuritySolutionsResourceGroupListOperationOptions() IoTSecuritySolutionsResourceGroupListOperationOptions {
	return IoTSecuritySolutionsResourceGroupListOperationOptions{}
}

func (o IoTSecuritySolutionsResourceGroupListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o IoTSecuritySolutionsResourceGroupListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// IoTSecuritySolutionsResourceGroupList ...
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupList(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions) (resp IoTSecuritySolutionsResourceGroupListOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsResourceGroupList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIoTSecuritySolutionsResourceGroupList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForIoTSecuritySolutionsResourceGroupList prepares the IoTSecuritySolutionsResourceGroupList request.
func (c IotSecuritySolutionsClient) preparerForIoTSecuritySolutionsResourceGroupList(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/iotSecuritySolutions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIoTSecuritySolutionsResourceGroupListWithNextLink prepares the IoTSecuritySolutionsResourceGroupList request with the given nextLink token.
func (c IotSecuritySolutionsClient) preparerForIoTSecuritySolutionsResourceGroupListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsResourceGroupList handles the response to the IoTSecuritySolutionsResourceGroupList request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIoTSecuritySolutionsResourceGroupList(resp *http.Response) (result IoTSecuritySolutionsResourceGroupListOperationResponse, err error) {
	type page struct {
		Values   []IoTSecuritySolutionModel `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IoTSecuritySolutionsResourceGroupListOperationResponse, err error) {
			req, err := c.preparerForIoTSecuritySolutionsResourceGroupListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIoTSecuritySolutionsResourceGroupList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsResourceGroupList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// IoTSecuritySolutionsResourceGroupListComplete retrieves all of the results into a single object
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupListComplete(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions) (IoTSecuritySolutionsResourceGroupListCompleteResult, error) {
	return c.IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate(ctx, id, options, IoTSecuritySolutionModelOperationPredicate{})
}

// IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsResourceGroupListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options IoTSecuritySolutionsResourceGroupListOperationOptions, predicate IoTSecuritySolutionModelOperationPredicate) (resp IoTSecuritySolutionsResourceGroupListCompleteResult, err error) {
	items := make([]IoTSecuritySolutionModel, 0)

	page, err := c.IoTSecuritySolutionsResourceGroupList(ctx, id, options)
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

	out := IoTSecuritySolutionsResourceGroupListCompleteResult{
		Items: items,
	}
	return out, nil
}
