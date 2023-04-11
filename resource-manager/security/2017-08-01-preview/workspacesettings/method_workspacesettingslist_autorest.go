package workspacesettings

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkspaceSetting

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceSettingsListOperationResponse, error)
}

type WorkspaceSettingsListCompleteResult struct {
	Items []WorkspaceSetting
}

func (r WorkspaceSettingsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceSettingsListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceSettingsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceSettingsList ...
func (c WorkspaceSettingsClient) WorkspaceSettingsList(ctx context.Context, id commonids.SubscriptionId) (resp WorkspaceSettingsListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSettingsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceSettingsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceSettingsList prepares the WorkspaceSettingsList request.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/workspaceSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceSettingsListWithNextLink prepares the WorkspaceSettingsList request with the given nextLink token.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceSettingsList handles the response to the WorkspaceSettingsList request. The method always
// closes the http.Response Body.
func (c WorkspaceSettingsClient) responderForWorkspaceSettingsList(resp *http.Response) (result WorkspaceSettingsListOperationResponse, err error) {
	type page struct {
		Values   []WorkspaceSetting `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceSettingsListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceSettingsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceSettingsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceSettingsListComplete retrieves all of the results into a single object
func (c WorkspaceSettingsClient) WorkspaceSettingsListComplete(ctx context.Context, id commonids.SubscriptionId) (WorkspaceSettingsListCompleteResult, error) {
	return c.WorkspaceSettingsListCompleteMatchingPredicate(ctx, id, WorkspaceSettingOperationPredicate{})
}

// WorkspaceSettingsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceSettingsClient) WorkspaceSettingsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate WorkspaceSettingOperationPredicate) (resp WorkspaceSettingsListCompleteResult, err error) {
	items := make([]WorkspaceSetting, 0)

	page, err := c.WorkspaceSettingsList(ctx, id)
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

	out := WorkspaceSettingsListCompleteResult{
		Items: items,
	}
	return out, nil
}
