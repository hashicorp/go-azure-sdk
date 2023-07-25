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

type ListMultiRoleMetricDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResourceMetricDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListMultiRoleMetricDefinitionsOperationResponse, error)
}

type ListMultiRoleMetricDefinitionsCompleteResult struct {
	Items []ResourceMetricDefinition
}

func (r ListMultiRoleMetricDefinitionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListMultiRoleMetricDefinitionsOperationResponse) LoadMore(ctx context.Context) (resp ListMultiRoleMetricDefinitionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListMultiRoleMetricDefinitions ...
func (c AppServiceEnvironmentsClient) ListMultiRoleMetricDefinitions(ctx context.Context, id HostingEnvironmentId) (resp ListMultiRoleMetricDefinitionsOperationResponse, err error) {
	req, err := c.preparerForListMultiRoleMetricDefinitions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListMultiRoleMetricDefinitions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListMultiRoleMetricDefinitions prepares the ListMultiRoleMetricDefinitions request.
func (c AppServiceEnvironmentsClient) preparerForListMultiRoleMetricDefinitions(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default/metricdefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListMultiRoleMetricDefinitionsWithNextLink prepares the ListMultiRoleMetricDefinitions request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListMultiRoleMetricDefinitionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListMultiRoleMetricDefinitions handles the response to the ListMultiRoleMetricDefinitions request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListMultiRoleMetricDefinitions(resp *http.Response) (result ListMultiRoleMetricDefinitionsOperationResponse, err error) {
	type page struct {
		Values   []ResourceMetricDefinition `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListMultiRoleMetricDefinitionsOperationResponse, err error) {
			req, err := c.preparerForListMultiRoleMetricDefinitionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListMultiRoleMetricDefinitions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRoleMetricDefinitions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListMultiRoleMetricDefinitionsComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRoleMetricDefinitionsComplete(ctx context.Context, id HostingEnvironmentId) (ListMultiRoleMetricDefinitionsCompleteResult, error) {
	return c.ListMultiRoleMetricDefinitionsCompleteMatchingPredicate(ctx, id, ResourceMetricDefinitionOperationPredicate{})
}

// ListMultiRoleMetricDefinitionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListMultiRoleMetricDefinitionsCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate ResourceMetricDefinitionOperationPredicate) (resp ListMultiRoleMetricDefinitionsCompleteResult, err error) {
	items := make([]ResourceMetricDefinition, 0)

	page, err := c.ListMultiRoleMetricDefinitions(ctx, id)
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

	out := ListMultiRoleMetricDefinitionsCompleteResult{
		Items: items,
	}
	return out, nil
}
