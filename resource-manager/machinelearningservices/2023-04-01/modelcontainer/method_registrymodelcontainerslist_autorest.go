package modelcontainer

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

type RegistryModelContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ModelContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryModelContainersListOperationResponse, error)
}

type RegistryModelContainersListCompleteResult struct {
	Items []ModelContainerResource
}

func (r RegistryModelContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryModelContainersListOperationResponse) LoadMore(ctx context.Context) (resp RegistryModelContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryModelContainersListOperationOptions struct {
	ListViewType *ListViewType
	Skip         *string
}

func DefaultRegistryModelContainersListOperationOptions() RegistryModelContainersListOperationOptions {
	return RegistryModelContainersListOperationOptions{}
}

func (o RegistryModelContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryModelContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	return out
}

// RegistryModelContainersList ...
func (c ModelContainerClient) RegistryModelContainersList(ctx context.Context, id RegistryId, options RegistryModelContainersListOperationOptions) (resp RegistryModelContainersListOperationResponse, err error) {
	req, err := c.preparerForRegistryModelContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryModelContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryModelContainersList prepares the RegistryModelContainersList request.
func (c ModelContainerClient) preparerForRegistryModelContainersList(ctx context.Context, id RegistryId, options RegistryModelContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/models", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryModelContainersListWithNextLink prepares the RegistryModelContainersList request with the given nextLink token.
func (c ModelContainerClient) preparerForRegistryModelContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryModelContainersList handles the response to the RegistryModelContainersList request. The method always
// closes the http.Response Body.
func (c ModelContainerClient) responderForRegistryModelContainersList(resp *http.Response) (result RegistryModelContainersListOperationResponse, err error) {
	type page struct {
		Values   []ModelContainerResource `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryModelContainersListOperationResponse, err error) {
			req, err := c.preparerForRegistryModelContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryModelContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryModelContainersListComplete retrieves all of the results into a single object
func (c ModelContainerClient) RegistryModelContainersListComplete(ctx context.Context, id RegistryId, options RegistryModelContainersListOperationOptions) (RegistryModelContainersListCompleteResult, error) {
	return c.RegistryModelContainersListCompleteMatchingPredicate(ctx, id, options, ModelContainerResourceOperationPredicate{})
}

// RegistryModelContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ModelContainerClient) RegistryModelContainersListCompleteMatchingPredicate(ctx context.Context, id RegistryId, options RegistryModelContainersListOperationOptions, predicate ModelContainerResourceOperationPredicate) (resp RegistryModelContainersListCompleteResult, err error) {
	items := make([]ModelContainerResource, 0)

	page, err := c.RegistryModelContainersList(ctx, id, options)
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

	out := RegistryModelContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
