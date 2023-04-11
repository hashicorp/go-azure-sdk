package securescore

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

type SecureScoresListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecureScoresListOperationResponse, error)
}

type SecureScoresListCompleteResult struct {
	Items []SecureScoreItem
}

func (r SecureScoresListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecureScoresListOperationResponse) LoadMore(ctx context.Context) (resp SecureScoresListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SecureScoresList ...
func (c SecureScoreClient) SecureScoresList(ctx context.Context, id commonids.SubscriptionId) (resp SecureScoresListOperationResponse, err error) {
	req, err := c.preparerForSecureScoresList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecureScoresList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecureScoresList prepares the SecureScoresList request.
func (c SecureScoreClient) preparerForSecureScoresList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/secureScores", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSecureScoresListWithNextLink prepares the SecureScoresList request with the given nextLink token.
func (c SecureScoreClient) preparerForSecureScoresListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecureScoresList handles the response to the SecureScoresList request. The method always
// closes the http.Response Body.
func (c SecureScoreClient) responderForSecureScoresList(resp *http.Response) (result SecureScoresListOperationResponse, err error) {
	type page struct {
		Values   []SecureScoreItem `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecureScoresListOperationResponse, err error) {
			req, err := c.preparerForSecureScoresListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecureScoresList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecureScoresListComplete retrieves all of the results into a single object
func (c SecureScoreClient) SecureScoresListComplete(ctx context.Context, id commonids.SubscriptionId) (SecureScoresListCompleteResult, error) {
	return c.SecureScoresListCompleteMatchingPredicate(ctx, id, SecureScoreItemOperationPredicate{})
}

// SecureScoresListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreClient) SecureScoresListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecureScoreItemOperationPredicate) (resp SecureScoresListCompleteResult, err error) {
	items := make([]SecureScoreItem, 0)

	page, err := c.SecureScoresList(ctx, id)
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

	out := SecureScoresListCompleteResult{
		Items: items,
	}
	return out, nil
}
