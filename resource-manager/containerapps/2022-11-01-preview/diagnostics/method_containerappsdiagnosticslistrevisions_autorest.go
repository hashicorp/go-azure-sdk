package diagnostics

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

type ContainerAppsDiagnosticsListRevisionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Revision

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ContainerAppsDiagnosticsListRevisionsOperationResponse, error)
}

type ContainerAppsDiagnosticsListRevisionsCompleteResult struct {
	Items []Revision
}

func (r ContainerAppsDiagnosticsListRevisionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ContainerAppsDiagnosticsListRevisionsOperationResponse) LoadMore(ctx context.Context) (resp ContainerAppsDiagnosticsListRevisionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ContainerAppsDiagnosticsListRevisionsOperationOptions struct {
	Filter *string
}

func DefaultContainerAppsDiagnosticsListRevisionsOperationOptions() ContainerAppsDiagnosticsListRevisionsOperationOptions {
	return ContainerAppsDiagnosticsListRevisionsOperationOptions{}
}

func (o ContainerAppsDiagnosticsListRevisionsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ContainerAppsDiagnosticsListRevisionsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ContainerAppsDiagnosticsListRevisions ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisions(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions) (resp ContainerAppsDiagnosticsListRevisionsOperationResponse, err error) {
	req, err := c.preparerForContainerAppsDiagnosticsListRevisions(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForContainerAppsDiagnosticsListRevisions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForContainerAppsDiagnosticsListRevisions prepares the ContainerAppsDiagnosticsListRevisions request.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsListRevisions(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/detectorProperties/revisionsApi/revisions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForContainerAppsDiagnosticsListRevisionsWithNextLink prepares the ContainerAppsDiagnosticsListRevisions request with the given nextLink token.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsListRevisionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForContainerAppsDiagnosticsListRevisions handles the response to the ContainerAppsDiagnosticsListRevisions request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForContainerAppsDiagnosticsListRevisions(resp *http.Response) (result ContainerAppsDiagnosticsListRevisionsOperationResponse, err error) {
	type page struct {
		Values   []Revision `json:"value"`
		NextLink *string    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ContainerAppsDiagnosticsListRevisionsOperationResponse, err error) {
			req, err := c.preparerForContainerAppsDiagnosticsListRevisionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForContainerAppsDiagnosticsListRevisions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListRevisions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ContainerAppsDiagnosticsListRevisionsComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisionsComplete(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions) (ContainerAppsDiagnosticsListRevisionsCompleteResult, error) {
	return c.ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate(ctx, id, options, RevisionOperationPredicate{})
}

// ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions, predicate RevisionOperationPredicate) (resp ContainerAppsDiagnosticsListRevisionsCompleteResult, err error) {
	items := make([]Revision, 0)

	page, err := c.ContainerAppsDiagnosticsListRevisions(ctx, id, options)
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

	out := ContainerAppsDiagnosticsListRevisionsCompleteResult{
		Items: items,
	}
	return out, nil
}
