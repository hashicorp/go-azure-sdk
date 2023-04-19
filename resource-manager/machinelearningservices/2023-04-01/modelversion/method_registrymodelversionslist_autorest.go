package modelversion

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

type RegistryModelVersionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ModelVersionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryModelVersionsListOperationResponse, error)
}

type RegistryModelVersionsListCompleteResult struct {
	Items []ModelVersionResource
}

func (r RegistryModelVersionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryModelVersionsListOperationResponse) LoadMore(ctx context.Context) (resp RegistryModelVersionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryModelVersionsListOperationOptions struct {
	Description  *string
	ListViewType *ListViewType
	OrderBy      *string
	Properties   *string
	Skip         *string
	Tags         *string
	Top          *int64
	Version      *string
}

func DefaultRegistryModelVersionsListOperationOptions() RegistryModelVersionsListOperationOptions {
	return RegistryModelVersionsListOperationOptions{}
}

func (o RegistryModelVersionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryModelVersionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Description != nil {
		out["description"] = *o.Description
	}

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Properties != nil {
		out["properties"] = *o.Properties
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Tags != nil {
		out["tags"] = *o.Tags
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	if o.Version != nil {
		out["version"] = *o.Version
	}

	return out
}

// RegistryModelVersionsList ...
func (c ModelVersionClient) RegistryModelVersionsList(ctx context.Context, id RegistryModelId, options RegistryModelVersionsListOperationOptions) (resp RegistryModelVersionsListOperationResponse, err error) {
	req, err := c.preparerForRegistryModelVersionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryModelVersionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryModelVersionsList prepares the RegistryModelVersionsList request.
func (c ModelVersionClient) preparerForRegistryModelVersionsList(ctx context.Context, id RegistryModelId, options RegistryModelVersionsListOperationOptions) (*http.Request, error) {
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

// preparerForRegistryModelVersionsListWithNextLink prepares the RegistryModelVersionsList request with the given nextLink token.
func (c ModelVersionClient) preparerForRegistryModelVersionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistryModelVersionsList handles the response to the RegistryModelVersionsList request. The method always
// closes the http.Response Body.
func (c ModelVersionClient) responderForRegistryModelVersionsList(resp *http.Response) (result RegistryModelVersionsListOperationResponse, err error) {
	type page struct {
		Values   []ModelVersionResource `json:"value"`
		NextLink *string                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryModelVersionsListOperationResponse, err error) {
			req, err := c.preparerForRegistryModelVersionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryModelVersionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryModelVersionsListComplete retrieves all of the results into a single object
func (c ModelVersionClient) RegistryModelVersionsListComplete(ctx context.Context, id RegistryModelId, options RegistryModelVersionsListOperationOptions) (RegistryModelVersionsListCompleteResult, error) {
	return c.RegistryModelVersionsListCompleteMatchingPredicate(ctx, id, options, ModelVersionResourceOperationPredicate{})
}

// RegistryModelVersionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ModelVersionClient) RegistryModelVersionsListCompleteMatchingPredicate(ctx context.Context, id RegistryModelId, options RegistryModelVersionsListOperationOptions, predicate ModelVersionResourceOperationPredicate) (resp RegistryModelVersionsListCompleteResult, err error) {
	items := make([]ModelVersionResource, 0)

	page, err := c.RegistryModelVersionsList(ctx, id, options)
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

	out := RegistryModelVersionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
