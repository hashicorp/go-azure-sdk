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

type ScriptCmdletsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ScriptCmdlet

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ScriptCmdletsListOperationResponse, error)
}

type ScriptCmdletsListCompleteResult struct {
	Items []ScriptCmdlet
}

func (r ScriptCmdletsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ScriptCmdletsListOperationResponse) LoadMore(ctx context.Context) (resp ScriptCmdletsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ScriptCmdletsList ...
func (c ScriptsClient) ScriptCmdletsList(ctx context.Context, id ScriptPackageId) (resp ScriptCmdletsListOperationResponse, err error) {
	req, err := c.preparerForScriptCmdletsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForScriptCmdletsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForScriptCmdletsList prepares the ScriptCmdletsList request.
func (c ScriptsClient) preparerForScriptCmdletsList(ctx context.Context, id ScriptPackageId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/scriptCmdlets", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForScriptCmdletsListWithNextLink prepares the ScriptCmdletsList request with the given nextLink token.
func (c ScriptsClient) preparerForScriptCmdletsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForScriptCmdletsList handles the response to the ScriptCmdletsList request. The method always
// closes the http.Response Body.
func (c ScriptsClient) responderForScriptCmdletsList(resp *http.Response) (result ScriptCmdletsListOperationResponse, err error) {
	type page struct {
		Values   []ScriptCmdlet `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ScriptCmdletsListOperationResponse, err error) {
			req, err := c.preparerForScriptCmdletsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForScriptCmdletsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ScriptCmdletsListComplete retrieves all of the results into a single object
func (c ScriptsClient) ScriptCmdletsListComplete(ctx context.Context, id ScriptPackageId) (ScriptCmdletsListCompleteResult, error) {
	return c.ScriptCmdletsListCompleteMatchingPredicate(ctx, id, ScriptCmdletOperationPredicate{})
}

// ScriptCmdletsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ScriptsClient) ScriptCmdletsListCompleteMatchingPredicate(ctx context.Context, id ScriptPackageId, predicate ScriptCmdletOperationPredicate) (resp ScriptCmdletsListCompleteResult, err error) {
	items := make([]ScriptCmdlet, 0)

	page, err := c.ScriptCmdletsList(ctx, id)
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

	out := ScriptCmdletsListCompleteResult{
		Items: items,
	}
	return out, nil
}
