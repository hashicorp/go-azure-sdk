package scripts

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

type ScriptPackagesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ScriptPackage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ScriptPackagesListOperationResponse, error)
}

type ScriptPackagesListCompleteResult struct {
	Items []ScriptPackage
}

func (r ScriptPackagesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ScriptPackagesListOperationResponse) LoadMore(ctx context.Context) (resp ScriptPackagesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ScriptPackagesList ...
func (c ScriptsClient) ScriptPackagesList(ctx context.Context, id PrivateCloudId) (resp ScriptPackagesListOperationResponse, err error) {
	req, err := c.preparerForScriptPackagesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForScriptPackagesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForScriptPackagesList prepares the ScriptPackagesList request.
func (c ScriptsClient) preparerForScriptPackagesList(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/scriptPackages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForScriptPackagesListWithNextLink prepares the ScriptPackagesList request with the given nextLink token.
func (c ScriptsClient) preparerForScriptPackagesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForScriptPackagesList handles the response to the ScriptPackagesList request. The method always
// closes the http.Response Body.
func (c ScriptsClient) responderForScriptPackagesList(resp *http.Response) (result ScriptPackagesListOperationResponse, err error) {
	type page struct {
		Values   []ScriptPackage `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ScriptPackagesListOperationResponse, err error) {
			req, err := c.preparerForScriptPackagesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForScriptPackagesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ScriptPackagesListComplete retrieves all of the results into a single object
func (c ScriptsClient) ScriptPackagesListComplete(ctx context.Context, id PrivateCloudId) (ScriptPackagesListCompleteResult, error) {
	return c.ScriptPackagesListCompleteMatchingPredicate(ctx, id, ScriptPackageOperationPredicate{})
}

// ScriptPackagesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ScriptsClient) ScriptPackagesListCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate ScriptPackageOperationPredicate) (resp ScriptPackagesListCompleteResult, err error) {
	items := make([]ScriptPackage, 0)

	page, err := c.ScriptPackagesList(ctx, id)
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

	out := ScriptPackagesListCompleteResult{
		Items: items,
	}
	return out, nil
}
