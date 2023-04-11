package securescorecontroldefinitions

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

type SecureScoreControlDefinitionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecureScoreControlDefinitionItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SecureScoreControlDefinitionsListOperationResponse, error)
}

type SecureScoreControlDefinitionsListCompleteResult struct {
	Items []SecureScoreControlDefinitionItem
}

func (r SecureScoreControlDefinitionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SecureScoreControlDefinitionsListOperationResponse) LoadMore(ctx context.Context) (resp SecureScoreControlDefinitionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SecureScoreControlDefinitionsList ...
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsList(ctx context.Context) (resp SecureScoreControlDefinitionsListOperationResponse, err error) {
	req, err := c.preparerForSecureScoreControlDefinitionsList(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSecureScoreControlDefinitionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSecureScoreControlDefinitionsList prepares the SecureScoreControlDefinitionsList request.
func (c SecureScoreControlDefinitionsClient) preparerForSecureScoreControlDefinitionsList(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Security/secureScoreControlDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSecureScoreControlDefinitionsListWithNextLink prepares the SecureScoreControlDefinitionsList request with the given nextLink token.
func (c SecureScoreControlDefinitionsClient) preparerForSecureScoreControlDefinitionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSecureScoreControlDefinitionsList handles the response to the SecureScoreControlDefinitionsList request. The method always
// closes the http.Response Body.
func (c SecureScoreControlDefinitionsClient) responderForSecureScoreControlDefinitionsList(resp *http.Response) (result SecureScoreControlDefinitionsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SecureScoreControlDefinitionsListOperationResponse, err error) {
			req, err := c.preparerForSecureScoreControlDefinitionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSecureScoreControlDefinitionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "securescorecontroldefinitions.SecureScoreControlDefinitionsClient", "SecureScoreControlDefinitionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SecureScoreControlDefinitionsListComplete retrieves all of the results into a single object
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsListComplete(ctx context.Context) (SecureScoreControlDefinitionsListCompleteResult, error) {
	return c.SecureScoreControlDefinitionsListCompleteMatchingPredicate(ctx, SecureScoreControlDefinitionItemOperationPredicate{})
}

// SecureScoreControlDefinitionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SecureScoreControlDefinitionsClient) SecureScoreControlDefinitionsListCompleteMatchingPredicate(ctx context.Context, predicate SecureScoreControlDefinitionItemOperationPredicate) (resp SecureScoreControlDefinitionsListCompleteResult, err error) {
	items := make([]SecureScoreControlDefinitionItem, 0)

	page, err := c.SecureScoreControlDefinitionsList(ctx)
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

	out := SecureScoreControlDefinitionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
