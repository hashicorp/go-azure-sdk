package authorizations

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

type AuthorizationListByAuthorizationProviderOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AuthorizationContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AuthorizationListByAuthorizationProviderOperationResponse, error)
}

type AuthorizationListByAuthorizationProviderCompleteResult struct {
	Items []AuthorizationContract
}

func (r AuthorizationListByAuthorizationProviderOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AuthorizationListByAuthorizationProviderOperationResponse) LoadMore(ctx context.Context) (resp AuthorizationListByAuthorizationProviderOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type AuthorizationListByAuthorizationProviderOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultAuthorizationListByAuthorizationProviderOperationOptions() AuthorizationListByAuthorizationProviderOperationOptions {
	return AuthorizationListByAuthorizationProviderOperationOptions{}
}

func (o AuthorizationListByAuthorizationProviderOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AuthorizationListByAuthorizationProviderOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// AuthorizationListByAuthorizationProvider ...
func (c AuthorizationsClient) AuthorizationListByAuthorizationProvider(ctx context.Context, id AuthorizationProviderId, options AuthorizationListByAuthorizationProviderOperationOptions) (resp AuthorizationListByAuthorizationProviderOperationResponse, err error) {
	req, err := c.preparerForAuthorizationListByAuthorizationProvider(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAuthorizationListByAuthorizationProvider(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAuthorizationListByAuthorizationProvider prepares the AuthorizationListByAuthorizationProvider request.
func (c AuthorizationsClient) preparerForAuthorizationListByAuthorizationProvider(ctx context.Context, id AuthorizationProviderId, options AuthorizationListByAuthorizationProviderOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/authorizations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAuthorizationListByAuthorizationProviderWithNextLink prepares the AuthorizationListByAuthorizationProvider request with the given nextLink token.
func (c AuthorizationsClient) preparerForAuthorizationListByAuthorizationProviderWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAuthorizationListByAuthorizationProvider handles the response to the AuthorizationListByAuthorizationProvider request. The method always
// closes the http.Response Body.
func (c AuthorizationsClient) responderForAuthorizationListByAuthorizationProvider(resp *http.Response) (result AuthorizationListByAuthorizationProviderOperationResponse, err error) {
	type page struct {
		Values   []AuthorizationContract `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AuthorizationListByAuthorizationProviderOperationResponse, err error) {
			req, err := c.preparerForAuthorizationListByAuthorizationProviderWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAuthorizationListByAuthorizationProvider(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizations.AuthorizationsClient", "AuthorizationListByAuthorizationProvider", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AuthorizationListByAuthorizationProviderComplete retrieves all of the results into a single object
func (c AuthorizationsClient) AuthorizationListByAuthorizationProviderComplete(ctx context.Context, id AuthorizationProviderId, options AuthorizationListByAuthorizationProviderOperationOptions) (AuthorizationListByAuthorizationProviderCompleteResult, error) {
	return c.AuthorizationListByAuthorizationProviderCompleteMatchingPredicate(ctx, id, options, AuthorizationContractOperationPredicate{})
}

// AuthorizationListByAuthorizationProviderCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AuthorizationsClient) AuthorizationListByAuthorizationProviderCompleteMatchingPredicate(ctx context.Context, id AuthorizationProviderId, options AuthorizationListByAuthorizationProviderOperationOptions, predicate AuthorizationContractOperationPredicate) (resp AuthorizationListByAuthorizationProviderCompleteResult, err error) {
	items := make([]AuthorizationContract, 0)

	page, err := c.AuthorizationListByAuthorizationProvider(ctx, id, options)
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

	out := AuthorizationListByAuthorizationProviderCompleteResult{
		Items: items,
	}
	return out, nil
}
