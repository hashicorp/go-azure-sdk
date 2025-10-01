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

type BackupProtectableItemsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadProtectableItemResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BackupProtectableItemsListOperationResponse, error)
}

type BackupProtectableItemsListCompleteResult struct {
	Items []WorkloadProtectableItemResource
}

func (r BackupProtectableItemsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BackupProtectableItemsListOperationResponse) LoadMore(ctx context.Context) (resp BackupProtectableItemsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type BackupProtectableItemsListOperationOptions struct {
	Filter *string
}

func DefaultBackupProtectableItemsListOperationOptions() BackupProtectableItemsListOperationOptions {
	return BackupProtectableItemsListOperationOptions{}
}

func (o BackupProtectableItemsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o BackupProtectableItemsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// BackupProtectableItemsList ...
func (c BmsClient) BackupProtectableItemsList(ctx context.Context, id VaultId, options BackupProtectableItemsListOperationOptions) (resp BackupProtectableItemsListOperationResponse, err error) {
	req, err := c.preparerForBackupProtectableItemsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBackupProtectableItemsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBackupProtectableItemsList prepares the BackupProtectableItemsList request.
func (c BmsClient) preparerForBackupProtectableItemsList(ctx context.Context, id VaultId, options BackupProtectableItemsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/backupProtectableItems", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBackupProtectableItemsListWithNextLink prepares the BackupProtectableItemsList request with the given nextLink token.
func (c BmsClient) preparerForBackupProtectableItemsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBackupProtectableItemsList handles the response to the BackupProtectableItemsList request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForBackupProtectableItemsList(resp *http.Response) (result BackupProtectableItemsListOperationResponse, err error) {
	type page struct {
		Values   []WorkloadProtectableItemResource `json:"value"`
		NextLink *string                           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BackupProtectableItemsListOperationResponse, err error) {
			req, err := c.preparerForBackupProtectableItemsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBackupProtectableItemsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectableItemsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BackupProtectableItemsListComplete retrieves all of the results into a single object
func (c BmsClient) BackupProtectableItemsListComplete(ctx context.Context, id VaultId, options BackupProtectableItemsListOperationOptions) (BackupProtectableItemsListCompleteResult, error) {
	return c.BackupProtectableItemsListCompleteMatchingPredicate(ctx, id, options, WorkloadProtectableItemResourceOperationPredicate{})
}

// BackupProtectableItemsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BmsClient) BackupProtectableItemsListCompleteMatchingPredicate(ctx context.Context, id VaultId, options BackupProtectableItemsListOperationOptions, predicate WorkloadProtectableItemResourceOperationPredicate) (resp BackupProtectableItemsListCompleteResult, err error) {
	items := make([]WorkloadProtectableItemResource, 0)

	page, err := c.BackupProtectableItemsList(ctx, id, options)
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

	out := BackupProtectableItemsListCompleteResult{
		Items: items,
	}
	return out, nil
}
