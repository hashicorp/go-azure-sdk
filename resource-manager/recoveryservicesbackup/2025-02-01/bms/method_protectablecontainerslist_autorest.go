package bms

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

type ProtectableContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProtectableContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ProtectableContainersListOperationResponse, error)
}

type ProtectableContainersListCompleteResult struct {
	Items []ProtectableContainerResource
}

func (r ProtectableContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ProtectableContainersListOperationResponse) LoadMore(ctx context.Context) (resp ProtectableContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ProtectableContainersListOperationOptions struct {
	Filter *string
}

func DefaultProtectableContainersListOperationOptions() ProtectableContainersListOperationOptions {
	return ProtectableContainersListOperationOptions{}
}

func (o ProtectableContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ProtectableContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ProtectableContainersList ...
func (c BmsClient) ProtectableContainersList(ctx context.Context, id BackupFabricId, options ProtectableContainersListOperationOptions) (resp ProtectableContainersListOperationResponse, err error) {
	req, err := c.preparerForProtectableContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForProtectableContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForProtectableContainersList prepares the ProtectableContainersList request.
func (c BmsClient) preparerForProtectableContainersList(ctx context.Context, id BackupFabricId, options ProtectableContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/protectableContainers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForProtectableContainersListWithNextLink prepares the ProtectableContainersList request with the given nextLink token.
func (c BmsClient) preparerForProtectableContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForProtectableContainersList handles the response to the ProtectableContainersList request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForProtectableContainersList(resp *http.Response) (result ProtectableContainersListOperationResponse, err error) {
	type page struct {
		Values   []ProtectableContainerResource `json:"value"`
		NextLink *string                        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ProtectableContainersListOperationResponse, err error) {
			req, err := c.preparerForProtectableContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForProtectableContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectableContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ProtectableContainersListComplete retrieves all of the results into a single object
func (c BmsClient) ProtectableContainersListComplete(ctx context.Context, id BackupFabricId, options ProtectableContainersListOperationOptions) (ProtectableContainersListCompleteResult, error) {
	return c.ProtectableContainersListCompleteMatchingPredicate(ctx, id, options, ProtectableContainerResourceOperationPredicate{})
}

// ProtectableContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BmsClient) ProtectableContainersListCompleteMatchingPredicate(ctx context.Context, id BackupFabricId, options ProtectableContainersListOperationOptions, predicate ProtectableContainerResourceOperationPredicate) (resp ProtectableContainersListCompleteResult, err error) {
	items := make([]ProtectableContainerResource, 0)

	page, err := c.ProtectableContainersList(ctx, id, options)
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

	out := ProtectableContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
