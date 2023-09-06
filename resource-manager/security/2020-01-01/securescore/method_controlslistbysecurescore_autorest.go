package securescore

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

type ControlsListBySecureScoreOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreControlDetails

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ControlsListBySecureScoreOperationResponse, error)
}

type ControlsListBySecureScoreCompleteResult struct {
	Items []SecureScoreControlDetails
}

func (r ControlsListBySecureScoreOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ControlsListBySecureScoreOperationResponse) LoadMore(ctx context.Context) (resp ControlsListBySecureScoreOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ControlsListBySecureScoreOperationOptions struct {
	Expand *ExpandControlsEnum
}

func DefaultControlsListBySecureScoreOperationOptions() ControlsListBySecureScoreOperationOptions {
	return ControlsListBySecureScoreOperationOptions{}
}

func (o ControlsListBySecureScoreOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ControlsListBySecureScoreOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// ControlsListBySecureScore ...
func (c SecureScoreClient) ControlsListBySecureScore(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions) (resp ControlsListBySecureScoreOperationResponse, err error) {
	req, err := c.preparerForControlsListBySecureScore(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForControlsListBySecureScore(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForControlsListBySecureScore prepares the ControlsListBySecureScore request.
func (c SecureScoreClient) preparerForControlsListBySecureScore(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/secureScoreControls", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForControlsListBySecureScoreWithNextLink prepares the ControlsListBySecureScore request with the given nextLink token.
func (c SecureScoreClient) preparerForControlsListBySecureScoreWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForControlsListBySecureScore handles the response to the ControlsListBySecureScore request. The method always
// closes the http.Response Body.
func (c SecureScoreClient) responderForControlsListBySecureScore(resp *http.Response) (result ControlsListBySecureScoreOperationResponse, err error) {
	type page struct {
		Values   []SecureScoreControlDetails `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ControlsListBySecureScoreOperationResponse, err error) {
			req, err := c.preparerForControlsListBySecureScoreWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForControlsListBySecureScore(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "ControlsListBySecureScore", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ControlsListBySecureScoreComplete retrieves all of the results into a single object
func (c SecureScoreClient) ControlsListBySecureScoreComplete(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions) (ControlsListBySecureScoreCompleteResult, error) {
	return c.ControlsListBySecureScoreCompleteMatchingPredicate(ctx, id, options, SecureScoreControlDetailsOperationPredicate{})
}

// ControlsListBySecureScoreCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreClient) ControlsListBySecureScoreCompleteMatchingPredicate(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions, predicate SecureScoreControlDetailsOperationPredicate) (resp ControlsListBySecureScoreCompleteResult, err error) {
	items := make([]SecureScoreControlDetails, 0)

	page, err := c.ControlsListBySecureScore(ctx, id, options)
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

	out := ControlsListBySecureScoreCompleteResult{
		Items: items,
	}
	return out, nil
}
