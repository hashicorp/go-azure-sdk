package sshpublickeys

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/odata"
)

// Copyright (c) TODO, Inc.

type ListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SshPublicKeyResource
	OData        *odata.OData

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBySubscriptionOperationResponse, error)
}

type ListBySubscriptionCompleteResult struct {
	Items []SshPublicKeyResource
}

func (r ListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp ListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBySubscription ...
func (c SshPublicKeysClient) ListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result ListBySubscriptionOperationResponse, err error) {
	req, err := c.Client2.NewGetRequest(ctx, fmt.Sprintf("%s/providers/Microsoft.Compute/sshPublicKeys", id.ID()), defaultApiVersion, odata.Query{})
	if err != nil {
		return
	}

	var resp *base.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	values := &struct {
		Values *[]SshPublicKeyResource `json:"value"`
	}{}
	if err = resp.Unmarshal(values); err != nil {
		return
	}
	result.Model = values.Values

	return
}

// ListBySubscriptionComplete retrieves all the results into a single object
func (c SshPublicKeysClient) ListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (ListBySubscriptionCompleteResult, error) {
	return c.ListBySubscriptionCompleteMatchingPredicate(ctx, id, SshPublicKeyResourceOperationPredicate{})
}

// ListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SshPublicKeysClient) ListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SshPublicKeyResourceOperationPredicate) (resp ListBySubscriptionCompleteResult, err error) {
	items := make([]SshPublicKeyResource, 0)

	page, err := c.ListBySubscription(ctx, id)
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

	out := ListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListBySubscription prepares the ListBySubscription request.
func (c SshPublicKeysClient) preparerForListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Compute/sshPublicKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBySubscriptionWithNextLink prepares the ListBySubscription request with the given nextLink token.
func (c SshPublicKeysClient) preparerForListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBySubscription handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (c SshPublicKeysClient) responderForListBySubscription(resp *http.Response) (result ListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []SshPublicKeyResource `json:"value"`
		NextLink *string                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sshpublickeys.SshPublicKeysClient", "ListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sshpublickeys.SshPublicKeysClient", "ListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sshpublickeys.SshPublicKeysClient", "ListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
