package experiments

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

type ListExecutionDetailsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ExperimentExecutionDetails

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListExecutionDetailsOperationResponse, error)
}

type ListExecutionDetailsCompleteResult struct {
	Items []ExperimentExecutionDetails
}

func (r ListExecutionDetailsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListExecutionDetailsOperationResponse) LoadMore(ctx context.Context) (resp ListExecutionDetailsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListExecutionDetails ...
func (c ExperimentsClient) ListExecutionDetails(ctx context.Context, id ExperimentId) (resp ListExecutionDetailsOperationResponse, err error) {
	req, err := c.preparerForListExecutionDetails(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListExecutionDetails(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListExecutionDetails prepares the ListExecutionDetails request.
func (c ExperimentsClient) preparerForListExecutionDetails(ctx context.Context, id ExperimentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/executionDetails", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListExecutionDetailsWithNextLink prepares the ListExecutionDetails request with the given nextLink token.
func (c ExperimentsClient) preparerForListExecutionDetailsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListExecutionDetails handles the response to the ListExecutionDetails request. The method always
// closes the http.Response Body.
func (c ExperimentsClient) responderForListExecutionDetails(resp *http.Response) (result ListExecutionDetailsOperationResponse, err error) {
	type page struct {
		Values   []ExperimentExecutionDetails `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListExecutionDetailsOperationResponse, err error) {
			req, err := c.preparerForListExecutionDetailsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListExecutionDetails(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListExecutionDetails", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListExecutionDetailsComplete retrieves all of the results into a single object
func (c ExperimentsClient) ListExecutionDetailsComplete(ctx context.Context, id ExperimentId) (ListExecutionDetailsCompleteResult, error) {
	return c.ListExecutionDetailsCompleteMatchingPredicate(ctx, id, ExperimentExecutionDetailsOperationPredicate{})
}

// ListExecutionDetailsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ExperimentsClient) ListExecutionDetailsCompleteMatchingPredicate(ctx context.Context, id ExperimentId, predicate ExperimentExecutionDetailsOperationPredicate) (resp ListExecutionDetailsCompleteResult, err error) {
	items := make([]ExperimentExecutionDetails, 0)

	page, err := c.ListExecutionDetails(ctx, id)
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

	out := ListExecutionDetailsCompleteResult{
		Items: items,
	}
	return out, nil
}
