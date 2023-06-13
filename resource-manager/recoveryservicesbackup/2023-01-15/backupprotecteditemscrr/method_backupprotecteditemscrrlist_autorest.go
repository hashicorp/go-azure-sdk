package backupprotecteditemscrr

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

type BackupProtectedItemsCrrListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProtectedItemResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BackupProtectedItemsCrrListOperationResponse, error)
}

type BackupProtectedItemsCrrListCompleteResult struct {
	Items []ProtectedItemResource
}

func (r BackupProtectedItemsCrrListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BackupProtectedItemsCrrListOperationResponse) LoadMore(ctx context.Context) (resp BackupProtectedItemsCrrListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type BackupProtectedItemsCrrListOperationOptions struct {
	Filter *string
}

func DefaultBackupProtectedItemsCrrListOperationOptions() BackupProtectedItemsCrrListOperationOptions {
	return BackupProtectedItemsCrrListOperationOptions{}
}

func (o BackupProtectedItemsCrrListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o BackupProtectedItemsCrrListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// BackupProtectedItemsCrrList ...
func (c BackupProtectedItemsCrrClient) BackupProtectedItemsCrrList(ctx context.Context, id VaultId, options BackupProtectedItemsCrrListOperationOptions) (resp BackupProtectedItemsCrrListOperationResponse, err error) {
	req, err := c.preparerForBackupProtectedItemsCrrList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBackupProtectedItemsCrrList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBackupProtectedItemsCrrList prepares the BackupProtectedItemsCrrList request.
func (c BackupProtectedItemsCrrClient) preparerForBackupProtectedItemsCrrList(ctx context.Context, id VaultId, options BackupProtectedItemsCrrListOperationOptions) (*http.Request, error) {
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

// preparerForBackupProtectedItemsCrrListWithNextLink prepares the BackupProtectedItemsCrrList request with the given nextLink token.
func (c BackupProtectedItemsCrrClient) preparerForBackupProtectedItemsCrrListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBackupProtectedItemsCrrList handles the response to the BackupProtectedItemsCrrList request. The method always
// closes the http.Response Body.
func (c BackupProtectedItemsCrrClient) responderForBackupProtectedItemsCrrList(resp *http.Response) (result BackupProtectedItemsCrrListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BackupProtectedItemsCrrListOperationResponse, err error) {
			req, err := c.preparerForBackupProtectedItemsCrrListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBackupProtectedItemsCrrList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "backupprotecteditemscrr.BackupProtectedItemsCrrClient", "BackupProtectedItemsCrrList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BackupProtectedItemsCrrListComplete retrieves all of the results into a single object
func (c BackupProtectedItemsCrrClient) BackupProtectedItemsCrrListComplete(ctx context.Context, id VaultId, options BackupProtectedItemsCrrListOperationOptions) (BackupProtectedItemsCrrListCompleteResult, error) {
	return c.BackupProtectedItemsCrrListCompleteMatchingPredicate(ctx, id, options, ProtectedItemResourceOperationPredicate{})
}

// BackupProtectedItemsCrrListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BackupProtectedItemsCrrClient) BackupProtectedItemsCrrListCompleteMatchingPredicate(ctx context.Context, id VaultId, options BackupProtectedItemsCrrListOperationOptions, predicate ProtectedItemResourceOperationPredicate) (resp BackupProtectedItemsCrrListCompleteResult, err error) {
	items := make([]ProtectedItemResource, 0)

	page, err := c.BackupProtectedItemsCrrList(ctx, id, options)
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

	out := BackupProtectedItemsCrrListCompleteResult{
		Items: items,
	}
	return out, nil
}
