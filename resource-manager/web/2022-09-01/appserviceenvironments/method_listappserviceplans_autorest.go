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

type ListAppServicePlansOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AppServicePlan

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAppServicePlansOperationResponse, error)
}

type ListAppServicePlansCompleteResult struct {
	Items []AppServicePlan
}

func (r ListAppServicePlansOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAppServicePlansOperationResponse) LoadMore(ctx context.Context) (resp ListAppServicePlansOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListAppServicePlans ...
func (c AppServiceEnvironmentsClient) ListAppServicePlans(ctx context.Context, id HostingEnvironmentId) (resp ListAppServicePlansOperationResponse, err error) {
	req, err := c.preparerForListAppServicePlans(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAppServicePlans(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAppServicePlans prepares the ListAppServicePlans request.
func (c AppServiceEnvironmentsClient) preparerForListAppServicePlans(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/serverFarms", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAppServicePlansWithNextLink prepares the ListAppServicePlans request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListAppServicePlansWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAppServicePlans handles the response to the ListAppServicePlans request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListAppServicePlans(resp *http.Response) (result ListAppServicePlansOperationResponse, err error) {
	type page struct {
		Values   []AppServicePlan `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAppServicePlansOperationResponse, err error) {
			req, err := c.preparerForListAppServicePlansWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAppServicePlans(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListAppServicePlans", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAppServicePlansComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListAppServicePlansComplete(ctx context.Context, id HostingEnvironmentId) (ListAppServicePlansCompleteResult, error) {
	return c.ListAppServicePlansCompleteMatchingPredicate(ctx, id, AppServicePlanOperationPredicate{})
}

// ListAppServicePlansCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListAppServicePlansCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate AppServicePlanOperationPredicate) (resp ListAppServicePlansCompleteResult, err error) {
	items := make([]AppServicePlan, 0)

	page, err := c.ListAppServicePlans(ctx, id)
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

	out := ListAppServicePlansCompleteResult{
		Items: items,
	}
	return out, nil
}
