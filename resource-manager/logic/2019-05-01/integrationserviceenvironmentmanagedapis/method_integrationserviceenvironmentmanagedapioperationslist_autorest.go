package integrationserviceenvironmentmanagedapis

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

type IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiOperation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, error)
}

type IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult struct {
	Items []ApiOperation
}

func (r IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse) LoadMore(ctx context.Context) (resp IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// IntegrationServiceEnvironmentManagedApiOperationsList ...
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsList(ctx context.Context, id ManagedApiId) (resp IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, err error) {
	req, err := c.preparerForIntegrationServiceEnvironmentManagedApiOperationsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIntegrationServiceEnvironmentManagedApiOperationsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForIntegrationServiceEnvironmentManagedApiOperationsList prepares the IntegrationServiceEnvironmentManagedApiOperationsList request.
func (c IntegrationServiceEnvironmentManagedApisClient) preparerForIntegrationServiceEnvironmentManagedApiOperationsList(ctx context.Context, id ManagedApiId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/apiOperations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIntegrationServiceEnvironmentManagedApiOperationsListWithNextLink prepares the IntegrationServiceEnvironmentManagedApiOperationsList request with the given nextLink token.
func (c IntegrationServiceEnvironmentManagedApisClient) preparerForIntegrationServiceEnvironmentManagedApiOperationsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIntegrationServiceEnvironmentManagedApiOperationsList handles the response to the IntegrationServiceEnvironmentManagedApiOperationsList request. The method always
// closes the http.Response Body.
func (c IntegrationServiceEnvironmentManagedApisClient) responderForIntegrationServiceEnvironmentManagedApiOperationsList(resp *http.Response) (result IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, err error) {
	type page struct {
		Values   []ApiOperation `json:"value"`
		NextLink *string        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IntegrationServiceEnvironmentManagedApiOperationsListOperationResponse, err error) {
			req, err := c.preparerForIntegrationServiceEnvironmentManagedApiOperationsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIntegrationServiceEnvironmentManagedApiOperationsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "integrationserviceenvironmentmanagedapis.IntegrationServiceEnvironmentManagedApisClient", "IntegrationServiceEnvironmentManagedApiOperationsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// IntegrationServiceEnvironmentManagedApiOperationsListComplete retrieves all of the results into a single object
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsListComplete(ctx context.Context, id ManagedApiId) (IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult, error) {
	return c.IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate(ctx, id, ApiOperationOperationPredicate{})
}

// IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IntegrationServiceEnvironmentManagedApisClient) IntegrationServiceEnvironmentManagedApiOperationsListCompleteMatchingPredicate(ctx context.Context, id ManagedApiId, predicate ApiOperationOperationPredicate) (resp IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult, err error) {
	items := make([]ApiOperation, 0)

	page, err := c.IntegrationServiceEnvironmentManagedApiOperationsList(ctx, id)
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

	out := IntegrationServiceEnvironmentManagedApiOperationsListCompleteResult{
		Items: items,
	}
	return out, nil
}
