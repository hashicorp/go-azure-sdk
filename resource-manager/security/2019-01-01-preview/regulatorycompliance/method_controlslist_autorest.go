package regulatorycompliance

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

type ControlsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RegulatoryComplianceControl

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ControlsListOperationResponse, error)
}

type ControlsListCompleteResult struct {
	Items []RegulatoryComplianceControl
}

func (r ControlsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ControlsListOperationResponse) LoadMore(ctx context.Context) (resp ControlsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ControlsListOperationOptions struct {
	Filter *string
}

func DefaultControlsListOperationOptions() ControlsListOperationOptions {
	return ControlsListOperationOptions{}
}

func (o ControlsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ControlsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ControlsList ...
func (c RegulatoryComplianceClient) ControlsList(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions) (resp ControlsListOperationResponse, err error) {
	req, err := c.preparerForControlsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForControlsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForControlsList prepares the ControlsList request.
func (c RegulatoryComplianceClient) preparerForControlsList(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/regulatoryComplianceControls", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForControlsListWithNextLink prepares the ControlsList request with the given nextLink token.
func (c RegulatoryComplianceClient) preparerForControlsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForControlsList handles the response to the ControlsList request. The method always
// closes the http.Response Body.
func (c RegulatoryComplianceClient) responderForControlsList(resp *http.Response) (result ControlsListOperationResponse, err error) {
	type page struct {
		Values   []RegulatoryComplianceControl `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ControlsListOperationResponse, err error) {
			req, err := c.preparerForControlsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForControlsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "ControlsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ControlsListComplete retrieves all of the results into a single object
func (c RegulatoryComplianceClient) ControlsListComplete(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions) (ControlsListCompleteResult, error) {
	return c.ControlsListCompleteMatchingPredicate(ctx, id, options, RegulatoryComplianceControlOperationPredicate{})
}

// ControlsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RegulatoryComplianceClient) ControlsListCompleteMatchingPredicate(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions, predicate RegulatoryComplianceControlOperationPredicate) (resp ControlsListCompleteResult, err error) {
	items := make([]RegulatoryComplianceControl, 0)

	page, err := c.ControlsList(ctx, id, options)
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

	out := ControlsListCompleteResult{
		Items: items,
	}
	return out, nil
}
