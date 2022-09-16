package workspaceazureadonlyauthentications

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

type AzureADOnlyAuthenticationsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AzureADOnlyAuthentication

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AzureADOnlyAuthenticationsListOperationResponse, error)
}

type AzureADOnlyAuthenticationsListCompleteResult struct {
	Items []AzureADOnlyAuthentication
}

func (r AzureADOnlyAuthenticationsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AzureADOnlyAuthenticationsListOperationResponse) LoadMore(ctx context.Context) (resp AzureADOnlyAuthenticationsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AzureADOnlyAuthenticationsList ...
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsList(ctx context.Context, id WorkspaceId) (resp AzureADOnlyAuthenticationsListOperationResponse, err error) {
	req, err := c.preparerForAzureADOnlyAuthenticationsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAzureADOnlyAuthenticationsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAzureADOnlyAuthenticationsList prepares the AzureADOnlyAuthenticationsList request.
func (c WorkspaceAzureADOnlyAuthenticationsClient) preparerForAzureADOnlyAuthenticationsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/azureADOnlyAuthentications", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAzureADOnlyAuthenticationsListWithNextLink prepares the AzureADOnlyAuthenticationsList request with the given nextLink token.
func (c WorkspaceAzureADOnlyAuthenticationsClient) preparerForAzureADOnlyAuthenticationsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAzureADOnlyAuthenticationsList handles the response to the AzureADOnlyAuthenticationsList request. The method always
// closes the http.Response Body.
func (c WorkspaceAzureADOnlyAuthenticationsClient) responderForAzureADOnlyAuthenticationsList(resp *http.Response) (result AzureADOnlyAuthenticationsListOperationResponse, err error) {
	type page struct {
		Values   []AzureADOnlyAuthentication `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AzureADOnlyAuthenticationsListOperationResponse, err error) {
			req, err := c.preparerForAzureADOnlyAuthenticationsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAzureADOnlyAuthenticationsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AzureADOnlyAuthenticationsListComplete retrieves all of the results into a single object
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsListComplete(ctx context.Context, id WorkspaceId) (AzureADOnlyAuthenticationsListCompleteResult, error) {
	return c.AzureADOnlyAuthenticationsListCompleteMatchingPredicate(ctx, id, AzureADOnlyAuthenticationOperationPredicate{})
}

// AzureADOnlyAuthenticationsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate AzureADOnlyAuthenticationOperationPredicate) (resp AzureADOnlyAuthenticationsListCompleteResult, err error) {
	items := make([]AzureADOnlyAuthentication, 0)

	page, err := c.AzureADOnlyAuthenticationsList(ctx, id)
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

	out := AzureADOnlyAuthenticationsListCompleteResult{
		Items: items,
	}
	return out, nil
}
