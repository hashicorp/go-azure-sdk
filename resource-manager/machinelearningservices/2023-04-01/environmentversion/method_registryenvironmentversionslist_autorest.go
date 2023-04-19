package environmentversion

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

type RegistryEnvironmentVersionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]EnvironmentVersionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryEnvironmentVersionsListOperationResponse, error)
}

type RegistryEnvironmentVersionsListCompleteResult struct {
	Items []EnvironmentVersionResource
}

func (r RegistryEnvironmentVersionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryEnvironmentVersionsListOperationResponse) LoadMore(ctx context.Context) (resp RegistryEnvironmentVersionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryEnvironmentVersionsListOperationOptions struct {
	ListViewType *ListViewType
	OrderBy      *string
	Skip         *string
	Top          *int64
}

func DefaultRegistryEnvironmentVersionsListOperationOptions() RegistryEnvironmentVersionsListOperationOptions {
	return RegistryEnvironmentVersionsListOperationOptions{}
}

func (o RegistryEnvironmentVersionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryEnvironmentVersionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// RegistryEnvironmentVersionsList ...
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsList(ctx context.Context, id RegistryEnvironmentId, options RegistryEnvironmentVersionsListOperationOptions) (resp RegistryEnvironmentVersionsListOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentVersionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryEnvironmentVersionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryEnvironmentVersionsList prepares the RegistryEnvironmentVersionsList request.
func (c EnvironmentVersionClient) preparerForRegistryEnvironmentVersionsList(ctx context.Context, id RegistryEnvironmentId, options RegistryEnvironmentVersionsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryEnvironmentVersionsListWithNextLink prepares the RegistryEnvironmentVersionsList request with the given nextLink token.
func (c EnvironmentVersionClient) preparerForRegistryEnvironmentVersionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryEnvironmentVersionsList handles the response to the RegistryEnvironmentVersionsList request. The method always
// closes the http.Response Body.
func (c EnvironmentVersionClient) responderForRegistryEnvironmentVersionsList(resp *http.Response) (result RegistryEnvironmentVersionsListOperationResponse, err error) {
	type page struct {
		Values   []EnvironmentVersionResource `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryEnvironmentVersionsListOperationResponse, err error) {
			req, err := c.preparerForRegistryEnvironmentVersionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryEnvironmentVersionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryEnvironmentVersionsListComplete retrieves all of the results into a single object
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsListComplete(ctx context.Context, id RegistryEnvironmentId, options RegistryEnvironmentVersionsListOperationOptions) (RegistryEnvironmentVersionsListCompleteResult, error) {
	return c.RegistryEnvironmentVersionsListCompleteMatchingPredicate(ctx, id, options, EnvironmentVersionResourceOperationPredicate{})
}

// RegistryEnvironmentVersionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsListCompleteMatchingPredicate(ctx context.Context, id RegistryEnvironmentId, options RegistryEnvironmentVersionsListOperationOptions, predicate EnvironmentVersionResourceOperationPredicate) (resp RegistryEnvironmentVersionsListCompleteResult, err error) {
	items := make([]EnvironmentVersionResource, 0)

	page, err := c.RegistryEnvironmentVersionsList(ctx, id, options)
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

	out := RegistryEnvironmentVersionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
