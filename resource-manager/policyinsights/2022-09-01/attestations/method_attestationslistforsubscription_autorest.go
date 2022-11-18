package attestations

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

type AttestationsListForSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Attestation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AttestationsListForSubscriptionOperationResponse, error)
}

type AttestationsListForSubscriptionCompleteResult struct {
	Items []Attestation
}

func (r AttestationsListForSubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AttestationsListForSubscriptionOperationResponse) LoadMore(ctx context.Context) (resp AttestationsListForSubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type AttestationsListForSubscriptionOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultAttestationsListForSubscriptionOperationOptions() AttestationsListForSubscriptionOperationOptions {
	return AttestationsListForSubscriptionOperationOptions{}
}

func (o AttestationsListForSubscriptionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AttestationsListForSubscriptionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// AttestationsListForSubscription ...
func (c AttestationsClient) AttestationsListForSubscription(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions) (resp AttestationsListForSubscriptionOperationResponse, err error) {
	req, err := c.preparerForAttestationsListForSubscription(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAttestationsListForSubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAttestationsListForSubscription prepares the AttestationsListForSubscription request.
func (c AttestationsClient) preparerForAttestationsListForSubscription(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/attestations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAttestationsListForSubscriptionWithNextLink prepares the AttestationsListForSubscription request with the given nextLink token.
func (c AttestationsClient) preparerForAttestationsListForSubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAttestationsListForSubscription handles the response to the AttestationsListForSubscription request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsListForSubscription(resp *http.Response) (result AttestationsListForSubscriptionOperationResponse, err error) {
	type page struct {
		Values   []Attestation `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AttestationsListForSubscriptionOperationResponse, err error) {
			req, err := c.preparerForAttestationsListForSubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAttestationsListForSubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForSubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AttestationsListForSubscriptionComplete retrieves all of the results into a single object
func (c AttestationsClient) AttestationsListForSubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions) (AttestationsListForSubscriptionCompleteResult, error) {
	return c.AttestationsListForSubscriptionCompleteMatchingPredicate(ctx, id, options, AttestationOperationPredicate{})
}

// AttestationsListForSubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AttestationsClient) AttestationsListForSubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions, predicate AttestationOperationPredicate) (resp AttestationsListForSubscriptionCompleteResult, err error) {
	items := make([]Attestation, 0)

	page, err := c.AttestationsListForSubscription(ctx, id, options)
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

	out := AttestationsListForSubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
