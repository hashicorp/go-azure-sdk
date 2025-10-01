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

type BackupProtectionContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProtectionContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (BackupProtectionContainersListOperationResponse, error)
}

type BackupProtectionContainersListCompleteResult struct {
	Items []ProtectionContainerResource
}

func (r BackupProtectionContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r BackupProtectionContainersListOperationResponse) LoadMore(ctx context.Context) (resp BackupProtectionContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type BackupProtectionContainersListOperationOptions struct {
	Filter *string
}

func DefaultBackupProtectionContainersListOperationOptions() BackupProtectionContainersListOperationOptions {
	return BackupProtectionContainersListOperationOptions{}
}

func (o BackupProtectionContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o BackupProtectionContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// BackupProtectionContainersList ...
func (c BmsClient) BackupProtectionContainersList(ctx context.Context, id VaultId, options BackupProtectionContainersListOperationOptions) (resp BackupProtectionContainersListOperationResponse, err error) {
	req, err := c.preparerForBackupProtectionContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForBackupProtectionContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForBackupProtectionContainersList prepares the BackupProtectionContainersList request.
func (c BmsClient) preparerForBackupProtectionContainersList(ctx context.Context, id VaultId, options BackupProtectionContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/backupProtectionContainers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForBackupProtectionContainersListWithNextLink prepares the BackupProtectionContainersList request with the given nextLink token.
func (c BmsClient) preparerForBackupProtectionContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForBackupProtectionContainersList handles the response to the BackupProtectionContainersList request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForBackupProtectionContainersList(resp *http.Response) (result BackupProtectionContainersListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result BackupProtectionContainersListOperationResponse, err error) {
			req, err := c.preparerForBackupProtectionContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForBackupProtectionContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "bms.BmsClient", "BackupProtectionContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// BackupProtectionContainersListComplete retrieves all of the results into a single object
func (c BmsClient) BackupProtectionContainersListComplete(ctx context.Context, id VaultId, options BackupProtectionContainersListOperationOptions) (BackupProtectionContainersListCompleteResult, error) {
	return c.BackupProtectionContainersListCompleteMatchingPredicate(ctx, id, options, ProtectionContainerResourceOperationPredicate{})
}

// BackupProtectionContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BmsClient) BackupProtectionContainersListCompleteMatchingPredicate(ctx context.Context, id VaultId, options BackupProtectionContainersListOperationOptions, predicate ProtectionContainerResourceOperationPredicate) (resp BackupProtectionContainersListCompleteResult, err error) {
	items := make([]ProtectionContainerResource, 0)

	page, err := c.BackupProtectionContainersList(ctx, id, options)
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

	out := BackupProtectionContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
