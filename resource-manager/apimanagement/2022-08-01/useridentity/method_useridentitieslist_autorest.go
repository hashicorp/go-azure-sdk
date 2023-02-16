package useridentity

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

type UserIdentitiesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]UserIdentityContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (UserIdentitiesListOperationResponse, error)
}

type UserIdentitiesListCompleteResult struct {
	Items []UserIdentityContract
}

func (r UserIdentitiesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r UserIdentitiesListOperationResponse) LoadMore(ctx context.Context) (resp UserIdentitiesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// UserIdentitiesList ...
func (c UserIdentityClient) UserIdentitiesList(ctx context.Context, id UserId) (resp UserIdentitiesListOperationResponse, err error) {
	req, err := c.preparerForUserIdentitiesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForUserIdentitiesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForUserIdentitiesList prepares the UserIdentitiesList request.
func (c UserIdentityClient) preparerForUserIdentitiesList(ctx context.Context, id UserId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/identities", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForUserIdentitiesListWithNextLink prepares the UserIdentitiesList request with the given nextLink token.
func (c UserIdentityClient) preparerForUserIdentitiesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForUserIdentitiesList handles the response to the UserIdentitiesList request. The method always
// closes the http.Response Body.
func (c UserIdentityClient) responderForUserIdentitiesList(resp *http.Response) (result UserIdentitiesListOperationResponse, err error) {
	type page struct {
		Values   []UserIdentityContract `json:"value"`
		NextLink *string                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result UserIdentitiesListOperationResponse, err error) {
			req, err := c.preparerForUserIdentitiesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForUserIdentitiesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "useridentity.UserIdentityClient", "UserIdentitiesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// UserIdentitiesListComplete retrieves all of the results into a single object
func (c UserIdentityClient) UserIdentitiesListComplete(ctx context.Context, id UserId) (UserIdentitiesListCompleteResult, error) {
	return c.UserIdentitiesListCompleteMatchingPredicate(ctx, id, UserIdentityContractOperationPredicate{})
}

// UserIdentitiesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c UserIdentityClient) UserIdentitiesListCompleteMatchingPredicate(ctx context.Context, id UserId, predicate UserIdentityContractOperationPredicate) (resp UserIdentitiesListCompleteResult, err error) {
	items := make([]UserIdentityContract, 0)

	page, err := c.UserIdentitiesList(ctx, id)
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

	out := UserIdentitiesListCompleteResult{
		Items: items,
	}
	return out, nil
}
