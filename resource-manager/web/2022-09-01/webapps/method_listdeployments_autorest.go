package webapps

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

type ListDeploymentsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Deployment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDeploymentsOperationResponse, error)
}

type ListDeploymentsCompleteResult struct {
	Items []Deployment
}

func (r ListDeploymentsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDeploymentsOperationResponse) LoadMore(ctx context.Context) (resp ListDeploymentsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDeployments ...
func (c WebAppsClient) ListDeployments(ctx context.Context, id SiteId) (resp ListDeploymentsOperationResponse, err error) {
	req, err := c.preparerForListDeployments(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDeployments(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDeployments prepares the ListDeployments request.
func (c WebAppsClient) preparerForListDeployments(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deployments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDeploymentsWithNextLink prepares the ListDeployments request with the given nextLink token.
func (c WebAppsClient) preparerForListDeploymentsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListDeployments handles the response to the ListDeployments request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListDeployments(resp *http.Response) (result ListDeploymentsOperationResponse, err error) {
	type page struct {
		Values   []Deployment `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDeploymentsOperationResponse, err error) {
			req, err := c.preparerForListDeploymentsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDeployments(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeployments", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDeploymentsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListDeploymentsComplete(ctx context.Context, id SiteId) (ListDeploymentsCompleteResult, error) {
	return c.ListDeploymentsCompleteMatchingPredicate(ctx, id, DeploymentOperationPredicate{})
}

// ListDeploymentsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListDeploymentsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate DeploymentOperationPredicate) (resp ListDeploymentsCompleteResult, err error) {
	items := make([]Deployment, 0)

	page, err := c.ListDeployments(ctx, id)
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

	out := ListDeploymentsCompleteResult{
		Items: items,
	}
	return out, nil
}
