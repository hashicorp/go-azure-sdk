package datacontainerregistry

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

type RegistryDataContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DataContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryDataContainersListOperationResponse, error)
}

type RegistryDataContainersListCompleteResult struct {
	Items []DataContainerResource
}

func (r RegistryDataContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryDataContainersListOperationResponse) LoadMore(ctx context.Context) (resp RegistryDataContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryDataContainersListOperationOptions struct {
	ListViewType *ListViewType
	Skip         *string
}

func DefaultRegistryDataContainersListOperationOptions() RegistryDataContainersListOperationOptions {
	return RegistryDataContainersListOperationOptions{}
}

func (o RegistryDataContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryDataContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	return out
}

// RegistryDataContainersList ...
func (c DataContainerRegistryClient) RegistryDataContainersList(ctx context.Context, id RegistryId, options RegistryDataContainersListOperationOptions) (resp RegistryDataContainersListOperationResponse, err error) {
	req, err := c.preparerForRegistryDataContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryDataContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryDataContainersList prepares the RegistryDataContainersList request.
func (c DataContainerRegistryClient) preparerForRegistryDataContainersList(ctx context.Context, id RegistryId, options RegistryDataContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/data", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryDataContainersListWithNextLink prepares the RegistryDataContainersList request with the given nextLink token.
func (c DataContainerRegistryClient) preparerForRegistryDataContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryDataContainersList handles the response to the RegistryDataContainersList request. The method always
// closes the http.Response Body.
func (c DataContainerRegistryClient) responderForRegistryDataContainersList(resp *http.Response) (result RegistryDataContainersListOperationResponse, err error) {
	type page struct {
		Values   []DataContainerResource `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryDataContainersListOperationResponse, err error) {
			req, err := c.preparerForRegistryDataContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryDataContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryDataContainersListComplete retrieves all of the results into a single object
func (c DataContainerRegistryClient) RegistryDataContainersListComplete(ctx context.Context, id RegistryId, options RegistryDataContainersListOperationOptions) (RegistryDataContainersListCompleteResult, error) {
	return c.RegistryDataContainersListCompleteMatchingPredicate(ctx, id, options, DataContainerResourceOperationPredicate{})
}

// RegistryDataContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DataContainerRegistryClient) RegistryDataContainersListCompleteMatchingPredicate(ctx context.Context, id RegistryId, options RegistryDataContainersListOperationOptions, predicate DataContainerResourceOperationPredicate) (resp RegistryDataContainersListCompleteResult, err error) {
	items := make([]DataContainerResource, 0)

	page, err := c.RegistryDataContainersList(ctx, id, options)
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

	out := RegistryDataContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
