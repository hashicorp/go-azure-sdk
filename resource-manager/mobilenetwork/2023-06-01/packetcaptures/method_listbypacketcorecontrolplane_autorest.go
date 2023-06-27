package packetcaptures

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

type ListByPacketCoreControlPlaneOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PacketCapture

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByPacketCoreControlPlaneOperationResponse, error)
}

type ListByPacketCoreControlPlaneCompleteResult struct {
	Items []PacketCapture
}

func (r ListByPacketCoreControlPlaneOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByPacketCoreControlPlaneOperationResponse) LoadMore(ctx context.Context) (resp ListByPacketCoreControlPlaneOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByPacketCoreControlPlane ...
func (c PacketCapturesClient) ListByPacketCoreControlPlane(ctx context.Context, id PacketCoreControlPlaneId) (resp ListByPacketCoreControlPlaneOperationResponse, err error) {
	req, err := c.preparerForListByPacketCoreControlPlane(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByPacketCoreControlPlane(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByPacketCoreControlPlane prepares the ListByPacketCoreControlPlane request.
func (c PacketCapturesClient) preparerForListByPacketCoreControlPlane(ctx context.Context, id PacketCoreControlPlaneId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/packetCaptures", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByPacketCoreControlPlaneWithNextLink prepares the ListByPacketCoreControlPlane request with the given nextLink token.
func (c PacketCapturesClient) preparerForListByPacketCoreControlPlaneWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByPacketCoreControlPlane handles the response to the ListByPacketCoreControlPlane request. The method always
// closes the http.Response Body.
func (c PacketCapturesClient) responderForListByPacketCoreControlPlane(resp *http.Response) (result ListByPacketCoreControlPlaneOperationResponse, err error) {
	type page struct {
		Values   []PacketCapture `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByPacketCoreControlPlaneOperationResponse, err error) {
			req, err := c.preparerForListByPacketCoreControlPlaneWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByPacketCoreControlPlane(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "packetcaptures.PacketCapturesClient", "ListByPacketCoreControlPlane", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByPacketCoreControlPlaneComplete retrieves all of the results into a single object
func (c PacketCapturesClient) ListByPacketCoreControlPlaneComplete(ctx context.Context, id PacketCoreControlPlaneId) (ListByPacketCoreControlPlaneCompleteResult, error) {
	return c.ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx, id, PacketCaptureOperationPredicate{})
}

// ListByPacketCoreControlPlaneCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PacketCapturesClient) ListByPacketCoreControlPlaneCompleteMatchingPredicate(ctx context.Context, id PacketCoreControlPlaneId, predicate PacketCaptureOperationPredicate) (resp ListByPacketCoreControlPlaneCompleteResult, err error) {
	items := make([]PacketCapture, 0)

	page, err := c.ListByPacketCoreControlPlane(ctx, id)
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

	out := ListByPacketCoreControlPlaneCompleteResult{
		Items: items,
	}
	return out, nil
}
