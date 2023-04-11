package securescorecontrols

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreControlDetails

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecureScoreControlsListOperationResponse, error)
}

type SecureScoreControlsListCompleteResult struct {
	Items []SecureScoreControlDetails
}

func (r SecureScoreControlsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecureScoreControlsListOperationResponse) LoadMore(ctx context.Context) (resp SecureScoreControlsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type SecureScoreControlsListOperationOptions struct {
	Expand *ExpandControlsEnum
}

func DefaultSecureScoreControlsListOperationOptions() SecureScoreControlsListOperationOptions {
	return SecureScoreControlsListOperationOptions{}
}

func (o SecureScoreControlsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SecureScoreControlsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// SecureScoreControlsList ...
func (c SecureScoreControlsClient) SecureScoreControlsList(ctx context.Context, id commonids.SubscriptionId, options SecureScoreControlsListOperationOptions) (resp SecureScoreControlsListOperationResponse, err error) {
	req, err := c.preparerForSecureScoreControlsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecureScoreControlsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecureScoreControlsList prepares the SecureScoreControlsList request.
func (c SecureScoreControlsClient) preparerForSecureScoreControlsList(ctx context.Context, id commonids.SubscriptionId, options SecureScoreControlsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/secureScoreControls", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSecureScoreControlsListWithNextLink prepares the SecureScoreControlsList request with the given nextLink token.
func (c SecureScoreControlsClient) preparerForSecureScoreControlsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecureScoreControlsList handles the response to the SecureScoreControlsList request. The method always
// closes the http.Response Body.
func (c SecureScoreControlsClient) responderForSecureScoreControlsList(resp *http.Response) (result SecureScoreControlsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecureScoreControlsListOperationResponse, err error) {
			req, err := c.preparerForSecureScoreControlsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecureScoreControlsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontrols.SecureScoreControlsClient", "SecureScoreControlsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecureScoreControlsListComplete retrieves all of the results into a single object
func (c SecureScoreControlsClient) SecureScoreControlsListComplete(ctx context.Context, id commonids.SubscriptionId, options SecureScoreControlsListOperationOptions) (SecureScoreControlsListCompleteResult, error) {
	return c.SecureScoreControlsListCompleteMatchingPredicate(ctx, id, options, SecureScoreControlDetailsOperationPredicate{})
}

// SecureScoreControlsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreControlsClient) SecureScoreControlsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options SecureScoreControlsListOperationOptions, predicate SecureScoreControlDetailsOperationPredicate) (resp SecureScoreControlsListCompleteResult, err error) {
	items := make([]SecureScoreControlDetails, 0)

	page, err := c.SecureScoreControlsList(ctx, id, options)
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

	out := SecureScoreControlsListCompleteResult{
		Items: items,
	}
	return out, nil
}
