package incrementalrestorepoints

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

type DiskRestorePointListByRestorePointOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DiskRestorePoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (DiskRestorePointListByRestorePointOperationResponse, error)
}

type DiskRestorePointListByRestorePointCompleteResult struct {
	Items []DiskRestorePoint
}

func (r DiskRestorePointListByRestorePointOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r DiskRestorePointListByRestorePointOperationResponse) LoadMore(ctx context.Context) (resp DiskRestorePointListByRestorePointOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// DiskRestorePointListByRestorePoint ...
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePoint(ctx context.Context, id RestorePointId) (resp DiskRestorePointListByRestorePointOperationResponse, err error) {
	req, err := c.preparerForDiskRestorePointListByRestorePoint(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForDiskRestorePointListByRestorePoint(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForDiskRestorePointListByRestorePoint prepares the DiskRestorePointListByRestorePoint request.
func (c IncrementalRestorePointsClient) preparerForDiskRestorePointListByRestorePoint(ctx context.Context, id RestorePointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/diskRestorePoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForDiskRestorePointListByRestorePointWithNextLink prepares the DiskRestorePointListByRestorePoint request with the given nextLink token.
func (c IncrementalRestorePointsClient) preparerForDiskRestorePointListByRestorePointWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForDiskRestorePointListByRestorePoint handles the response to the DiskRestorePointListByRestorePoint request. The method always
// closes the http.Response Body.
func (c IncrementalRestorePointsClient) responderForDiskRestorePointListByRestorePoint(resp *http.Response) (result DiskRestorePointListByRestorePointOperationResponse, err error) {
	type page struct {
		Values   []DiskRestorePoint `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result DiskRestorePointListByRestorePointOperationResponse, err error) {
			req, err := c.preparerForDiskRestorePointListByRestorePointWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForDiskRestorePointListByRestorePoint(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "incrementalrestorepoints.IncrementalRestorePointsClient", "DiskRestorePointListByRestorePoint", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// DiskRestorePointListByRestorePointComplete retrieves all of the results into a single object
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePointComplete(ctx context.Context, id RestorePointId) (DiskRestorePointListByRestorePointCompleteResult, error) {
	return c.DiskRestorePointListByRestorePointCompleteMatchingPredicate(ctx, id, DiskRestorePointOperationPredicate{})
}

// DiskRestorePointListByRestorePointCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePointCompleteMatchingPredicate(ctx context.Context, id RestorePointId, predicate DiskRestorePointOperationPredicate) (resp DiskRestorePointListByRestorePointCompleteResult, err error) {
	items := make([]DiskRestorePoint, 0)

	page, err := c.DiskRestorePointListByRestorePoint(ctx, id)
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

	out := DiskRestorePointListByRestorePointCompleteResult{
		Items: items,
	}
	return out, nil
}
