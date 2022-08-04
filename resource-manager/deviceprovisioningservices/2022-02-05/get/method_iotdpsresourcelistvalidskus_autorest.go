package get

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

type IotDpsResourcelistValidSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IotDpsSkuDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IotDpsResourcelistValidSkusOperationResponse, error)
}

type IotDpsResourcelistValidSkusCompleteResult struct {
	Items []IotDpsSkuDefinition
}

func (r IotDpsResourcelistValidSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IotDpsResourcelistValidSkusOperationResponse) LoadMore(ctx context.Context) (resp IotDpsResourcelistValidSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// IotDpsResourcelistValidSkus ...
func (c GETClient) IotDpsResourcelistValidSkus(ctx context.Context, id ProvisioningServiceId) (resp IotDpsResourcelistValidSkusOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourcelistValidSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIotDpsResourcelistValidSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// IotDpsResourcelistValidSkusComplete retrieves all of the results into a single object
func (c GETClient) IotDpsResourcelistValidSkusComplete(ctx context.Context, id ProvisioningServiceId) (IotDpsResourcelistValidSkusCompleteResult, error) {
	return c.IotDpsResourcelistValidSkusCompleteMatchingPredicate(ctx, id, IotDpsSkuDefinitionOperationPredicate{})
}

// IotDpsResourcelistValidSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GETClient) IotDpsResourcelistValidSkusCompleteMatchingPredicate(ctx context.Context, id ProvisioningServiceId, predicate IotDpsSkuDefinitionOperationPredicate) (resp IotDpsResourcelistValidSkusCompleteResult, err error) {
	items := make([]IotDpsSkuDefinition, 0)

	page, err := c.IotDpsResourcelistValidSkus(ctx, id)
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

	out := IotDpsResourcelistValidSkusCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForIotDpsResourcelistValidSkus prepares the IotDpsResourcelistValidSkus request.
func (c GETClient) preparerForIotDpsResourcelistValidSkus(ctx context.Context, id ProvisioningServiceId) (*http.Request, error) {
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

// preparerForIotDpsResourcelistValidSkusWithNextLink prepares the IotDpsResourcelistValidSkus request with the given nextLink token.
func (c GETClient) preparerForIotDpsResourcelistValidSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIotDpsResourcelistValidSkus handles the response to the IotDpsResourcelistValidSkus request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourcelistValidSkus(resp *http.Response) (result IotDpsResourcelistValidSkusOperationResponse, err error) {
	type page struct {
		Values   []IotDpsSkuDefinition `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IotDpsResourcelistValidSkusOperationResponse, err error) {
			req, err := c.preparerForIotDpsResourcelistValidSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIotDpsResourcelistValidSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourcelistValidSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
