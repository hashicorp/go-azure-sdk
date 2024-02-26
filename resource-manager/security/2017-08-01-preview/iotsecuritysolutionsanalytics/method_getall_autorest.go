package iotsecuritysolutionsanalytics

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

type GetAllOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IoTSecuritySolutionAnalyticsModel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAllOperationResponse, error)
}

type GetAllCompleteResult struct {
	Items []IoTSecuritySolutionAnalyticsModel
}

func (r GetAllOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAllOperationResponse) LoadMore(ctx context.Context) (resp GetAllOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetAll ...
func (c IoTSecuritySolutionsAnalyticsClient) GetAll(ctx context.Context, id IotSecuritySolutionId) (resp GetAllOperationResponse, err error) {
	req, err := c.preparerForGetAll(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAll(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAll prepares the GetAll request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForGetAll(ctx context.Context, id IotSecuritySolutionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/analyticsModels", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAllWithNextLink prepares the GetAll request with the given nextLink token.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForGetAllWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAll handles the response to the GetAll request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForGetAll(resp *http.Response) (result GetAllOperationResponse, err error) {
	type page struct {
		Values   []IoTSecuritySolutionAnalyticsModel `json:"value"`
		NextLink *string                             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAllOperationResponse, err error) {
			req, err := c.preparerForGetAllWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAll(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "GetAll", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAllComplete retrieves all of the results into a single object
func (c IoTSecuritySolutionsAnalyticsClient) GetAllComplete(ctx context.Context, id IotSecuritySolutionId) (GetAllCompleteResult, error) {
	return c.GetAllCompleteMatchingPredicate(ctx, id, IoTSecuritySolutionAnalyticsModelOperationPredicate{})
}

// GetAllCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IoTSecuritySolutionsAnalyticsClient) GetAllCompleteMatchingPredicate(ctx context.Context, id IotSecuritySolutionId, predicate IoTSecuritySolutionAnalyticsModelOperationPredicate) (resp GetAllCompleteResult, err error) {
	items := make([]IoTSecuritySolutionAnalyticsModel, 0)

	page, err := c.GetAll(ctx, id)
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

	out := GetAllCompleteResult{
		Items: items,
	}
	return out, nil
}
