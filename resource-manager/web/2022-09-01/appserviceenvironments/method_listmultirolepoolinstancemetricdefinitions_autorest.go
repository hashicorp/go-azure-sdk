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

type ListMultiRolePoolInstanceMetricDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResourceMetricDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListMultiRolePoolInstanceMetricDefinitionsOperationResponse, error)
}

type ListMultiRolePoolInstanceMetricDefinitionsCompleteResult struct {
	Items []ResourceMetricDefinition
}

func (r ListMultiRolePoolInstanceMetricDefinitionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListMultiRolePoolInstanceMetricDefinitionsOperationResponse) LoadMore(ctx context.Context) (resp ListMultiRolePoolInstanceMetricDefinitionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListMultiRolePoolInstanceMetricDefinitions ...
func (c AppServiceEnvironmentsClient) ListMultiRolePoolInstanceMetricDefinitions(ctx context.Context, id DefaultInstanceId) (resp ListMultiRolePoolInstanceMetricDefinitionsOperationResponse, err error) {
	req, err := c.preparerForListMultiRolePoolInstanceMetricDefinitions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListMultiRolePoolInstanceMetricDefinitions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListMultiRolePoolInstanceMetricDefinitions prepares the ListMultiRolePoolInstanceMetricDefinitions request.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePoolInstanceMetricDefinitions(ctx context.Context, id DefaultInstanceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/metricdefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListMultiRolePoolInstanceMetricDefinitionsWithNextLink prepares the ListMultiRolePoolInstanceMetricDefinitions request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePoolInstanceMetricDefinitionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListMultiRolePoolInstanceMetricDefinitions handles the response to the ListMultiRolePoolInstanceMetricDefinitions request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListMultiRolePoolInstanceMetricDefinitions(resp *http.Response) (result ListMultiRolePoolInstanceMetricDefinitionsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListMultiRolePoolInstanceMetricDefinitionsOperationResponse, err error) {
			req, err := c.preparerForListMultiRolePoolInstanceMetricDefinitionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListMultiRolePoolInstanceMetricDefinitions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolInstanceMetricDefinitions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListMultiRolePoolInstanceMetricDefinitionsComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRolePoolInstanceMetricDefinitionsComplete(ctx context.Context, id DefaultInstanceId) (ListMultiRolePoolInstanceMetricDefinitionsCompleteResult, error) {
	return c.ListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate(ctx, id, ResourceMetricDefinitionOperationPredicate{})
}

// ListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListMultiRolePoolInstanceMetricDefinitionsCompleteMatchingPredicate(ctx context.Context, id DefaultInstanceId, predicate ResourceMetricDefinitionOperationPredicate) (resp ListMultiRolePoolInstanceMetricDefinitionsCompleteResult, err error) {
	items := make([]ResourceMetricDefinition, 0)

	page, err := c.ListMultiRolePoolInstanceMetricDefinitions(ctx, id)
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

	out := ListMultiRolePoolInstanceMetricDefinitionsCompleteResult{
		Items: items,
	}
	return out, nil
}
