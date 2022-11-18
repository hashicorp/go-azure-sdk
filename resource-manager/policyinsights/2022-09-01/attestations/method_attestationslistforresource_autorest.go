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

type AttestationsListForResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Attestation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AttestationsListForResourceOperationResponse, error)
}

type AttestationsListForResourceCompleteResult struct {
	Items []Attestation
}

func (r AttestationsListForResourceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AttestationsListForResourceOperationResponse) LoadMore(ctx context.Context) (resp AttestationsListForResourceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type AttestationsListForResourceOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultAttestationsListForResourceOperationOptions() AttestationsListForResourceOperationOptions {
	return AttestationsListForResourceOperationOptions{}
}

func (o AttestationsListForResourceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AttestationsListForResourceOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// AttestationsListForResource ...
func (c AttestationsClient) AttestationsListForResource(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions) (resp AttestationsListForResourceOperationResponse, err error) {
	req, err := c.preparerForAttestationsListForResource(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAttestationsListForResource(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAttestationsListForResource prepares the AttestationsListForResource request.
func (c AttestationsClient) preparerForAttestationsListForResource(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions) (*http.Request, error) {
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

// preparerForAttestationsListForResourceWithNextLink prepares the AttestationsListForResource request with the given nextLink token.
func (c AttestationsClient) preparerForAttestationsListForResourceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAttestationsListForResource handles the response to the AttestationsListForResource request. The method always
// closes the http.Response Body.
func (c AttestationsClient) responderForAttestationsListForResource(resp *http.Response) (result AttestationsListForResourceOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AttestationsListForResourceOperationResponse, err error) {
			req, err := c.preparerForAttestationsListForResourceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAttestationsListForResource(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "attestations.AttestationsClient", "AttestationsListForResource", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AttestationsListForResourceComplete retrieves all of the results into a single object
func (c AttestationsClient) AttestationsListForResourceComplete(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions) (AttestationsListForResourceCompleteResult, error) {
	return c.AttestationsListForResourceCompleteMatchingPredicate(ctx, id, options, AttestationOperationPredicate{})
}

// AttestationsListForResourceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AttestationsClient) AttestationsListForResourceCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options AttestationsListForResourceOperationOptions, predicate AttestationOperationPredicate) (resp AttestationsListForResourceCompleteResult, err error) {
	items := make([]Attestation, 0)

	page, err := c.AttestationsListForResource(ctx, id, options)
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

	out := AttestationsListForResourceCompleteResult{
		Items: items,
	}
	return out, nil
}
