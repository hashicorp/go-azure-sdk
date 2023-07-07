package flexibleservercapabilities

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

type ServerCapabilitiesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FlexibleServerCapability

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ServerCapabilitiesListOperationResponse, error)
}

type ServerCapabilitiesListCompleteResult struct {
	Items []FlexibleServerCapability
}

func (r ServerCapabilitiesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ServerCapabilitiesListOperationResponse) LoadMore(ctx context.Context) (resp ServerCapabilitiesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ServerCapabilitiesList ...
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesList(ctx context.Context, id FlexibleServerId) (resp ServerCapabilitiesListOperationResponse, err error) {
	req, err := c.preparerForServerCapabilitiesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForServerCapabilitiesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForServerCapabilitiesList prepares the ServerCapabilitiesList request.
func (c FlexibleServerCapabilitiesClient) preparerForServerCapabilitiesList(ctx context.Context, id FlexibleServerId) (*http.Request, error) {
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

// preparerForServerCapabilitiesListWithNextLink prepares the ServerCapabilitiesList request with the given nextLink token.
func (c FlexibleServerCapabilitiesClient) preparerForServerCapabilitiesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForServerCapabilitiesList handles the response to the ServerCapabilitiesList request. The method always
// closes the http.Response Body.
func (c FlexibleServerCapabilitiesClient) responderForServerCapabilitiesList(resp *http.Response) (result ServerCapabilitiesListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ServerCapabilitiesListOperationResponse, err error) {
			req, err := c.preparerForServerCapabilitiesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForServerCapabilitiesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "flexibleservercapabilities.FlexibleServerCapabilitiesClient", "ServerCapabilitiesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ServerCapabilitiesListComplete retrieves all of the results into a single object
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesListComplete(ctx context.Context, id FlexibleServerId) (ServerCapabilitiesListCompleteResult, error) {
	return c.ServerCapabilitiesListCompleteMatchingPredicate(ctx, id, FlexibleServerCapabilityOperationPredicate{})
}

// ServerCapabilitiesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c FlexibleServerCapabilitiesClient) ServerCapabilitiesListCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate FlexibleServerCapabilityOperationPredicate) (resp ServerCapabilitiesListCompleteResult, err error) {
	items := make([]FlexibleServerCapability, 0)

	page, err := c.ServerCapabilitiesList(ctx, id)
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

	out := ServerCapabilitiesListCompleteResult{
		Items: items,
	}
	return out, nil
}
