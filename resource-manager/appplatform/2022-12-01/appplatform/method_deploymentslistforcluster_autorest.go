package appplatform

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

type DeploymentsListForClusterOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (DeploymentsListForClusterOperationResponse, error)
}

type DeploymentsListForClusterCompleteResult struct {
	Items []DeploymentResource
}

func (r DeploymentsListForClusterOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r DeploymentsListForClusterOperationResponse) LoadMore(ctx context.Context) (resp DeploymentsListForClusterOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type DeploymentsListForClusterOperationOptions struct {
	Version *[]string
}

func DefaultDeploymentsListForClusterOperationOptions() DeploymentsListForClusterOperationOptions {
	return DeploymentsListForClusterOperationOptions{}
}

func (o DeploymentsListForClusterOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeploymentsListForClusterOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Version != nil {
		out["version"] = *o.Version
	}

	return out
}

// DeploymentsListForCluster ...
func (c AppPlatformClient) DeploymentsListForCluster(ctx context.Context, id SpringId, options DeploymentsListForClusterOperationOptions) (resp DeploymentsListForClusterOperationResponse, err error) {
	req, err := c.preparerForDeploymentsListForCluster(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForDeploymentsListForCluster(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForDeploymentsListForCluster prepares the DeploymentsListForCluster request.
func (c AppPlatformClient) preparerForDeploymentsListForCluster(ctx context.Context, id SpringId, options DeploymentsListForClusterOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/deployments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForDeploymentsListForClusterWithNextLink prepares the DeploymentsListForCluster request with the given nextLink token.
func (c AppPlatformClient) preparerForDeploymentsListForClusterWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForDeploymentsListForCluster handles the response to the DeploymentsListForCluster request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForDeploymentsListForCluster(resp *http.Response) (result DeploymentsListForClusterOperationResponse, err error) {
	type page struct {
		Values   []DeploymentResource `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result DeploymentsListForClusterOperationResponse, err error) {
			req, err := c.preparerForDeploymentsListForClusterWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForDeploymentsListForCluster(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsListForCluster", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// DeploymentsListForClusterComplete retrieves all of the results into a single object
func (c AppPlatformClient) DeploymentsListForClusterComplete(ctx context.Context, id SpringId, options DeploymentsListForClusterOperationOptions) (DeploymentsListForClusterCompleteResult, error) {
	return c.DeploymentsListForClusterCompleteMatchingPredicate(ctx, id, options, DeploymentResourceOperationPredicate{})
}

// DeploymentsListForClusterCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) DeploymentsListForClusterCompleteMatchingPredicate(ctx context.Context, id SpringId, options DeploymentsListForClusterOperationOptions, predicate DeploymentResourceOperationPredicate) (resp DeploymentsListForClusterCompleteResult, err error) {
	items := make([]DeploymentResource, 0)

	page, err := c.DeploymentsListForCluster(ctx, id, options)
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

	out := DeploymentsListForClusterCompleteResult{
		Items: items,
	}
	return out, nil
}
