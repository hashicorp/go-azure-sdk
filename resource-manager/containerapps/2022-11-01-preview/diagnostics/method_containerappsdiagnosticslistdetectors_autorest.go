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

type ContainerAppsDiagnosticsListDetectorsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Diagnostics

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ContainerAppsDiagnosticsListDetectorsOperationResponse, error)
}

type ContainerAppsDiagnosticsListDetectorsCompleteResult struct {
	Items []Diagnostics
}

func (r ContainerAppsDiagnosticsListDetectorsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ContainerAppsDiagnosticsListDetectorsOperationResponse) LoadMore(ctx context.Context) (resp ContainerAppsDiagnosticsListDetectorsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ContainerAppsDiagnosticsListDetectors ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsListDetectors(ctx context.Context, id ContainerAppId) (resp ContainerAppsDiagnosticsListDetectorsOperationResponse, err error) {
	req, err := c.preparerForContainerAppsDiagnosticsListDetectors(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForContainerAppsDiagnosticsListDetectors(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForContainerAppsDiagnosticsListDetectors prepares the ContainerAppsDiagnosticsListDetectors request.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsListDetectors(ctx context.Context, id ContainerAppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectors", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForContainerAppsDiagnosticsListDetectorsWithNextLink prepares the ContainerAppsDiagnosticsListDetectors request with the given nextLink token.
func (c DiagnosticsClient) preparerForContainerAppsDiagnosticsListDetectorsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForContainerAppsDiagnosticsListDetectors handles the response to the ContainerAppsDiagnosticsListDetectors request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForContainerAppsDiagnosticsListDetectors(resp *http.Response) (result ContainerAppsDiagnosticsListDetectorsOperationResponse, err error) {
	type page struct {
		Values   []Diagnostics `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ContainerAppsDiagnosticsListDetectorsOperationResponse, err error) {
			req, err := c.preparerForContainerAppsDiagnosticsListDetectorsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForContainerAppsDiagnosticsListDetectors(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ContainerAppsDiagnosticsListDetectors", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ContainerAppsDiagnosticsListDetectorsComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ContainerAppsDiagnosticsListDetectorsComplete(ctx context.Context, id ContainerAppId) (ContainerAppsDiagnosticsListDetectorsCompleteResult, error) {
	return c.ContainerAppsDiagnosticsListDetectorsCompleteMatchingPredicate(ctx, id, DiagnosticsOperationPredicate{})
}

// ContainerAppsDiagnosticsListDetectorsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ContainerAppsDiagnosticsListDetectorsCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate DiagnosticsOperationPredicate) (resp ContainerAppsDiagnosticsListDetectorsCompleteResult, err error) {
	items := make([]Diagnostics, 0)

	page, err := c.ContainerAppsDiagnosticsListDetectors(ctx, id)
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

	out := ContainerAppsDiagnosticsListDetectorsCompleteResult{
		Items: items,
	}
	return out, nil
}
