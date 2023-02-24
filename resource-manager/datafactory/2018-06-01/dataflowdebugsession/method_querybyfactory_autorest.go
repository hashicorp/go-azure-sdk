package dataflowdebugsession

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

type QueryByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DataFlowDebugSessionInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (QueryByFactoryOperationResponse, error)
}

type QueryByFactoryCompleteResult struct {
	Items []DataFlowDebugSessionInfo
}

func (r QueryByFactoryOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r QueryByFactoryOperationResponse) LoadMore(ctx context.Context) (resp QueryByFactoryOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// QueryByFactory ...
func (c DataFlowDebugSessionClient) QueryByFactory(ctx context.Context, id FactoryId) (resp QueryByFactoryOperationResponse, err error) {
	req, err := c.preparerForQueryByFactory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForQueryByFactory(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForQueryByFactory prepares the QueryByFactory request.
func (c DataFlowDebugSessionClient) preparerForQueryByFactory(ctx context.Context, id FactoryId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/queryDataFlowDebugSessions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForQueryByFactoryWithNextLink prepares the QueryByFactory request with the given nextLink token.
func (c DataFlowDebugSessionClient) preparerForQueryByFactoryWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForQueryByFactory handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (c DataFlowDebugSessionClient) responderForQueryByFactory(resp *http.Response) (result QueryByFactoryOperationResponse, err error) {
	type page struct {
		Values   []DataFlowDebugSessionInfo `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result QueryByFactoryOperationResponse, err error) {
			req, err := c.preparerForQueryByFactoryWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForQueryByFactory(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "QueryByFactory", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// QueryByFactoryComplete retrieves all of the results into a single object
func (c DataFlowDebugSessionClient) QueryByFactoryComplete(ctx context.Context, id FactoryId) (QueryByFactoryCompleteResult, error) {
	return c.QueryByFactoryCompleteMatchingPredicate(ctx, id, DataFlowDebugSessionInfoOperationPredicate{})
}

// QueryByFactoryCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DataFlowDebugSessionClient) QueryByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate DataFlowDebugSessionInfoOperationPredicate) (resp QueryByFactoryCompleteResult, err error) {
	items := make([]DataFlowDebugSessionInfo, 0)

	page, err := c.QueryByFactory(ctx, id)
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

	out := QueryByFactoryCompleteResult{
		Items: items,
	}
	return out, nil
}
