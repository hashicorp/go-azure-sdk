package appserviceenvironments

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

type GetOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]OutboundEnvironmentEndpoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetOutboundNetworkDependenciesEndpointsOperationResponse, error)
}

type GetOutboundNetworkDependenciesEndpointsCompleteResult struct {
	Items []OutboundEnvironmentEndpoint
}

func (r GetOutboundNetworkDependenciesEndpointsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetOutboundNetworkDependenciesEndpointsOperationResponse) LoadMore(ctx context.Context) (resp GetOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetOutboundNetworkDependenciesEndpoints ...
func (c AppServiceEnvironmentsClient) GetOutboundNetworkDependenciesEndpoints(ctx context.Context, id HostingEnvironmentId) (resp GetOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	req, err := c.preparerForGetOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetOutboundNetworkDependenciesEndpoints(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetOutboundNetworkDependenciesEndpoints prepares the GetOutboundNetworkDependenciesEndpoints request.
func (c AppServiceEnvironmentsClient) preparerForGetOutboundNetworkDependenciesEndpoints(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/outboundNetworkDependenciesEndpoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetOutboundNetworkDependenciesEndpointsWithNextLink prepares the GetOutboundNetworkDependenciesEndpoints request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForGetOutboundNetworkDependenciesEndpointsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetOutboundNetworkDependenciesEndpoints handles the response to the GetOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetOutboundNetworkDependenciesEndpoints(resp *http.Response) (result GetOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	type page struct {
		Values   []OutboundEnvironmentEndpoint `json:"value"`
		NextLink *string                       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
			req, err := c.preparerForGetOutboundNetworkDependenciesEndpointsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetOutboundNetworkDependenciesEndpoints(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetOutboundNetworkDependenciesEndpointsComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) GetOutboundNetworkDependenciesEndpointsComplete(ctx context.Context, id HostingEnvironmentId) (GetOutboundNetworkDependenciesEndpointsCompleteResult, error) {
	return c.GetOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx, id, OutboundEnvironmentEndpointOperationPredicate{})
}

// GetOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) GetOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate OutboundEnvironmentEndpointOperationPredicate) (resp GetOutboundNetworkDependenciesEndpointsCompleteResult, err error) {
	items := make([]OutboundEnvironmentEndpoint, 0)

	page, err := c.GetOutboundNetworkDependenciesEndpoints(ctx, id)
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

	out := GetOutboundNetworkDependenciesEndpointsCompleteResult{
		Items: items,
	}
	return out, nil
}
