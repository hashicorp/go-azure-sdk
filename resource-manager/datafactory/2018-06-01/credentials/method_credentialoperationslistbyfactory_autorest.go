package credentials

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

type CredentialOperationsListByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ManagedIdentityCredentialResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (CredentialOperationsListByFactoryOperationResponse, error)
}

type CredentialOperationsListByFactoryCompleteResult struct {
	Items []ManagedIdentityCredentialResource
}

func (r CredentialOperationsListByFactoryOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r CredentialOperationsListByFactoryOperationResponse) LoadMore(ctx context.Context) (resp CredentialOperationsListByFactoryOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// CredentialOperationsListByFactory ...
func (c CredentialsClient) CredentialOperationsListByFactory(ctx context.Context, id FactoryId) (resp CredentialOperationsListByFactoryOperationResponse, err error) {
	req, err := c.preparerForCredentialOperationsListByFactory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForCredentialOperationsListByFactory(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForCredentialOperationsListByFactory prepares the CredentialOperationsListByFactory request.
func (c CredentialsClient) preparerForCredentialOperationsListByFactory(ctx context.Context, id FactoryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/credentials", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForCredentialOperationsListByFactoryWithNextLink prepares the CredentialOperationsListByFactory request with the given nextLink token.
func (c CredentialsClient) preparerForCredentialOperationsListByFactoryWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForCredentialOperationsListByFactory handles the response to the CredentialOperationsListByFactory request. The method always
// closes the http.Response Body.
func (c CredentialsClient) responderForCredentialOperationsListByFactory(resp *http.Response) (result CredentialOperationsListByFactoryOperationResponse, err error) {
	type page struct {
		Values   []ManagedIdentityCredentialResource `json:"value"`
		NextLink *string                             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result CredentialOperationsListByFactoryOperationResponse, err error) {
			req, err := c.preparerForCredentialOperationsListByFactoryWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForCredentialOperationsListByFactory(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "credentials.CredentialsClient", "CredentialOperationsListByFactory", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// CredentialOperationsListByFactoryComplete retrieves all of the results into a single object
func (c CredentialsClient) CredentialOperationsListByFactoryComplete(ctx context.Context, id FactoryId) (CredentialOperationsListByFactoryCompleteResult, error) {
	return c.CredentialOperationsListByFactoryCompleteMatchingPredicate(ctx, id, ManagedIdentityCredentialResourceOperationPredicate{})
}

// CredentialOperationsListByFactoryCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c CredentialsClient) CredentialOperationsListByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate ManagedIdentityCredentialResourceOperationPredicate) (resp CredentialOperationsListByFactoryCompleteResult, err error) {
	items := make([]ManagedIdentityCredentialResource, 0)

	page, err := c.CredentialOperationsListByFactory(ctx, id)
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

	out := CredentialOperationsListByFactoryCompleteResult{
		Items: items,
	}
	return out, nil
}
