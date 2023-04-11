package securescorecontroldefinitions

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

type SecureScoreControlDefinitionsListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreControlDefinitionItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecureScoreControlDefinitionsListBySubscriptionOperationResponse, error)
}

type SecureScoreControlDefinitionsListBySubscriptionCompleteResult struct {
	Items []SecureScoreControlDefinitionItem
}

func (r SecureScoreControlDefinitionsListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecureScoreControlDefinitionsListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp SecureScoreControlDefinitionsListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SecureScoreControlDefinitionsListBySubscription ...
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp SecureScoreControlDefinitionsListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForSecureScoreControlDefinitionsListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecureScoreControlDefinitionsListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecureScoreControlDefinitionsListBySubscription prepares the SecureScoreControlDefinitionsListBySubscription request.
func (c SecureScoreControlDefinitionsClient) preparerForSecureScoreControlDefinitionsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/secureScoreControlDefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSecureScoreControlDefinitionsListBySubscriptionWithNextLink prepares the SecureScoreControlDefinitionsListBySubscription request with the given nextLink token.
func (c SecureScoreControlDefinitionsClient) preparerForSecureScoreControlDefinitionsListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecureScoreControlDefinitionsListBySubscription handles the response to the SecureScoreControlDefinitionsListBySubscription request. The method always
// closes the http.Response Body.
func (c SecureScoreControlDefinitionsClient) responderForSecureScoreControlDefinitionsListBySubscription(resp *http.Response) (result SecureScoreControlDefinitionsListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []SecureScoreControlDefinitionItem `json:"value"`
		NextLink *string                            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecureScoreControlDefinitionsListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForSecureScoreControlDefinitionsListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecureScoreControlDefinitionsListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecureScoreControlDefinitionsListBySubscriptionComplete retrieves all of the results into a single object
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (SecureScoreControlDefinitionsListBySubscriptionCompleteResult, error) {
	return c.SecureScoreControlDefinitionsListBySubscriptionCompleteMatchingPredicate(ctx, id, SecureScoreControlDefinitionItemOperationPredicate{})
}

// SecureScoreControlDefinitionsListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecureScoreControlDefinitionItemOperationPredicate) (resp SecureScoreControlDefinitionsListBySubscriptionCompleteResult, err error) {
	items := make([]SecureScoreControlDefinitionItem, 0)

	page, err := c.SecureScoreControlDefinitionsListBySubscription(ctx, id)
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

	out := SecureScoreControlDefinitionsListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
