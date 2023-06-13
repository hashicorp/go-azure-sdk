package recoverypointscrr

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

type RecoveryPointsCrrListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RecoveryPointResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RecoveryPointsCrrListOperationResponse, error)
}

type RecoveryPointsCrrListCompleteResult struct {
	Items []RecoveryPointResource
}

func (r RecoveryPointsCrrListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RecoveryPointsCrrListOperationResponse) LoadMore(ctx context.Context) (resp RecoveryPointsCrrListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RecoveryPointsCrrListOperationOptions struct {
	Filter *string
}

func DefaultRecoveryPointsCrrListOperationOptions() RecoveryPointsCrrListOperationOptions {
	return RecoveryPointsCrrListOperationOptions{}
}

func (o RecoveryPointsCrrListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RecoveryPointsCrrListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// RecoveryPointsCrrList ...
func (c RecoveryPointsCrrClient) RecoveryPointsCrrList(ctx context.Context, id ProtectedItemId, options RecoveryPointsCrrListOperationOptions) (resp RecoveryPointsCrrListOperationResponse, err error) {
	req, err := c.preparerForRecoveryPointsCrrList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRecoveryPointsCrrList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRecoveryPointsCrrList prepares the RecoveryPointsCrrList request.
func (c RecoveryPointsCrrClient) preparerForRecoveryPointsCrrList(ctx context.Context, id ProtectedItemId, options RecoveryPointsCrrListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/recoveryPoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRecoveryPointsCrrListWithNextLink prepares the RecoveryPointsCrrList request with the given nextLink token.
func (c RecoveryPointsCrrClient) preparerForRecoveryPointsCrrListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRecoveryPointsCrrList handles the response to the RecoveryPointsCrrList request. The method always
// closes the http.Response Body.
func (c RecoveryPointsCrrClient) responderForRecoveryPointsCrrList(resp *http.Response) (result RecoveryPointsCrrListOperationResponse, err error) {
	type page struct {
		Values   []RecoveryPointResource `json:"value"`
		NextLink *string                 `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RecoveryPointsCrrListOperationResponse, err error) {
			req, err := c.preparerForRecoveryPointsCrrListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRecoveryPointsCrrList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypointscrr.RecoveryPointsCrrClient", "RecoveryPointsCrrList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RecoveryPointsCrrListComplete retrieves all of the results into a single object
func (c RecoveryPointsCrrClient) RecoveryPointsCrrListComplete(ctx context.Context, id ProtectedItemId, options RecoveryPointsCrrListOperationOptions) (RecoveryPointsCrrListCompleteResult, error) {
	return c.RecoveryPointsCrrListCompleteMatchingPredicate(ctx, id, options, RecoveryPointResourceOperationPredicate{})
}

// RecoveryPointsCrrListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecoveryPointsCrrClient) RecoveryPointsCrrListCompleteMatchingPredicate(ctx context.Context, id ProtectedItemId, options RecoveryPointsCrrListOperationOptions, predicate RecoveryPointResourceOperationPredicate) (resp RecoveryPointsCrrListCompleteResult, err error) {
	items := make([]RecoveryPointResource, 0)

	page, err := c.RecoveryPointsCrrList(ctx, id, options)
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

	out := RecoveryPointsCrrListCompleteResult{
		Items: items,
	}
	return out, nil
}
