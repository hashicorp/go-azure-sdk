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

type IoTSecuritySolutionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IoTSecuritySolutionModel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IoTSecuritySolutionsListOperationResponse, error)
}

type IoTSecuritySolutionsListCompleteResult struct {
	Items []IoTSecuritySolutionModel
}

func (r IoTSecuritySolutionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IoTSecuritySolutionsListOperationResponse) LoadMore(ctx context.Context) (resp IoTSecuritySolutionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type IoTSecuritySolutionsListOperationOptions struct {
	Filter *string
}

func DefaultIoTSecuritySolutionsListOperationOptions() IoTSecuritySolutionsListOperationOptions {
	return IoTSecuritySolutionsListOperationOptions{}
}

func (o IoTSecuritySolutionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o IoTSecuritySolutionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// IoTSecuritySolutionsList ...
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsList(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions) (resp IoTSecuritySolutionsListOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIoTSecuritySolutionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForIoTSecuritySolutionsList prepares the IoTSecuritySolutionsList request.
func (c IotSecuritySolutionsClient) preparerForIoTSecuritySolutionsList(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions) (*http.Request, error) {
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

// preparerForIoTSecuritySolutionsListWithNextLink prepares the IoTSecuritySolutionsList request with the given nextLink token.
func (c IotSecuritySolutionsClient) preparerForIoTSecuritySolutionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsList handles the response to the IoTSecuritySolutionsList request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIoTSecuritySolutionsList(resp *http.Response) (result IoTSecuritySolutionsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IoTSecuritySolutionsListOperationResponse, err error) {
			req, err := c.preparerForIoTSecuritySolutionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIoTSecuritySolutionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IoTSecuritySolutionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// IoTSecuritySolutionsListComplete retrieves all of the results into a single object
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsListComplete(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions) (IoTSecuritySolutionsListCompleteResult, error) {
	return c.IoTSecuritySolutionsListCompleteMatchingPredicate(ctx, id, options, IoTSecuritySolutionModelOperationPredicate{})
}

// IoTSecuritySolutionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IotSecuritySolutionsClient) IoTSecuritySolutionsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options IoTSecuritySolutionsListOperationOptions, predicate IoTSecuritySolutionModelOperationPredicate) (resp IoTSecuritySolutionsListCompleteResult, err error) {
	items := make([]IoTSecuritySolutionModel, 0)

	page, err := c.IoTSecuritySolutionsList(ctx, id, options)
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

	out := IoTSecuritySolutionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
