package msiximage

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

type ExpandOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ExpandMsixImage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ExpandOperationResponse, error)
}

type ExpandCompleteResult struct {
	Items []ExpandMsixImage
}

func (r ExpandOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ExpandOperationResponse) LoadMore(ctx context.Context) (resp ExpandOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// Expand ...
func (c MsixImageClient) Expand(ctx context.Context, id HostPoolId, input MSIXImageURI) (resp ExpandOperationResponse, err error) {
	req, err := c.preparerForExpand(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForExpand(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ExpandComplete retrieves all of the results into a single object
func (c MsixImageClient) ExpandComplete(ctx context.Context, id HostPoolId, input MSIXImageURI) (ExpandCompleteResult, error) {
	return c.ExpandCompleteMatchingPredicate(ctx, id, input, ExpandMsixImageOperationPredicate{})
}

// ExpandCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MsixImageClient) ExpandCompleteMatchingPredicate(ctx context.Context, id HostPoolId, input MSIXImageURI, predicate ExpandMsixImageOperationPredicate) (resp ExpandCompleteResult, err error) {
	items := make([]ExpandMsixImage, 0)

	page, err := c.Expand(ctx, id, input)
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

	out := ExpandCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForExpand prepares the Expand request.
func (c MsixImageClient) preparerForExpand(ctx context.Context, id HostPoolId, input MSIXImageURI) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/expandMsixImage", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForExpandWithNextLink prepares the Expand request with the given nextLink token.
func (c MsixImageClient) preparerForExpandWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForExpand handles the response to the Expand request. The method always
// closes the http.Response Body.
func (c MsixImageClient) responderForExpand(resp *http.Response) (result ExpandOperationResponse, err error) {
	type page struct {
		Values   []ExpandMsixImage `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ExpandOperationResponse, err error) {
			req, err := c.preparerForExpandWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForExpand(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "msiximage.MsixImageClient", "Expand", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
