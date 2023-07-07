package locationbasedcapabilities

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

type ExecuteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FlexibleServerCapability

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ExecuteOperationResponse, error)
}

type ExecuteCompleteResult struct {
	Items []FlexibleServerCapability
}

func (r ExecuteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ExecuteOperationResponse) LoadMore(ctx context.Context) (resp ExecuteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// Execute ...
func (c LocationBasedCapabilitiesClient) Execute(ctx context.Context, id LocationId) (resp ExecuteOperationResponse, err error) {
	req, err := c.preparerForExecute(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForExecute(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForExecute prepares the Execute request.
func (c LocationBasedCapabilitiesClient) preparerForExecute(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capabilities", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForExecuteWithNextLink prepares the Execute request with the given nextLink token.
func (c LocationBasedCapabilitiesClient) preparerForExecuteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForExecute handles the response to the Execute request. The method always
// closes the http.Response Body.
func (c LocationBasedCapabilitiesClient) responderForExecute(resp *http.Response) (result ExecuteOperationResponse, err error) {
	type page struct {
		Values   []FlexibleServerCapability `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ExecuteOperationResponse, err error) {
			req, err := c.preparerForExecuteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForExecute(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "locationbasedcapabilities.LocationBasedCapabilitiesClient", "Execute", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ExecuteComplete retrieves all of the results into a single object
func (c LocationBasedCapabilitiesClient) ExecuteComplete(ctx context.Context, id LocationId) (ExecuteCompleteResult, error) {
	return c.ExecuteCompleteMatchingPredicate(ctx, id, FlexibleServerCapabilityOperationPredicate{})
}

// ExecuteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c LocationBasedCapabilitiesClient) ExecuteCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate FlexibleServerCapabilityOperationPredicate) (resp ExecuteCompleteResult, err error) {
	items := make([]FlexibleServerCapability, 0)

	page, err := c.Execute(ctx, id)
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

	out := ExecuteCompleteResult{
		Items: items,
	}
	return out, nil
}
