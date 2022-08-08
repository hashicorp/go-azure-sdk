package containerappsauthconfigs

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

type ListByContainerAppOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AuthConfig

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByContainerAppOperationResponse, error)
}

type ListByContainerAppCompleteResult struct {
	Items []AuthConfig
}

func (r ListByContainerAppOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByContainerAppOperationResponse) LoadMore(ctx context.Context) (resp ListByContainerAppOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByContainerApp ...
func (c ContainerAppsAuthConfigsClient) ListByContainerApp(ctx context.Context, id ContainerAppId) (resp ListByContainerAppOperationResponse, err error) {
	req, err := c.preparerForListByContainerApp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByContainerApp(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByContainerApp prepares the ListByContainerApp request.
func (c ContainerAppsAuthConfigsClient) preparerForListByContainerApp(ctx context.Context, id ContainerAppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/authConfigs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByContainerAppWithNextLink prepares the ListByContainerApp request with the given nextLink token.
func (c ContainerAppsAuthConfigsClient) preparerForListByContainerAppWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByContainerApp handles the response to the ListByContainerApp request. The method always
// closes the http.Response Body.
func (c ContainerAppsAuthConfigsClient) responderForListByContainerApp(resp *http.Response) (result ListByContainerAppOperationResponse, err error) {
	type page struct {
		Values   []AuthConfig `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByContainerAppOperationResponse, err error) {
			req, err := c.preparerForListByContainerAppWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByContainerApp(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "containerappsauthconfigs.ContainerAppsAuthConfigsClient", "ListByContainerApp", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByContainerAppComplete retrieves all of the results into a single object
func (c ContainerAppsAuthConfigsClient) ListByContainerAppComplete(ctx context.Context, id ContainerAppId) (ListByContainerAppCompleteResult, error) {
	return c.ListByContainerAppCompleteMatchingPredicate(ctx, id, AuthConfigOperationPredicate{})
}

// ListByContainerAppCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ContainerAppsAuthConfigsClient) ListByContainerAppCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate AuthConfigOperationPredicate) (resp ListByContainerAppCompleteResult, err error) {
	items := make([]AuthConfig, 0)

	page, err := c.ListByContainerApp(ctx, id)
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

	out := ListByContainerAppCompleteResult{
		Items: items,
	}
	return out, nil
}
