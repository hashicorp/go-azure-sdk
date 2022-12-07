package appplatform

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

type BindingsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BindingResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BindingsListOperationResponse, error)
}

type BindingsListCompleteResult struct {
	Items []BindingResource
}

func (r BindingsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BindingsListOperationResponse) LoadMore(ctx context.Context) (resp BindingsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// BindingsList ...
func (c AppPlatformClient) BindingsList(ctx context.Context, id AppId) (resp BindingsListOperationResponse, err error) {
	req, err := c.preparerForBindingsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBindingsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBindingsList prepares the BindingsList request.
func (c AppPlatformClient) preparerForBindingsList(ctx context.Context, id AppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/bindings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBindingsListWithNextLink prepares the BindingsList request with the given nextLink token.
func (c AppPlatformClient) preparerForBindingsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBindingsList handles the response to the BindingsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForBindingsList(resp *http.Response) (result BindingsListOperationResponse, err error) {
	type page struct {
		Values   []BindingResource `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BindingsListOperationResponse, err error) {
			req, err := c.preparerForBindingsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBindingsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BindingsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BindingsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) BindingsListComplete(ctx context.Context, id AppId) (BindingsListCompleteResult, error) {
	return c.BindingsListCompleteMatchingPredicate(ctx, id, BindingResourceOperationPredicate{})
}

// BindingsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) BindingsListCompleteMatchingPredicate(ctx context.Context, id AppId, predicate BindingResourceOperationPredicate) (resp BindingsListCompleteResult, err error) {
	items := make([]BindingResource, 0)

	page, err := c.BindingsList(ctx, id)
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

	out := BindingsListCompleteResult{
		Items: items,
	}
	return out, nil
}
