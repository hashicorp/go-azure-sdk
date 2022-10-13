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

type DeploymentsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (DeploymentsListOperationResponse, error)
}

type DeploymentsListCompleteResult struct {
	Items []DeploymentResource
}

func (r DeploymentsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r DeploymentsListOperationResponse) LoadMore(ctx context.Context) (resp DeploymentsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type DeploymentsListOperationOptions struct {
	Version *[]string
}

func DefaultDeploymentsListOperationOptions() DeploymentsListOperationOptions {
	return DeploymentsListOperationOptions{}
}

func (o DeploymentsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeploymentsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Version != nil {
		out["version"] = *o.Version
	}

	return out
}

// DeploymentsList ...
func (c AppPlatformClient) DeploymentsList(ctx context.Context, id AppId, options DeploymentsListOperationOptions) (resp DeploymentsListOperationResponse, err error) {
	req, err := c.preparerForDeploymentsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForDeploymentsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForDeploymentsList prepares the DeploymentsList request.
func (c AppPlatformClient) preparerForDeploymentsList(ctx context.Context, id AppId, options DeploymentsListOperationOptions) (*http.Request, error) {
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

// preparerForDeploymentsListWithNextLink prepares the DeploymentsList request with the given nextLink token.
func (c AppPlatformClient) preparerForDeploymentsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForDeploymentsList handles the response to the DeploymentsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForDeploymentsList(resp *http.Response) (result DeploymentsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result DeploymentsListOperationResponse, err error) {
			req, err := c.preparerForDeploymentsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForDeploymentsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// DeploymentsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) DeploymentsListComplete(ctx context.Context, id AppId, options DeploymentsListOperationOptions) (DeploymentsListCompleteResult, error) {
	return c.DeploymentsListCompleteMatchingPredicate(ctx, id, options, DeploymentResourceOperationPredicate{})
}

// DeploymentsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) DeploymentsListCompleteMatchingPredicate(ctx context.Context, id AppId, options DeploymentsListOperationOptions, predicate DeploymentResourceOperationPredicate) (resp DeploymentsListCompleteResult, err error) {
	items := make([]DeploymentResource, 0)

	page, err := c.DeploymentsList(ctx, id, options)
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

	out := DeploymentsListCompleteResult{
		Items: items,
	}
	return out, nil
}
