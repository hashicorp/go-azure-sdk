package automationrules

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

type AutomationRulesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AutomationRule

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AutomationRulesListOperationResponse, error)
}

type AutomationRulesListCompleteResult struct {
	Items []AutomationRule
}

func (r AutomationRulesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AutomationRulesListOperationResponse) LoadMore(ctx context.Context) (resp AutomationRulesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AutomationRulesList ...
func (c AutomationRulesClient) AutomationRulesList(ctx context.Context, id WorkspaceId) (resp AutomationRulesListOperationResponse, err error) {
	req, err := c.preparerForAutomationRulesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAutomationRulesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// AutomationRulesListComplete retrieves all of the results into a single object
func (c AutomationRulesClient) AutomationRulesListComplete(ctx context.Context, id WorkspaceId) (AutomationRulesListCompleteResult, error) {
	return c.AutomationRulesListCompleteMatchingPredicate(ctx, id, AutomationRuleOperationPredicate{})
}

// AutomationRulesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AutomationRulesClient) AutomationRulesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate AutomationRuleOperationPredicate) (resp AutomationRulesListCompleteResult, err error) {
	items := make([]AutomationRule, 0)

	page, err := c.AutomationRulesList(ctx, id)
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

	out := AutomationRulesListCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForAutomationRulesList prepares the AutomationRulesList request.
func (c AutomationRulesClient) preparerForAutomationRulesList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/automationRules", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAutomationRulesListWithNextLink prepares the AutomationRulesList request with the given nextLink token.
func (c AutomationRulesClient) preparerForAutomationRulesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAutomationRulesList handles the response to the AutomationRulesList request. The method always
// closes the http.Response Body.
func (c AutomationRulesClient) responderForAutomationRulesList(resp *http.Response) (result AutomationRulesListOperationResponse, err error) {
	type page struct {
		Values   []AutomationRule `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AutomationRulesListOperationResponse, err error) {
			req, err := c.preparerForAutomationRulesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAutomationRulesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
