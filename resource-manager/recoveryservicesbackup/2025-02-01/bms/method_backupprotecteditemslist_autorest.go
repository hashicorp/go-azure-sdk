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

type BackupProtectedItemsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProtectedItemResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BackupProtectedItemsListOperationResponse, error)
}

type BackupProtectedItemsListCompleteResult struct {
	Items []ProtectedItemResource
}

func (r BackupProtectedItemsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BackupProtectedItemsListOperationResponse) LoadMore(ctx context.Context) (resp BackupProtectedItemsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type BackupProtectedItemsListOperationOptions struct {
	Filter *string
}

func DefaultBackupProtectedItemsListOperationOptions() BackupProtectedItemsListOperationOptions {
	return BackupProtectedItemsListOperationOptions{}
}

func (o BackupProtectedItemsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o BackupProtectedItemsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// BackupProtectedItemsList ...
func (c BmsClient) BackupProtectedItemsList(ctx context.Context, id VaultId, options BackupProtectedItemsListOperationOptions) (resp BackupProtectedItemsListOperationResponse, err error) {
	req, err := c.preparerForBackupProtectedItemsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBackupProtectedItemsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBackupProtectedItemsList prepares the BackupProtectedItemsList request.
func (c BmsClient) preparerForBackupProtectedItemsList(ctx context.Context, id VaultId, options BackupProtectedItemsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/backupProtectedItems", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBackupProtectedItemsListWithNextLink prepares the BackupProtectedItemsList request with the given nextLink token.
func (c BmsClient) preparerForBackupProtectedItemsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBackupProtectedItemsList handles the response to the BackupProtectedItemsList request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForBackupProtectedItemsList(resp *http.Response) (result BackupProtectedItemsListOperationResponse, err error) {
	type page struct {
		Values   []ProtectedItemResource `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BackupProtectedItemsListOperationResponse, err error) {
			req, err := c.preparerForBackupProtectedItemsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBackupProtectedItemsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectedItemsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BackupProtectedItemsListComplete retrieves all of the results into a single object
func (c BmsClient) BackupProtectedItemsListComplete(ctx context.Context, id VaultId, options BackupProtectedItemsListOperationOptions) (BackupProtectedItemsListCompleteResult, error) {
	return c.BackupProtectedItemsListCompleteMatchingPredicate(ctx, id, options, ProtectedItemResourceOperationPredicate{})
}

// BackupProtectedItemsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BmsClient) BackupProtectedItemsListCompleteMatchingPredicate(ctx context.Context, id VaultId, options BackupProtectedItemsListOperationOptions, predicate ProtectedItemResourceOperationPredicate) (resp BackupProtectedItemsListCompleteResult, err error) {
	items := make([]ProtectedItemResource, 0)

	page, err := c.BackupProtectedItemsList(ctx, id, options)
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

	out := BackupProtectedItemsListCompleteResult{
		Items: items,
	}
	return out, nil
}
