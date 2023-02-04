package migrates

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareSitesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMwareSite

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (VMwareSitesListBySubscriptionOperationResponse, error)
}

type VMwareSitesListBySubscriptionCompleteResult struct {
	Items []VMwareSite
}

func (r VMwareSitesListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r VMwareSitesListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp VMwareSitesListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// VMwareSitesListBySubscription ...
func (c MigratesClient) VMwareSitesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp VMwareSitesListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForVMwareSitesListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForVMwareSitesListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForVMwareSitesListBySubscription prepares the VMwareSitesListBySubscription request.
func (c MigratesClient) preparerForVMwareSitesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.OffAzure/vmwareSites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForVMwareSitesListBySubscriptionWithNextLink prepares the VMwareSitesListBySubscription request with the given nextLink token.
func (c MigratesClient) preparerForVMwareSitesListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForVMwareSitesListBySubscription handles the response to the VMwareSitesListBySubscription request. The method always
// closes the http.Response Body.
func (c MigratesClient) responderForVMwareSitesListBySubscription(resp *http.Response) (result VMwareSitesListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []VMwareSite `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result VMwareSitesListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForVMwareSitesListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForVMwareSitesListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrates.MigratesClient", "VMwareSitesListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// VMwareSitesListBySubscriptionComplete retrieves all of the results into a single object
func (c MigratesClient) VMwareSitesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (VMwareSitesListBySubscriptionCompleteResult, error) {
	return c.VMwareSitesListBySubscriptionCompleteMatchingPredicate(ctx, id, VMwareSiteOperationPredicate{})
}

// VMwareSitesListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MigratesClient) VMwareSitesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate VMwareSiteOperationPredicate) (resp VMwareSitesListBySubscriptionCompleteResult, err error) {
	items := make([]VMwareSite, 0)

	page, err := c.VMwareSitesListBySubscription(ctx, id)
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

	out := VMwareSitesListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
