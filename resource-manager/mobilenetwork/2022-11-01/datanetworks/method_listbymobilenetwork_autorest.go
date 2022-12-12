package datanetworks

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

type ListByMobileNetworkOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DataNetwork

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByMobileNetworkOperationResponse, error)
}

type ListByMobileNetworkCompleteResult struct {
	Items []DataNetwork
}

func (r ListByMobileNetworkOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByMobileNetworkOperationResponse) LoadMore(ctx context.Context) (resp ListByMobileNetworkOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByMobileNetwork ...
func (c DataNetworksClient) ListByMobileNetwork(ctx context.Context, id MobileNetworkId) (resp ListByMobileNetworkOperationResponse, err error) {
	req, err := c.preparerForListByMobileNetwork(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByMobileNetwork(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByMobileNetwork prepares the ListByMobileNetwork request.
func (c DataNetworksClient) preparerForListByMobileNetwork(ctx context.Context, id MobileNetworkId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dataNetworks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByMobileNetworkWithNextLink prepares the ListByMobileNetwork request with the given nextLink token.
func (c DataNetworksClient) preparerForListByMobileNetworkWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByMobileNetwork handles the response to the ListByMobileNetwork request. The method always
// closes the http.Response Body.
func (c DataNetworksClient) responderForListByMobileNetwork(resp *http.Response) (result ListByMobileNetworkOperationResponse, err error) {
	type page struct {
		Values   []DataNetwork `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByMobileNetworkOperationResponse, err error) {
			req, err := c.preparerForListByMobileNetworkWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByMobileNetwork(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datanetworks.DataNetworksClient", "ListByMobileNetwork", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByMobileNetworkComplete retrieves all of the results into a single object
func (c DataNetworksClient) ListByMobileNetworkComplete(ctx context.Context, id MobileNetworkId) (ListByMobileNetworkCompleteResult, error) {
	return c.ListByMobileNetworkCompleteMatchingPredicate(ctx, id, DataNetworkOperationPredicate{})
}

// ListByMobileNetworkCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DataNetworksClient) ListByMobileNetworkCompleteMatchingPredicate(ctx context.Context, id MobileNetworkId, predicate DataNetworkOperationPredicate) (resp ListByMobileNetworkCompleteResult, err error) {
	items := make([]DataNetwork, 0)

	page, err := c.ListByMobileNetwork(ctx, id)
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

	out := ListByMobileNetworkCompleteResult{
		Items: items,
	}
	return out, nil
}
