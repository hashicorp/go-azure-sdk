package virtualmachinescalesets

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

type GetOSUpgradeHistoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]UpgradeOperationHistoricalStatusInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetOSUpgradeHistoryOperationResponse, error)
}

type GetOSUpgradeHistoryCompleteResult struct {
	Items []UpgradeOperationHistoricalStatusInfo
}

func (r GetOSUpgradeHistoryOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetOSUpgradeHistoryOperationResponse) LoadMore(ctx context.Context) (resp GetOSUpgradeHistoryOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetOSUpgradeHistory ...
func (c VirtualMachineScaleSetsClient) GetOSUpgradeHistory(ctx context.Context, id VirtualMachineScaleSetId) (resp GetOSUpgradeHistoryOperationResponse, err error) {
	req, err := c.preparerForGetOSUpgradeHistory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetOSUpgradeHistory(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// GetOSUpgradeHistoryComplete retrieves all of the results into a single object
func (c VirtualMachineScaleSetsClient) GetOSUpgradeHistoryComplete(ctx context.Context, id VirtualMachineScaleSetId) (GetOSUpgradeHistoryCompleteResult, error) {
	return c.GetOSUpgradeHistoryCompleteMatchingPredicate(ctx, id, UpgradeOperationHistoricalStatusInfoOperationPredicate{})
}

// GetOSUpgradeHistoryCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VirtualMachineScaleSetsClient) GetOSUpgradeHistoryCompleteMatchingPredicate(ctx context.Context, id VirtualMachineScaleSetId, predicate UpgradeOperationHistoricalStatusInfoOperationPredicate) (resp GetOSUpgradeHistoryCompleteResult, err error) {
	items := make([]UpgradeOperationHistoricalStatusInfo, 0)

	page, err := c.GetOSUpgradeHistory(ctx, id)
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

	out := GetOSUpgradeHistoryCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForGetOSUpgradeHistory prepares the GetOSUpgradeHistory request.
func (c VirtualMachineScaleSetsClient) preparerForGetOSUpgradeHistory(ctx context.Context, id VirtualMachineScaleSetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/osUpgradeHistory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetOSUpgradeHistoryWithNextLink prepares the GetOSUpgradeHistory request with the given nextLink token.
func (c VirtualMachineScaleSetsClient) preparerForGetOSUpgradeHistoryWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetOSUpgradeHistory handles the response to the GetOSUpgradeHistory request. The method always
// closes the http.Response Body.
func (c VirtualMachineScaleSetsClient) responderForGetOSUpgradeHistory(resp *http.Response) (result GetOSUpgradeHistoryOperationResponse, err error) {
	type page struct {
		Values   []UpgradeOperationHistoricalStatusInfo `json:"value"`
		NextLink *string                                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetOSUpgradeHistoryOperationResponse, err error) {
			req, err := c.preparerForGetOSUpgradeHistoryWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetOSUpgradeHistory(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachinescalesets.VirtualMachineScaleSetsClient", "GetOSUpgradeHistory", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
