package softdeletedcontainers

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

type DeletedProtectionContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProtectionContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (DeletedProtectionContainersListOperationResponse, error)
}

type DeletedProtectionContainersListCompleteResult struct {
	Items []ProtectionContainerResource
}

func (r DeletedProtectionContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r DeletedProtectionContainersListOperationResponse) LoadMore(ctx context.Context) (resp DeletedProtectionContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type DeletedProtectionContainersListOperationOptions struct {
	Filter *string
}

func DefaultDeletedProtectionContainersListOperationOptions() DeletedProtectionContainersListOperationOptions {
	return DeletedProtectionContainersListOperationOptions{}
}

func (o DeletedProtectionContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeletedProtectionContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// DeletedProtectionContainersList ...
func (c SoftDeletedContainersClient) DeletedProtectionContainersList(ctx context.Context, id VaultId, options DeletedProtectionContainersListOperationOptions) (resp DeletedProtectionContainersListOperationResponse, err error) {
	req, err := c.preparerForDeletedProtectionContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForDeletedProtectionContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForDeletedProtectionContainersList prepares the DeletedProtectionContainersList request.
func (c SoftDeletedContainersClient) preparerForDeletedProtectionContainersList(ctx context.Context, id VaultId, options DeletedProtectionContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/backupDeletedProtectionContainers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForDeletedProtectionContainersListWithNextLink prepares the DeletedProtectionContainersList request with the given nextLink token.
func (c SoftDeletedContainersClient) preparerForDeletedProtectionContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForDeletedProtectionContainersList handles the response to the DeletedProtectionContainersList request. The method always
// closes the http.Response Body.
func (c SoftDeletedContainersClient) responderForDeletedProtectionContainersList(resp *http.Response) (result DeletedProtectionContainersListOperationResponse, err error) {
	type page struct {
		Values   []ProtectionContainerResource `json:"value"`
		NextLink *string                       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result DeletedProtectionContainersListOperationResponse, err error) {
			req, err := c.preparerForDeletedProtectionContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForDeletedProtectionContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "softdeletedcontainers.SoftDeletedContainersClient", "DeletedProtectionContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// DeletedProtectionContainersListComplete retrieves all of the results into a single object
func (c SoftDeletedContainersClient) DeletedProtectionContainersListComplete(ctx context.Context, id VaultId, options DeletedProtectionContainersListOperationOptions) (DeletedProtectionContainersListCompleteResult, error) {
	return c.DeletedProtectionContainersListCompleteMatchingPredicate(ctx, id, options, ProtectionContainerResourceOperationPredicate{})
}

// DeletedProtectionContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SoftDeletedContainersClient) DeletedProtectionContainersListCompleteMatchingPredicate(ctx context.Context, id VaultId, options DeletedProtectionContainersListOperationOptions, predicate ProtectionContainerResourceOperationPredicate) (resp DeletedProtectionContainersListCompleteResult, err error) {
	items := make([]ProtectionContainerResource, 0)

	page, err := c.DeletedProtectionContainersList(ctx, id, options)
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

	out := DeletedProtectionContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
