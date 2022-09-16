package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

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

type WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DedicatedSQLminimalTlsSettings

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse, error)
}

type WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteResult struct {
	Items []DedicatedSQLminimalTlsSettings
}

func (r WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList ...
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList prepares the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList request.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dedicatedSQLminimalTlsSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListWithNextLink prepares the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList request with the given nextLink token.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList handles the response to the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(resp *http.Response) (result WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse, err error) {
	type page struct {
		Values   []DedicatedSQLminimalTlsSettings `json:"value"`
		NextLink *string                          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteMatchingPredicate(ctx, id, DedicatedSQLminimalTlsSettingsOperationPredicate{})
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate DedicatedSQLminimalTlsSettingsOperationPredicate) (resp WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteResult, err error) {
	items := make([]DedicatedSQLminimalTlsSettings, 0)

	page, err := c.WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsList(ctx, id)
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

	out := WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsListCompleteResult{
		Items: items,
	}
	return out, nil
}
