package securitycontacts

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

type SecurityContactsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecurityContact

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecurityContactsListOperationResponse, error)
}

type SecurityContactsListCompleteResult struct {
	Items []SecurityContact
}

func (r SecurityContactsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecurityContactsListOperationResponse) LoadMore(ctx context.Context) (resp SecurityContactsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SecurityContactsList ...
func (c SecurityContactsClient) SecurityContactsList(ctx context.Context, id commonids.SubscriptionId) (resp SecurityContactsListOperationResponse, err error) {
	req, err := c.preparerForSecurityContactsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecurityContactsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecurityContactsList prepares the SecurityContactsList request.
func (c SecurityContactsClient) preparerForSecurityContactsList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/securityContacts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSecurityContactsListWithNextLink prepares the SecurityContactsList request with the given nextLink token.
func (c SecurityContactsClient) preparerForSecurityContactsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecurityContactsList handles the response to the SecurityContactsList request. The method always
// closes the http.Response Body.
func (c SecurityContactsClient) responderForSecurityContactsList(resp *http.Response) (result SecurityContactsListOperationResponse, err error) {
	type page struct {
		Values   []SecurityContact `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecurityContactsListOperationResponse, err error) {
			req, err := c.preparerForSecurityContactsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecurityContactsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securitycontacts.SecurityContactsClient", "SecurityContactsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecurityContactsListComplete retrieves all of the results into a single object
func (c SecurityContactsClient) SecurityContactsListComplete(ctx context.Context, id commonids.SubscriptionId) (SecurityContactsListCompleteResult, error) {
	return c.SecurityContactsListCompleteMatchingPredicate(ctx, id, SecurityContactOperationPredicate{})
}

// SecurityContactsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecurityContactsClient) SecurityContactsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecurityContactOperationPredicate) (resp SecurityContactsListCompleteResult, err error) {
	items := make([]SecurityContact, 0)

	page, err := c.SecurityContactsList(ctx, id)
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

	out := SecurityContactsListCompleteResult{
		Items: items,
	}
	return out, nil
}
