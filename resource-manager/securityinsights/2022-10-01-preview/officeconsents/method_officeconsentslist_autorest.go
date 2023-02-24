package officeconsents

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

type OfficeConsentsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]OfficeConsent

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (OfficeConsentsListOperationResponse, error)
}

type OfficeConsentsListCompleteResult struct {
	Items []OfficeConsent
}

func (r OfficeConsentsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r OfficeConsentsListOperationResponse) LoadMore(ctx context.Context) (resp OfficeConsentsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// OfficeConsentsList ...
func (c OfficeConsentsClient) OfficeConsentsList(ctx context.Context, id WorkspaceId) (resp OfficeConsentsListOperationResponse, err error) {
	req, err := c.preparerForOfficeConsentsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForOfficeConsentsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForOfficeConsentsList prepares the OfficeConsentsList request.
func (c OfficeConsentsClient) preparerForOfficeConsentsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/officeConsents", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForOfficeConsentsListWithNextLink prepares the OfficeConsentsList request with the given nextLink token.
func (c OfficeConsentsClient) preparerForOfficeConsentsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForOfficeConsentsList handles the response to the OfficeConsentsList request. The method always
// closes the http.Response Body.
func (c OfficeConsentsClient) responderForOfficeConsentsList(resp *http.Response) (result OfficeConsentsListOperationResponse, err error) {
	type page struct {
		Values   []OfficeConsent `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result OfficeConsentsListOperationResponse, err error) {
			req, err := c.preparerForOfficeConsentsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForOfficeConsentsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// OfficeConsentsListComplete retrieves all of the results into a single object
func (c OfficeConsentsClient) OfficeConsentsListComplete(ctx context.Context, id WorkspaceId) (OfficeConsentsListCompleteResult, error) {
	return c.OfficeConsentsListCompleteMatchingPredicate(ctx, id, OfficeConsentOperationPredicate{})
}

// OfficeConsentsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c OfficeConsentsClient) OfficeConsentsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate OfficeConsentOperationPredicate) (resp OfficeConsentsListCompleteResult, err error) {
	items := make([]OfficeConsent, 0)

	page, err := c.OfficeConsentsList(ctx, id)
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

	out := OfficeConsentsListCompleteResult{
		Items: items,
	}
	return out, nil
}
