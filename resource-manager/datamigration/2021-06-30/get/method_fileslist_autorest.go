package get

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

type FilesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProjectFile

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (FilesListOperationResponse, error)
}

type FilesListCompleteResult struct {
	Items []ProjectFile
}

func (r FilesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r FilesListOperationResponse) LoadMore(ctx context.Context) (resp FilesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// FilesList ...
func (c GETClient) FilesList(ctx context.Context, id ProjectId) (resp FilesListOperationResponse, err error) {
	req, err := c.preparerForFilesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForFilesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForFilesList prepares the FilesList request.
func (c GETClient) preparerForFilesList(ctx context.Context, id ProjectId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/files", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForFilesListWithNextLink prepares the FilesList request with the given nextLink token.
func (c GETClient) preparerForFilesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForFilesList handles the response to the FilesList request. The method always
// closes the http.Response Body.
func (c GETClient) responderForFilesList(resp *http.Response) (result FilesListOperationResponse, err error) {
	type page struct {
		Values   []ProjectFile `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result FilesListOperationResponse, err error) {
			req, err := c.preparerForFilesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForFilesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "FilesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// FilesListComplete retrieves all of the results into a single object
func (c GETClient) FilesListComplete(ctx context.Context, id ProjectId) (FilesListCompleteResult, error) {
	return c.FilesListCompleteMatchingPredicate(ctx, id, ProjectFileOperationPredicate{})
}

// FilesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GETClient) FilesListCompleteMatchingPredicate(ctx context.Context, id ProjectId, predicate ProjectFileOperationPredicate) (resp FilesListCompleteResult, err error) {
	items := make([]ProjectFile, 0)

	page, err := c.FilesList(ctx, id)
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

	out := FilesListCompleteResult{
		Items: items,
	}
	return out, nil
}
