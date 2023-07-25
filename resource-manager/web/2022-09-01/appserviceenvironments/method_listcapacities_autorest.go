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

type ListCapacitiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StampCapacity

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListCapacitiesOperationResponse, error)
}

type ListCapacitiesCompleteResult struct {
	Items []StampCapacity
}

func (r ListCapacitiesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListCapacitiesOperationResponse) LoadMore(ctx context.Context) (resp ListCapacitiesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListCapacities ...
func (c AppServiceEnvironmentsClient) ListCapacities(ctx context.Context, id HostingEnvironmentId) (resp ListCapacitiesOperationResponse, err error) {
	req, err := c.preparerForListCapacities(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListCapacities(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListCapacities prepares the ListCapacities request.
func (c AppServiceEnvironmentsClient) preparerForListCapacities(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capacities/compute", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListCapacitiesWithNextLink prepares the ListCapacities request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListCapacitiesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListCapacities handles the response to the ListCapacities request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListCapacities(resp *http.Response) (result ListCapacitiesOperationResponse, err error) {
	type page struct {
		Values   []StampCapacity `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListCapacitiesOperationResponse, err error) {
			req, err := c.preparerForListCapacitiesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListCapacities(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListCapacities", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListCapacitiesComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListCapacitiesComplete(ctx context.Context, id HostingEnvironmentId) (ListCapacitiesCompleteResult, error) {
	return c.ListCapacitiesCompleteMatchingPredicate(ctx, id, StampCapacityOperationPredicate{})
}

// ListCapacitiesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListCapacitiesCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate StampCapacityOperationPredicate) (resp ListCapacitiesCompleteResult, err error) {
	items := make([]StampCapacity, 0)

	page, err := c.ListCapacities(ctx, id)
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

	out := ListCapacitiesCompleteResult{
		Items: items,
	}
	return out, nil
}
