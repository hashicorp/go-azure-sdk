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

type SecureScoreControlsListBySecureScoreOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreControlDetails

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecureScoreControlsListBySecureScoreOperationResponse, error)
}

type SecureScoreControlsListBySecureScoreCompleteResult struct {
	Items []SecureScoreControlDetails
}

func (r SecureScoreControlsListBySecureScoreOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecureScoreControlsListBySecureScoreOperationResponse) LoadMore(ctx context.Context) (resp SecureScoreControlsListBySecureScoreOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type SecureScoreControlsListBySecureScoreOperationOptions struct {
	Expand *ExpandControlsEnum
}

func DefaultSecureScoreControlsListBySecureScoreOperationOptions() SecureScoreControlsListBySecureScoreOperationOptions {
	return SecureScoreControlsListBySecureScoreOperationOptions{}
}

func (o SecureScoreControlsListBySecureScoreOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SecureScoreControlsListBySecureScoreOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// SecureScoreControlsListBySecureScore ...
func (c SecureScoreClient) SecureScoreControlsListBySecureScore(ctx context.Context, id SecureScoreId, options SecureScoreControlsListBySecureScoreOperationOptions) (resp SecureScoreControlsListBySecureScoreOperationResponse, err error) {
	req, err := c.preparerForSecureScoreControlsListBySecureScore(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecureScoreControlsListBySecureScore(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecureScoreControlsListBySecureScore prepares the SecureScoreControlsListBySecureScore request.
func (c SecureScoreClient) preparerForSecureScoreControlsListBySecureScore(ctx context.Context, id SecureScoreId, options SecureScoreControlsListBySecureScoreOperationOptions) (*http.Request, error) {
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

// preparerForSecureScoreControlsListBySecureScoreWithNextLink prepares the SecureScoreControlsListBySecureScore request with the given nextLink token.
func (c SecureScoreClient) preparerForSecureScoreControlsListBySecureScoreWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecureScoreControlsListBySecureScore handles the response to the SecureScoreControlsListBySecureScore request. The method always
// closes the http.Response Body.
func (c SecureScoreClient) responderForSecureScoreControlsListBySecureScore(resp *http.Response) (result SecureScoreControlsListBySecureScoreOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecureScoreControlsListBySecureScoreOperationResponse, err error) {
			req, err := c.preparerForSecureScoreControlsListBySecureScoreWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecureScoreControlsListBySecureScore(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoreControlsListBySecureScore", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecureScoreControlsListBySecureScoreComplete retrieves all of the results into a single object
func (c SecureScoreClient) SecureScoreControlsListBySecureScoreComplete(ctx context.Context, id SecureScoreId, options SecureScoreControlsListBySecureScoreOperationOptions) (SecureScoreControlsListBySecureScoreCompleteResult, error) {
	return c.SecureScoreControlsListBySecureScoreCompleteMatchingPredicate(ctx, id, options, SecureScoreControlDetailsOperationPredicate{})
}

// SecureScoreControlsListBySecureScoreCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreClient) SecureScoreControlsListBySecureScoreCompleteMatchingPredicate(ctx context.Context, id SecureScoreId, options SecureScoreControlsListBySecureScoreOperationOptions, predicate SecureScoreControlDetailsOperationPredicate) (resp SecureScoreControlsListBySecureScoreCompleteResult, err error) {
	items := make([]SecureScoreControlDetails, 0)

	page, err := c.SecureScoreControlsListBySecureScore(ctx, id, options)
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

	out := SecureScoreControlsListBySecureScoreCompleteResult{
		Items: items,
	}
	return out, nil
}
